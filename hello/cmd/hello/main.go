package main

import (
	"github.com/jm96441n/gotools/hello"
)

func main() {
	p := hello.NewPrinter()
	p.PrintGreeting()
	p.PrintTime()
}
