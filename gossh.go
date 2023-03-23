package main

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)
var (
	client         *ssh.Client     //ssh客户端
	session        *ssh.Session    //session会话
	w         *io.WriteCloser //io输入命令
	r         *io.Reader      //io获取结果
	in             chan string     //命令写入管道
	out            chan string     //结果输出管道
	lastCmd string
	exit int
	buf []byte
 	wg sync.WaitGroup
 	goNum sync.WaitGroup
 	reCount int
)

func MuxShell() error {
	config := &ssh.ClientConfig{
		User: "admin",
		Auth: []ssh.AuthMethod{
			ssh.Password("Runlian@2012"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: time.Second * 6,
		Config: ssh.Config{
			Ciphers: []string{
				"aes128-ctr",
				"aes192-ctr",
				"aes256-ctr",
				"aes512-ctr",
				"rsa512-ctr",
				"hmac-sha2-256",
				"hmac-sha2-512",
				"hmac-ripemd160",
				"aes128-gcm@openssh.com",
				"chacha20-poly1305@openssh.com",
				"arcfour256",
				"arcfour128",
				"aes128-cbc",
				"des-cbc",
				"3des-cbc",
				"aes192-cbc",
				"aes256-cbc",
				"curve25519-sha256@libssh.org",
				"ecdh-sha2-nistp256",
				"ecdh-sha2-nistp384",
				"ecdh-sha2-nistp521",
				"diffie-hellman-group14-sha1",
				"diffie-hellman-group1-sha1",
				"diffie-hellman-group-exchange-sha1",
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
	}
	client, err := ssh.Dial("tcp", "10.54.5.154:22", config)
	if err != nil {
		panic(err)
	}else{
		fmt.Println("连接成功")
	}

	if session == nil {
		if session, err = client.NewSession(); err != nil {
			log.Printf("创建会话异常, %v", err)
			return err
		} else {
			in = make(chan string, 5)
			out = make(chan string, 5)
			buf = make([]byte, 1024*1024)
			wg = sync.WaitGroup{}
			goNum = sync.WaitGroup{}

			modes := ssh.TerminalModes{
				ssh.ECHO:          0, // enable echoing
				ssh.ECHOCTL:       0,
				ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kb aud
				ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kb aud
			}

			// Request pseudo terminal
			if err := session.RequestPty("xterm", 10000, 200, modes); err != nil {
				log.Printf("请求伪终端执行命令异常, %v", err)
				return err
			}
			if ww, err := session.StdinPipe(); err != nil {
				log.Printf("连接输入通道异常， %v", err)
				return err
			} else {
				w = &ww
			}
			if rr, err := session.StdoutPipe(); err != nil {
				log.Printf("连接输出通道异常, %v", err)
				return err
			} else {
				r = &rr
			}
			go func() {
				defer func() {
					goNum.Done()
				}()
				goNum.Add(1)
				var cmd string
				for {
					select {
					case cmd = <-in:
						fmt.Println(cmd)
						wg.Wait()
						wg.Add(1)
						(*w).Write([]byte(cmd + "\n"))
						lastCmd = cmd
					default:

					}
					if exit == 1 {
						break
					}
				}
			}()
			go func() {
				defer func() {
					goNum.Done()
				}()
				goNum.Add(1)
				var (
					t    int
					last int
				)
				for {
					if exit == 1 {
						break
					}
					n, err := (*r).Read(buf[t:])
					if err != nil {
						fmt.Println(err)
						//close(in)
						//close(out)
						return
					}
					t += n
					//assuming the $PS1 == 'sh-4.3$ '
					if strings.Contains(string(buf[t-n:t]), ">") {
						out <- string(buf[last:t])
						t = 0
						reCount = 0
						lastCmd = ""
						wg.Done()
						last = t
					}
				}
			}()

			if err := session.Shell(); err != nil {
				log.Fatalf("登录shell异常:%v", err)
				return err
			}
		}
	}
	return nil
}

func asdasd(cmd string) (ret string,err error) {
	err = MuxShell()
	if err != nil {
		log.Printf("异常, %v", err)
	}
	in <- cmd
	for {
		select {
		case ret = <- out:
			if strings.Contains(ret, ">") {
				return ret,nil
			}
		default:
			time.Sleep(10 * time.Millisecond)
			reCount += 1
		}
		if reCount == 1000 {
			break
		}
	}
	return ret,errors.New("timeout")
}

func main() {
	var wgg sync.WaitGroup
	cmds := []string{"111","show arp"}
	for _, cmd := range cmds {
		wgg.Add(1)
		if w, err := asdasd(cmd); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(w)
		}
		wgg.Done()
	}
	exit = 1
	wgg.Wait()

	if session != nil {
		session.Close()
	}
	goNum.Wait()
	close(in)
	close(out)
	if client != nil {
		client.Close()
	}
}
