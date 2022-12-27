package hello

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type Printer struct {
	Output io.Writer
	Input  io.Reader
	TimeFn func() time.Time
}

func NewPrinter() *Printer {
	return &Printer{
		Output: os.Stdout,
		Input:  os.Stdin,
		TimeFn: time.Now,
	}
}

func (p *Printer) PrintGreeting() {
	scanner := bufio.NewScanner(p.Input)
	var name string
	scanner.Scan()
	name = scanner.Text()

	fmt.Fprintf(p.Output, "hello %s\n", name)
}

func (p *Printer) PrintTime() {
	curTime := p.TimeFn()
	minutes := curTime.Minute()
	militaryTimeHour := curTime.Hour()
	hour := militaryTimeHour % 12
	if hour == 0 {
		hour = 12
	}
	fmt.Fprintf(p.Output, "It's %d minutes past %d\n", minutes, hour)
}
