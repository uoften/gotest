package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var jobs chan int
	//close(jobs) //关闭未赋值的chan会panic: close of nil channel
	jobs = make(chan int)
	//close(jobs) //正常关闭
	timeout := make(chan bool)
	var wg sync.WaitGroup
	go func() {
		time.Sleep(time.Second * 1)
		timeout <- true
	}()

	go func() {
		for i := 1; ; i++ {
			select {
			case <-timeout:
				close(jobs)
				return

			default:
				jobs <- i
				fmt.Println("produce:", i)
			}
			time.Sleep(time.Millisecond*100)
		}
	}()
	//jobs <- 1 //向已经关闭的chan写数据会panic: send on closed channel
	//close(jobs) //关闭已经关闭的chan会panic: close of closed channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range jobs {
			fmt.Println("consume:", i)
			time.Sleep(time.Millisecond*100)
		}
	}()
	wg.Wait()
}