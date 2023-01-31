package main

import "github.com/jm96441n/gotools/pipeline"

func main() {
	pipeline.FromString("hello world\n").Stdout()
}
