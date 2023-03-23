package main

import (
	"fmt"
	"time"
)
var (
	chanA chan int
)
func init() {
	chanA = make(chan int,10)
	go func() {
		for i:=0;i<10;i++ {
			time.Sleep(900*time.Millisecond)
			chanA <- i
		}
	}()
}

func main() {
	end:=false
	for {
		select {
		case a,ok := <-chanA:
			if ok {
				fmt.Println(a)
			}
		case <-time.After(1 * time.Second):
			end=true
		}
		if end {
			break
		}
	}
}