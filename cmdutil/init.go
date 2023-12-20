package cmdutil

import (
	"os"
	"os/exec"
)

var cmdPipeline []*exec.Cmd

var Stdin = os.Stdin
var Stderr = os.Stderr
var Stdout = os.Stdout
var Cmd_alive bool
var Out []*os.File
var In *os.File
var Err []*os.File
var FilesToClose []*os.File
var InBufferFile *os.File
var OutBufferFile *os.File
var ErrBufferFile *os.File
var CwdColor int = 95
var pipe int = -1

func init() {
	Stdin = os.Stdin
	Stderr = os.Stderr
	Stdout = os.Stdout
	Cmd_alive = true
	InBufferFile, _ = os.CreateTemp("", ".inbuffer")
	OutBufferFile, _ = os.CreateTemp("", ".outbuffer")
}
