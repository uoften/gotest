package main

import (
	"fmt"
	"github.com/chromedp/cdproto/network"
)

func main() {
	fmt.Println(network.Enable())
}
