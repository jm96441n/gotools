package main

import (
	"os"

	"github.com/jm96441n/gotools/findgo"
)

func main() {
	fileSys := os.DirFS(os.Args[1])
	findgo.Files(fileSys)
}
