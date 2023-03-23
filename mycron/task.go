package mycron

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func NewTask() *task {
	return &task{c: cron.New()}
}

type task struct {
	c *cron.Cron
}

func (t *task) Add(expr string, task func()) *task {
	_, err := t.c.AddFunc(expr, task)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func (t *task) Run() {
	t.c.Start()
}