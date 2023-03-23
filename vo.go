package main

import (
	"fmt"
	"strconv"
)

func float64ToInt64(value float64) int64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	str := strconv.FormatFloat(value, 'E', -1, 64)
	num, _ := strconv.ParseInt(str, 10, 0)
	return num
}

func main() {
	if true && (1==1 || 3==2) && true {
		fmt.Println(111111)
	}
}
