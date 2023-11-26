package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func newCmd(c1 *exec.Cmd, c2 *exec.Cmd) *exec.Cmd {
	// run command in c2 with cwd of c1
	cmd := exec.Command(c2.Path, c2.Args[1:]...)
	cmd.Dir = getcwd(c1)
	return cmd
}

func run(c *exec.Cmd, command string, args ...string) error {
	c.Path = command
	c.Args = append(c.Args, args...)
	output, err := c.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func joinPath(pathA string, pathB string) string {
	return filepath.Join(pathA, pathB)
}

func pathExists(c *exec.Cmd, path string) bool {
	return newCmd(c, exec.Command("ls", path)).Run() == nil
}

func getcwd(c *exec.Cmd) string {
	if c.Dir != "" {
		return c.Dir
	}
	cmd := exec.Command("pwd")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return strings.TrimSpace(string(output))
}

func mkdir(c *exec.Cmd, path string) error {
	return newCmd(c, exec.Command("mkdir", "-p", path)).Run()
}

func chdir(c *exec.Cmd, path string) error {
	c.Dir = joinPath(getcwd(c), path)
	return newCmd(c, exec.Command("ls")).Run()
}

func main() {
	cmd := exec.Command("pwd")
	fmt.Println(getcwd(cmd))
}