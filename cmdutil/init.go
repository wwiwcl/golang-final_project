package cmdutil

import (
	"os"
)

var Stdin = os.Stdin
var Stderr = os.Stderr
var Stdout = os.Stdout
var Cmd_alive bool
var Out []*os.File
var In []*os.File
var Err []*os.File
var InBufferFile *os.File
var OutBufferFile *os.File
var ErrBufferFile *os.File

func init() {
	Stdin = os.Stdin
	Stderr = os.Stderr
	Stdout = os.Stdout
	Cmd_alive = true
	InBufferFile, _ = os.CreateTemp("", ".inbuffer")
	OutBufferFile, _ = os.CreateTemp("", ".outbuffer")
	ErrBufferFile, _ = os.CreateTemp("", ".errbuffer")
	ResetBuffer()
	ResetRedirection()
}
