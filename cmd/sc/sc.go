package main

import (
	"os"
	"github.com/kost/gosc/shell"
)

func main() {
	for _, arg := range os.Args[1:] {
		shell.ExecShellCode_b64(arg)
	}
}
