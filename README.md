# pty (forked version)

Pty is a Go package for using unix pseudo-terminals.

## pty.Start2

stdout and stderr separated.

## Install

    go get github.com/erukiti/pty

## Example

```go
package main

import (
	"github.com/erukiti/pty"
	"io"
	"os"
	"os/exec"
)

func main() {
	c := exec.Command("grep", "--color=auto", "bar")
	f, e, err := pty.Start2(c)
	if err != nil {
		panic(err)
	}

	go func() {
		f.Write([]byte("foo\n"))
		f.Write([]byte("bar\n"))
		f.Write([]byte("baz\n"))
		f.Write([]byte{4}) // EOT
	}()
	io.Copy(os.Stdout, f)
	io.Copy(os.Stderr, e)
}
```
