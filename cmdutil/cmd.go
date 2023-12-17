package cmdutil

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var Cmd_alive bool = true
var original_stdin = os.Stdin
var original_stderr = os.Stderr
var original_stdout = os.Stdout

func NewCmd(c1 *exec.Cmd, c2 *exec.Cmd) *exec.Cmd {
	// run command in c2 with cwd of c1
	cmd := exec.Command(c2.Path, c2.Args[1:]...)
	cmd.Dir = Getcwd(c1)
	return cmd
}

/*
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
*/

func InSliceString(e string, slice []string) int {
	i := 0
	for _, s := range slice {
		if s == e {
			return i
		}
	}
	return -1
}

func keysOfStringMap(inputMap map[string]func(c *exec.Cmd, args ...string) error) []string {
	keys := make([]string, 0, len(inputMap))
	for key := range inputMap {
		keys = append(keys, key)
	}
	return keys
}

func JoinPath(pathA string, pathB string) string {
	return filepath.Join(pathA, pathB)
}

func IsAbsPath(path string) bool {
	return filepath.IsAbs(path)
}

func PathExists(c *exec.Cmd, path string) bool {
	return NewCmd(c, exec.Command("ls", path)).Run() == nil
}

func Getcwd(args ...*exec.Cmd) string {
	var c *exec.Cmd
	if len(args) > 0 {
		c = args[0]
	} else {
		c = exec.Command("ls")
	}
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

func redirection(mode int, file *os.File) { // 0: stdin, 1: stdout, 2: stderr
	if mode == 0 {
		os.Stdin = file
	} else if mode == 1 {
		os.Stdout = file
	} else {
		os.Stderr = file
	}
}

func checkRedirection(mode int, args *[]string) (bool, *os.File, error) {
	if mode == 0 {
		rediretInFile := InSliceString("<", *args)
		if rediretInFile >= 0 {
			fileout, err := os.Open((*args)[rediretInFile+1])
			if err != nil {
				return false, os.Stdin, err
			}
			*args = append((*args)[:rediretInFile], (*args)[rediretInFile+2:]...)
			return true, fileout, nil
		}
		return false, os.Stdin, nil
	}
	if mode == 1 {
		rediretOutFile := InSliceString(">", *args)
		if rediretOutFile >= 0 {
			fileout, err := os.Create((*args)[rediretOutFile+1])
			if err != nil {
				return false, os.Stdout, err
			}
			*args = append((*args)[:rediretOutFile], (*args)[rediretOutFile+2:]...)
			return true, fileout, nil
		}
		return false, os.Stdout, nil
	}
	return false, os.Stderr, nil
}

func Runcmd(c *exec.Cmd, command string, args ...string) error {
	defer func() {
		os.Stdin = original_stdin
		os.Stderr = original_stderr
		os.Stdout = original_stdout
	}()
	// redirection stdin
	redirectin, filein, err := checkRedirection(0, &args)
	if err != nil {
		return err
	}
	if redirectin {
		defer filein.Close()
		redirection(0, filein)
	}
	// redirection stdout
	redirectout, fileout, err := checkRedirection(1, &args)
	if err != nil {
		return err
	}
	if redirectout {
		defer fileout.Close()
		redirection(1, fileout)
	}
	// special commands
	if InSliceString(command, keysOfStringMap(command_keyword)) >= 0 {
		return RunSpecCase(c, command, args...)
	}
	// run command
	cNew := NewCmd(c, exec.Command(command, args...))
	output, err := cNew.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(output))
	return nil
}
