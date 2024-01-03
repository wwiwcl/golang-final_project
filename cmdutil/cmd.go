package cmdutil

import (
	"os"
	"os/exec"
)

func CloseFiles() {
	for _, f := range FilesToClose {
		f.Close()
	}
	FilesToClose = []*os.File{}
}

func NewCmd(c1 *exec.Cmd, c2 *exec.Cmd) *exec.Cmd {
	// run command in c2 with cwd of c1
	cmd := exec.Command(c2.Args[0], c2.Args[1:]...)
	cmd.Dir, _ = Getcwd(c1)
	return cmd
}

func Getcwd(args ...*exec.Cmd) (string, error) {
	var c *exec.Cmd
	if len(args) > 0 {
		c = args[0]
		if c.Dir != "" {
			return c.Dir, nil
		}
	}
	return DefaultWd, nil
}

func Runcmd(c *exec.Cmd, args ...string) error {
	setBackgroundProcess(&args)
	if !backgroundProcess {
		defer CloseFiles()
	}
	return pipeline(c, args...)
}
