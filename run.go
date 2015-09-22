package pty

import (
	"os"
	"os/exec"
	"syscall"
)

// Start assigns a pseudo-terminal tty os.File to c.Stdin, c.Stdout,
// and c.Stderr, calls c.Start, and returns the File of the tty's
// corresponding pty.
func Start(c *exec.Cmd) (pty *os.File, err error) {
	pty, tty, err := Open()
	if err != nil {
		return nil, err
	}
	defer tty.Close()
	c.Stdout = tty
	c.Stdin = tty
	c.Stderr = tty
	c.SysProcAttr = &syscall.SysProcAttr{Setctty: true, Setsid: true}
	err = c.Start()
	if err != nil {
		pty.Close()
		return nil, err
	}
	return pty, err
}

func Start2(c *exec.Cmd) (pty *os.File, stderr *os.File, err error) {
	stderr, w, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	pty, tty, err := Open()
	if err != nil {
		return nil, nil, err
	}
	defer tty.Close()

	c.Stdout = tty
	c.Stdin = tty
	c.Stderr = w
	c.SysProcAttr = &syscall.SysProcAttr{Setctty: true, Setsid: true}
	err = c.Start()
	if err != nil {
		pty.Close()
		return nil, nil, err
	}

	return pty, stderr, nil
}
