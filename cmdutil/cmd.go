package cmdutil

import (
	"os"
	"os/exec"
	"strings"
)

func CloseBuffer() {
	os.Remove(InBufferFile.Name())
	os.Remove(ErrBufferFile.Name())
	os.Remove(OutBufferFile.Name())
}

func NewCmd(c1 *exec.Cmd, c2 *exec.Cmd) *exec.Cmd {
	// run command in c2 with cwd of c1
	cmd := exec.Command(c2.Path, c2.Args[1:]...)
	cmd.Dir, _ = Getcwd(c1)
	return cmd
}

func Getcwd(args ...*exec.Cmd) (string, error) {
	var c *exec.Cmd
	if len(args) > 0 {
		c = args[0]
	} else {
		c = exec.Command("ls")
	}
	if c.Dir != "" {
		return c.Dir, nil
	}
	cmd := exec.Command("pwd")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func Runcmd(c *exec.Cmd, args ...string) error {
	defer resetRedirection()
	defer resetBuffer()
	return pipeline(c, args...)
}
