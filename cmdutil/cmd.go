package cmdutil

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var Cmd_alive bool = true
var original_stdin = os.Stdin
var original_stderr = os.Stderr
var original_stdout = os.Stdout
var Out []io.WriteCloser = []io.WriteCloser{os.Stdout}
var In []io.ReadCloser = []io.ReadCloser{os.Stdin}
var Err []io.WriteCloser = []io.WriteCloser{os.Stderr}

func Input() ([]byte, error) {
	var read []byte
	var err error
	for _, inputs := range In {
		read, err = io.ReadAll(inputs)
		if err != nil {
			return []byte{}, err
		}
	}
	return read, nil
}

func Output(contain []byte) ([]byte, error) {
	for _, outputs := range Out {
		_, err := outputs.Write(contain)
		if err != nil {
			return []byte{}, err
		}
	}
	return contain, nil
}

func Errput(contain string) (string, error) {
	for _, outputs := range Err {
		_, err := io.WriteString(outputs, contain)
		if err != nil {
			return "", err
		}
	}
	return contain, nil
}

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

func InSliceString(e []string, slice []string) int {
	i := 0
	for _, s := range slice {
		for _, v := range e {
			if v == s {
				return i
			}
		}
		i++
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
		In = append(In, file)
		if In[0] == os.Stdin {
			In = In[1:]
		}
	} else if mode == 1 {
		os.Stdout = file
		Out = append(Out, file)
		if Out[0] == os.Stdout {
			Out = Out[1:]
		}
	} else {
		os.Stderr = file
		Err = append(Err, file)
		if Err[0] == os.Stderr {
			Err = Err[1:]
		}
	}
}

func checkRedirection(mode int, args *[]string) (bool, []*os.File, error) {
	var returnFiles []*os.File
	if mode == 0 {
		rediretInFile := InSliceString([]string{"<"}, *args)
		for rediretInFile >= 0 {
			fileout, err := os.Open((*args)[rediretInFile+1])
			if err != nil {
				return false, []*os.File{}, err
			}
			*args = append((*args)[:rediretInFile], (*args)[rediretInFile+2:]...)
			returnFiles = append(returnFiles, fileout)
			rediretInFile = InSliceString([]string{"<"}, *args)
		}
		return (len(returnFiles) > 0), returnFiles, nil
	}
	if mode == 1 {
		rediretOutFile := InSliceString([]string{">", "1>"}, *args)
		for rediretOutFile >= 0 {
			fileout, err := os.Create((*args)[rediretOutFile+1])
			if err != nil {
				return false, []*os.File{}, err
			}
			*args = append((*args)[:rediretOutFile], (*args)[rediretOutFile+2:]...)
			returnFiles = append(returnFiles, fileout)
			rediretOutFile = InSliceString([]string{">", "1>"}, *args)
		}
		rediretOutFile = InSliceString([]string{"1>&2"}, *args)
		if rediretOutFile >= 0 {
			*args = append((*args)[:rediretOutFile], (*args)[rediretOutFile+1:]...)
			returnFiles = append(returnFiles, os.Stderr)
		}
		return (len(returnFiles) > 0), returnFiles, nil
	}
	if mode == 2 {
		rediretErrFile := InSliceString([]string{"2>", ">>"}, *args)
		for rediretErrFile >= 0 {
			fileout, err := os.Create((*args)[rediretErrFile+1])
			if err != nil {
				return false, []*os.File{}, err
			}
			*args = append((*args)[:rediretErrFile], (*args)[rediretErrFile+2:]...)
			returnFiles = append(returnFiles, fileout)
			rediretErrFile = InSliceString([]string{"2>", ">>"}, *args)
		}
		rediretErrFile = InSliceString([]string{"2>&1"}, *args)
		if rediretErrFile >= 0 {
			*args = append((*args)[:rediretErrFile], (*args)[rediretErrFile+1:]...)
			returnFiles = append(returnFiles, os.Stdout)
		}
		return (len(returnFiles) > 0), returnFiles, nil
	}
	return false, []*os.File{}, nil
}

func resetRedirection(mode ...int) {
	if len(mode) == 0 {
		mode = []int{0, 1, 2}
	}
	if mode[0] == 0 {
		os.Stdin = original_stdin
		In = []io.ReadCloser{os.Stdin}
		mode = mode[1:]
	}
	if len(mode) == 0 {
		return
	}
	if mode[0] == 1 {
		os.Stdout = original_stdout
		Out = []io.WriteCloser{os.Stdout}
		mode = mode[1:]
	}
	if len(mode) == 0 {
		return
	}
	if mode[0] == 2 {
		os.Stderr = original_stderr
		Err = []io.WriteCloser{os.Stderr}
		return
	}
}

func Pipeline(c *exec.Cmd, args ...string) error {
	// redirection stdin
	redirectin, filein, err := checkRedirection(0, &args)
	if err != nil {
		return err
	}
	if redirectin {
		resetRedirection(0)
		for i := 0; i < len(filein); i++ {
			if filein[i] != os.Stdin && filein[i] != os.Stdout && filein[i] != os.Stderr {
				defer filein[i].Close()
			}
			redirection(0, filein[i])
		}
	}
	// redirection stdout
	redirectout, fileout, err := checkRedirection(1, &args)
	if err != nil {
		return err
	}
	if redirectout {
		resetRedirection(1)
		for i := 0; i < len(fileout); i++ {
			if fileout[i] != os.Stdin && fileout[i] != os.Stdout && fileout[i] != os.Stderr {
				defer fileout[i].Close()
			}
			redirection(1, fileout[i])
		}
	}
	// redirection stderr
	redirecterr, fileerr, err := checkRedirection(2, &args)
	if err != nil {
		return err
	}
	if redirecterr {
		resetRedirection(2)
		for i := 0; i < len(fileerr); i++ {
			if fileerr[i] != os.Stdin && fileerr[i] != os.Stdout && fileerr[i] != os.Stderr {
				defer fileerr[i].Close()
			}
			redirection(2, fileerr[i])
		}
	}
	return nil
}

func Runcmd(c *exec.Cmd, command string, args ...string) error {
	defer resetRedirection()
	// special commands
	if InSliceString([]string{command}, keysOfStringMap(command_keyword)) >= 0 {
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
