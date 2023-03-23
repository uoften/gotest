package main

import (
	"fmt"
	"time"
)

func main()  {

	var dt time.Time
	loc, _ := time.LoadLocation("Asia/Shanghai")
	dt, _ = time.ParseInLocation("2006-01-02", "2021-06-20", loc)
	fmt.Println(dt)

	t := time.Now()
	fmt.Println(int(t.Unix()))
}
