package main

import (
	"fmt"
	"github.com/robinson/gos7"
	"go.uber.org/zap"
	"gotest/common/utils"
	"log"
	"os"
	"time"
)

const (
	tcpDevice = "192.168.56.1"
	rack      = 0
	slot      = 1
)

func CheckError(err error)  {
	utils.Logger.Info("",zap.Any("",err))
}

func main(){
	// TCPClient
	fmt.Printf("connect to %v", tcpDevice)
	fmt.Println()
	handler := gos7.NewTCPClientHandler(tcpDevice, rack, slot)
	handler.Timeout = 200 * time.Second
	handler.IdleTimeout = 200 * time.Second
	handler.Logger = log.New(os.Stdout, "tcp: ", log.LstdFlags)
	// 创建连接，可以基于这个新建多个会话（client）
	err := handler.Connect()
	utils.Logger.Info("",zap.Any("",err))
	defer handler.Close()

	// 新建一个会话
	client := gos7.NewClient(handler)
	address := 100
	start := 0
	size := 4
	buffer := make([]byte, 255)
	value := float32(3) // 写入的数据

	// 向DB中写入数据
	var helper gos7.Helper
	helper.SetValueAt(buffer, 0, value)
	err = client.AGWriteDB(address, start, size, buffer)
	utils.Logger.Info("",zap.Any("",err))
	fmt.Printf("write to db%v start:%v size:%v value:%v", address, start, size, value)
	fmt.Println()

	// 从DB中读取数据
	buf := make([]byte, 255)
	err = client.AGReadDB(address, start, size, buf)
	utils.Logger.Info("",zap.Any("",err))
	var s7 gos7.Helper
	var result float32 // 结果
	s7.GetValueAt(buf, 0, &result)
	fmt.Printf("read value:%v from db%v start:%v size:%v ", result, address, start, size)
	fmt.Println()
}

