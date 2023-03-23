package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	tests := []string{
		"test1",
		"test2",
	}
	for _, test := range tests {
		wg.Add(1)
		test := test
		go func() {
			handle(test,&wg)
		}()
	}
	wg.Wait()
}

func handle(str string,wg *sync.WaitGroup) {
	fmt.Println(str)
	wg.Done()
}

//输出：
//test2
//test2
