package main

import (
	"github.com/kost/gosc/shell"
)

func main() {
	for _, arg := range os.Args[1:] {
		ExecShellCode_b64(arg)
	}
}
