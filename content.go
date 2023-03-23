package main
import (
	"context"
	"fmt"
	"time"
)

func main() {
	//创建一个context根节点
	root := context.Background()

	//根据根节点派生出子节点
	ctx,cancel := context.WithCancel(root)

	//开启协程
	go test(ctx)

	//十秒后取消子节点context
	time.Sleep(time.Second * 10)
	cancel()

	time.Sleep(time.Second * 1)
}

//将context作为参数传入函数
func test(ctx context.Context)  {
	//循环，每秒检测一次当前context是否被取消
	for {
		time.Sleep(1 * time.Second)
		select {
		//如果检测到当前context被取消（通道被关闭）
		case <-ctx.Done():
			//关闭协程
			fmt.Println("done")
			return
		default:
			//正常运行
			fmt.Println("work")
		}
	}
}