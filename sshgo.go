package main

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"log"
	"net"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Cli struct {
	IP       string      //IP地址
	Username string      //用户名
	Password string      //密码
	Port     int         //端口号
	client   *ssh.Client //ssh客户端
	session  *ssh.Session
	buffer   []byte
	t        int
	in       chan string
	out      chan string
	lastCmd  string
	writer   *io.WriteCloser
	reader   *io.Reader
	wg       sync.WaitGroup
	goNum    sync.WaitGroup
	exit     int
	cmdHint  string
}

//创建命令行对象
//@param ip IP地址
//@param username 用户名
//@param password 密码
//@param port 端口号,默认22
func New(ip string, username string, password string, port ...int) *Cli {
	cli := new(Cli)
	cli.IP = ip
	cli.Username = username
	cli.Password = password
	if len(port) <= 0 {
		cli.Port = 22
	} else {
		cli.Port = port[0]
	}
	return cli
}

//连接
func (c *Cli) connect() error {
	var authMethods []ssh.AuthMethod
	keyboardInteractiveChallenge := func(
		user,
		instruction string,
		questions []string,
		echos []bool,
	) (answers []string, err error) {

		if len(questions) == 0 {
			return []string{}, nil
		}

		answers = make([]string, len(questions))
		for i := range questions {
			yes, _ := regexp.MatchString("`*yes*`", questions[i])
			if yes {
				answers[i] = "yes"

			} else {
				answers[i] = c.Password
			}
		}
		return answers, nil
	}
	authMethods = append(authMethods, ssh.KeyboardInteractive(keyboardInteractiveChallenge))
	authMethods = append(authMethods, ssh.Password(c.Password))

	sshConfig := &ssh.ClientConfig{
		Config:            ssh.Config{
			Ciphers:        []string{
				"aes128-ctr",
				"aes192-ctr",
				"aes256-ctr",
				"aes128-gcm@openssh.com",
				"chacha20-poly1305@openssh.com",
				"arcfour256",
				"arcfour128",
				"aes128-cbc",
				"des-cbc",
				"3des-cbc",
			},
			KeyExchanges: []string{
				"curve25519-sha256@libssh.org",
				"ecdh-sha2-nistp256",
				"ecdh-sha2-nistp384",
				"ecdh-sha2-nistp521",
				"diffie-hellman-group14-sha1",
				"diffie-hellman-group1-sha1",
				"diffie-hellman-group-exchange-sha1",
			},
		},
		User: c.Username,
		Auth: authMethods,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	addr := fmt.Sprintf("%s:%d", c.IP, c.Port)
	sshClient, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return err
	}
	c.client = sshClient
	return nil
}
// Close 请勿调整此函数代码顺序，否则可能导致无法正常退出
func (c *Cli) Close() {
	log.Printf("cli ready to exit")
	//设置退出标志，关闭cmd executer
	c.exit = 1
	//关闭result getter
	c.session.Close()
	//等待相关协程退出
	c.goNum.Wait()
	close(c.in)
	close(c.out)
	c.client.Close()
	(*c.writer).Close()
	log.Printf("cli exit over")
}
//过滤换行空格
func rmu0000(s string) string {
	str := make([]rune, 0, len(s))
	for _, v := range []rune(s) {
		if v == 0 {
			continue
		}
		str = append(str, v)
	}
	return string(str)
}

//执行shell
//@param shell shell脚本命令
func (c *Cli) InitSession() error {
	if c.client == nil {
		if err := c.connect(); err != nil {
			return err
		}
	}
	if c.session == nil {
		if session, err := c.client.NewSession(); err != nil {
			log.Fatalf("创建会话异常, %v", err)
			return err
		} else {
			c.in = make(chan string, 5)
			c.out = make(chan string, 5)
			c.buffer = make([]byte, 1024*1024)
			c.wg = sync.WaitGroup{}
			c.goNum = sync.WaitGroup{}

			modes := ssh.TerminalModes{
				ssh.ECHO:          0, // enable echoing
				ssh.ECHOCTL:       0,
				ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
				ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
			}

			// Request pseudo terminal
			if err := session.RequestPty("vt100", 10000, 200, modes); err != nil {
				log.Fatalf("请求伪终端执行命令异常, %v", err)
				return err
			}
			if w, err := session.StdinPipe(); err != nil {
				log.Fatalf("连接输入通道异常， %v", err)
				return err
			} else {
				c.writer = &w
			}
			if r, err := session.StdoutPipe(); err != nil {
				log.Fatalf("连接输出通道异常, %v", err)
				return err
			} else {
				c.reader = &r
			}
			go func() {
				defer c.goNum.Done()
				c.goNum.Add(1)
				log.Println("cmd executer start")
				var cmd string
				for {
					select {
					case cmd = <-c.in:
						defer c.wg.Wait()
						c.wg.Add(1)
						(*c.writer).Write([]byte(cmd + "\n"))
						c.lastCmd = cmd
					default:
						time.Sleep(10 * time.Millisecond)
					}
					if c.exit == 1 {
						break
					}
				}
			}()
			go func() {
				defer c.goNum.Done()
				c.goNum.Add(1)
				log.Println("result geter start")
				last := 0
				var result string
				for {
					if c.exit == 1 {
						break
					}
					if n, err := (*c.reader).Read(c.buffer[c.t:]); err != nil {
						return
					} else if n == 0 {
						if c.out == nil {
							break
						}
						time.Sleep(10 * time.Millisecond)
						continue
					} else {
						c.t += n
						//if strings.Contains(string(c.buffer[c.t-n:c.t]), "More") {
						//	(*c.writer).Write([]byte(" "))
						//	continue
						//}
					}
					result = string(c.buffer[last:c.t])

					//过滤分页More，支持关闭分页时不需要:---- More ----\r\r([ ]+)\r
					//reg := regexp.MustCompile("---- More ----\\r\\r([ ]+)\\r")
					//result = reg.ReplaceAllString(result, "")

					ok, _ := regexp.Match(c.cmdHint, []byte(result))
					if ok {
						if c.lastCmd != "" && strings.Contains(result, c.lastCmd) {
							regcmd := regexp.MustCompile(c.cmdHint)
							result = regcmd.ReplaceAllString(result, "")
							c.out <- result
							c.wg.Done()
							c.lastCmd = ""
						} else {
							//获取最后一行过滤后作为提示符
							flysnowRegexp := regexp.MustCompile("(.*)$")
							params := flysnowRegexp.FindStringSubmatch(result)
							c.cmdHint = rmu0000(params[len(params)-1])
							//result = result+c.cmdHint
						}
						last = c.t
					}
				}
			}()
			c.session = session
			if err := c.session.Shell(); err != nil {
				log.Fatalf("登录shell异常:%v", err)
			}
		}
	}
	return nil
}
// RunTerminal 执行带交互的命令
func (c *Cli) RunTerminal(shell string) (string, error) {
	c.InitSession()
	c.in <- shell
	var ret string
	failed := 0
	for {
		select {
		case ret = <-c.out:
			if strings.Contains(ret, shell) {
				return ret, nil
			}
		default:
			time.Sleep(10 * time.Millisecond)
			failed += 1
		}
		if failed == 1000 {
			fmt.Printf("cmd timeout\n")
			break
		}
	}
	return "", errors.New("cmd execute timeout")
}
type hostResult struct {
	Id     bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Ip     string        `json:"ip"`
	Cmd    string        `json:"cmd"`
	Result string        `json:"result"`
}
const Table = "host_ssh_result"

func withTable(f func(*mgo.Collection) error) error {
	// 连接数据库
	dialInfo := &mgo.DialInfo{
		Addrs: []string{"10.54.7.251:27017"}, //远程(或本地)服务器地址及端口号
		Direct: false,
		Timeout: time.Second * 3,
		Username: "admin",
		Password: "S8sOd$2sD0cp",
		PoolLimit: 4096, // Session.SetPoolLimit
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		fmt.Println("mgo链接错误")
	}
	defer session.Close()
	c := session.DB("test").C(Table)
	return f(c)
}
func main() {
	start := time.Now()
	cli := New("10.54.7.234", "admin", "Runlian@2012", 22)
	commands := []string{
		"screen-length disable",
		"dis version",
		"dis arp",
		//"dis mac-address",
		//"dis cu",
	}
	for _, cmd := range commands {
		hostCmdResult := new(hostResult)
		hostCmdResult.Ip = cli.IP
		hostCmdResult.Cmd = cmd
		if w, err := cli.RunTerminal(cmd); err != nil {
			hostCmdResult.Result = err.Error()
			fmt.Printf("%v", err)
		} else {
			hostCmdResult.Result = w
			fmt.Printf("%v", w)
		}
		withTable(func(c *mgo.Collection) error {
			return c.Insert(hostCmdResult)
		})
	}

	elapsed := time.Since(start)
	fmt.Println("执行完成耗时：", elapsed)
	cli.Close()
	time.Sleep(10 * time.Second)
}