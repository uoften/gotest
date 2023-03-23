package main

import (
	"fmt"
	"sync"
)

var x int = 0

var wgg sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wgg.Done()
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
}
func main() {
	wgg.Add(1)
	//2个协程出现
	go add()
	wgg.Add(1)
	go add()
	wgg.Wait()
	fmt.Println(x)
}
