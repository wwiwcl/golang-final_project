package cmdutil

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

var cmdPipeline []*exec.Cmd

var Stdin = os.Stdin
var Stderr = os.Stderr
var Stdout = os.Stdout
var Cmd_alive bool
var FilesToClose []*os.File
var pipebuf []*bytes.Buffer
var CwdColor int = 95
var pipe int = -1
var DefaultWd string
var backgroundProcess bool = false

func init() {
	Cmd_alive = true
	err := error(nil)
	DefaultWd, err = initGetcwd()
	if err != nil {
		panic(err)
	}
}

func initGetcwd() (string, error) {
	cmd := exec.Command("pwd")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
