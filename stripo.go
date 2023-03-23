package main

import (
	"fmt"
	"strings"
)

func main() {
	tracer := "死神来了,死神byebye"
	comma := strings.Index(tracer, ",")
	fmt.Println(tracer[comma:])
}