package main

import (
	"fmt"
	"sync"
	"time"
)

type cli struct {
	run bool
	startRead bool
	reCount int
	wg sync.WaitGroup
}
// 结果监听服务
func (s *cli) outListener() {
	var err error
	defer func() {
		s.wg.Done()
		fmt.Println("监听结束")
		if err != nil {
			fmt.Printf("host:%s listener close by:%v\n", s.run, err)
		}
		if er := recover(); er != nil {
			fmt.Printf("result geter err:%v\n", err)
		}
	}()
	s.wg.Add(1)
	for s.run {
		fmt.Println(111)
		if s.startRead {
			if err != nil {
				return
			}
		} else {
			time.Sleep(50 * time.Millisecond)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func main()  {
	s := new(cli)
	s.run = true
	go s.outListener()
	time.Sleep(60*time.Millisecond)
	s.startRead = true
	s.Close()
	fmt.Println("程序结束")
}

func (c *cli) Close()  {
	c.run = false
	c.wg.Wait()
}
