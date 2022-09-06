package main

import (
	"os"
	"github.com/kost/gosc/gosh"
)

func main() {
	for _, arg := range os.Args[1:] {
		gosh.RemoteShell(arg)
	}
}
