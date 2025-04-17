package utils

import (
	"errors"
	"fmt"
	"golang.org/x/sys/windows"
	"net"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

var cmd *exec.Cmd

func StartupDangoOCR() error {
	if isPortOpen(6666) {
		return errors.New("6666端口已被使用")
	}

	cmd = exec.Command("ocr/startOCR.exe")

	// 隐藏窗口（Windows）
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow:    true,
			CreationFlags: windows.CREATE_NO_WINDOW,
		}
	}

	return cmd.Start()
}

func StopDangoOCR() error {
	if cmd == nil || cmd.Process == nil {
		return errors.New("ocr进程未启动")
	}
	return cmd.Process.Kill()
}

func isPortOpen(port int) bool {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", fmt.Sprintf(":%d", port), timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
