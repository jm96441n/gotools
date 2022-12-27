package count

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Counter struct {
	Output    io.Writer
	CurVal    int
	StopAt    int
	DelayTime time.Duration
}

func NewCounter() *Counter {
	return &Counter{
		Output:    os.Stdout,
		DelayTime: time.Duration(10 * time.Minute),
	}
}

func (c *Counter) Next() int {
	val := c.CurVal
	c.CurVal += 1
	return val
}

func (c *Counter) Run() {
	for c.CurVal <= c.StopAt && c.StopAt != 0 {
		val := c.Next()
		fmt.Fprintln(c.Output, val)
		time.Sleep(c.DelayTime)
	}
}
