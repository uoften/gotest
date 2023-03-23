package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type CpuCount struct {
	start      time.Time
	last       time.Time
	totalStart int64
	totalLast  int64
	useLast    int64
	Pid        int
	Interval   int
	run        bool
	over       bool
	cpunum     int64
}

func readline(fd string) string {
	fi, err := os.Open(fd)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
	n, _, _ := r.ReadLine()
	return strings.TrimSpace(string(n))
}

func InitCpuCount() *CpuCount {
	c := CpuCount{}
	c.Pid = os.Getpid()
	c.initCPU()
	c.start = time.Now()
	c.Interval = 1
	c.run = true
	return &c
}

func (T *CpuCount) initCPU() {
	fi, err := os.Open("/proc/stat")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
	for {
		n, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		line := strings.TrimSpace(string(n))
		infos := strings.Fields(line)
		if infos[0] == "cpu" {
			var tSum int64
			for i, j := range infos {
				if i != 0 {
					t, _ := strconv.ParseInt(j, 10, 64)
					tSum += t
				}
			}
			T.totalStart = tSum
			T.totalLast = tSum
			T.useLast = 0
		} else if strings.HasPrefix(infos[0], "cpu") {
			T.cpunum++
		}
	}
}

func (T *CpuCount) getUse() (int64, int64) {
	sysStat := readline("/proc/stat")
	var tSum, uSum int64
	for i, j := range strings.Fields(sysStat) {
		if i != 0 {
			t, _ := strconv.ParseInt(j, 10, 64)
			tSum += t
		}
	}
	myStat := readline(fmt.Sprintf("/proc/%d/stat", T.Pid))
	for i, j := range strings.Fields(myStat) {
		if i > 12 && i < 17 {
			t, _ := strconv.ParseInt(j, 10, 64)
			uSum += t
		}
	}
	return tSum, uSum
}

func (T *CpuCount) check() {
	step := 0
	for T.run {
		time.Sleep(1 * time.Second)
		if step > T.Interval {
			step = 0
			t, u := T.getUse()
			tdelay := t - T.totalLast
			udelay := u - T.useLast
			use := (udelay * 100) / (tdelay / T.cpunum)
			use2 := float64(use)
			T.totalLast = t
			T.useLast = u
			log.Printf("cpu usage:%.2f%%\n", use2)
		} else {
			step += 1
		}
	}
	T.over = true
}
func (T *CpuCount) stop() {
	T.run = false
	for !T.over {
		time.Sleep(10 * time.Millisecond)
	}
	t, u := T.getUse()
	tDelay := t - T.totalStart
	use := (u * 100) / (tDelay / T.cpunum)
	use2 := float64(use)
	log.Printf("total cpu usage:%.2f%%\n", use2)
}

func main()  {
	cpu := InitCpuCount()
	go cpu.check()
	defer cpu.stop()
}