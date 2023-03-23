package main

import (
	"errors"
	"fmt"
)

type HostCli struct {
	ip string
	Username string
	Password string
}

type HandleGroup interface {
	newCli() int
	Conn(i, j int) bool
	Session(i, j int)
}

func Execute(h HandleGroup) error {
	err := errors.New("new err")
	return err
}

func (h HostCli) newCli() int               { return 1 }
func (h HostCli) Conn(i int, j int) bool { return true }
func (h HostCli) Session(i int, j int)      {}

// 只要参数的绑定方法中实现了接口定义的方法，就允许把参数当作接口类型来调用
func main() {
	cli := HostCli{"127.0.0.1","",""}
	n := Execute(cli)
	fmt.Println(n)
}