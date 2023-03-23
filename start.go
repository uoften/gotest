package main

import (
	"fmt"
	"time"
)

func monthInterval(y int, m time.Month) (firstDay, lastDay time.Time) {
	firstDay = time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
	lastDay = time.Date(y, m+1, 1, 0, 0, 0, -1, time.UTC)
	return firstDay, lastDay
}

func main() {
	var (
		y int
		m time.Month
	)

	y, m, _ = time.Now().Date()
	first, last := monthInterval(y, m)
	fmt.Println(first.Format("2006-01-02"))
	fmt.Println(last.Format("2006-01-02"))

	y, m = 2018, time.Month(2)
	first, last = monthInterval(y, m)
	fmt.Println(first.Format("2006-01-02"))
	fmt.Println(last.Format("2006-01-02"))
}