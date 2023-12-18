package cmdutil

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ResetBuffer(pipe ...bool) error {
	os.Remove(InBufferFile.Name())
	os.Remove(ErrBufferFile.Name())
	InBufferFile, _ = os.CreateTemp("", ".inbuffer")
	if len(pipe) > 0 {
		err := pipelinePass()
		if err != nil {
			return err
		}
	}
	os.Remove(OutBufferFile.Name())
	OutBufferFile, _ = os.CreateTemp("", ".outbuffer")
	ErrBufferFile, _ = os.CreateTemp("", ".errbuffer")
	os.Stdin = InBufferFile
	os.Stdout = OutBufferFile
	os.Stderr = ErrBufferFile
	return nil
}

var contents []byte

func CloseBuffer() {
	os.Remove(InBufferFile.Name())
	os.Remove(ErrBufferFile.Name())
	os.Remove(OutBufferFile.Name())
}

func pipelinePass() error {
	_, err := InBufferFile.Write(contents)
	InBufferFile.Sync()
	if err != nil {
		return err
	}
	return nil
}

func output(contain []byte) ([]byte, error) {
	for _, outputs := range Out {
		_, err := outputs.Write(contain)
		if err != nil {
			return []byte{}, err
		}
		outputs.Sync()
	}
	return contain, nil
}

func errput(contain []byte) ([]byte, error) {
	for _, outputs := range Err {
		_, err := outputs.Write(contain)
		if err != nil {
			return []byte{}, err
		}
		outputs.Sync()
	}
	return contain, nil
}

func outputsAfterRun() error {
	defer ResetBuffer()
	_, err := os.Stdout.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	contents, err := io.ReadAll(os.Stdout)
	if err != nil {
		return err
	}
	_, err = output(contents)
	if err != nil {
		return err
	}
	contents, err = io.ReadAll(os.Stderr)
	if err != nil {
		return err
	}
	os.Stderr.Close()
	_, err = errput(contents)
	if err != nil {
		return err
	}
	return nil
}

func NewCmd(c1 *exec.Cmd, c2 *exec.Cmd) *exec.Cmd {
	// run command in c2 with cwd of c1
	cmd := exec.Command(c2.Path, c2.Args[1:]...)
	cmd.Dir, _ = Getcwd(c1)
	return cmd
}

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

func redirection(mode int, file *os.File) { // 0: stdin, 1: stdout, 2: stderr
	if mode == 0 {
		os.Stdin = file
	} else if mode == 1 {
		//os.Stdout = file
		Out = append(Out, file)
		if Out[0] == Stdout {
			Out = Out[1:]
		}
	} else {
		//os.Stderr = file
		Err = append(Err, file)
		if Err[0] == Stderr {
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
			returnFiles = append(returnFiles, Stderr)
		}
		rediretOutFile = InSliceString([]string{">&1", "1>&1"}, *args)
		if rediretOutFile >= 0 {
			*args = append((*args)[:rediretOutFile], (*args)[rediretOutFile+1:]...)
			returnFiles = append(returnFiles, Stdout)
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
			returnFiles = append(returnFiles, Stdout)
		}
		rediretErrFile = InSliceString([]string{"2>&2", ">>&2"}, *args)
		if rediretErrFile >= 0 {
			*args = append((*args)[:rediretErrFile], (*args)[rediretErrFile+1:]...)
			returnFiles = append(returnFiles, Stderr)
		}
		return (len(returnFiles) > 0), returnFiles, nil
	}
	return false, []*os.File{}, nil
}

func ResetRedirection(mode ...int) {
	if len(mode) == 0 {
		mode = []int{0, 1, 2}
	}
	if mode[0] == 0 {
		mode = mode[1:]
	}
	if len(mode) == 0 {
		return
	}
	if mode[0] == 1 {
		Out = []*os.File{Stdout}
		mode = mode[1:]
	}
	if len(mode) == 0 {
		return
	}
	if mode[0] == 2 {
		Err = []*os.File{Stderr}
		return
	}
}

func run(c *exec.Cmd, command string, args ...string) error {
	// redirection stdin
	redirectin, filein, err := checkRedirection(0, &args)
	if err != nil {
		return err
	}
	if redirectin {
		ResetRedirection(0)
		for i := 0; i < len(filein); i++ {
			if filein[i] != Stdin && filein[i] != Stdout && filein[i] != Stderr {
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
		ResetRedirection(1)
		for i := 0; i < len(fileout); i++ {
			if fileout[i] != Stdin && fileout[i] != Stdout && fileout[i] != Stderr {
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
		ResetRedirection(2)
		for i := 0; i < len(fileerr); i++ {
			if fileerr[i] != Stdin && fileerr[i] != Stdout && fileerr[i] != Stderr {
				defer fileerr[i].Close()
			}
			redirection(2, fileerr[i])
		}
	}
	defer outputsAfterRun()
	// run specific
	if InSliceString([]string{command}, keysOfStringMap(command_keyword)) >= 0 {
		return RunSpecCase(c, command, args...)
	}
	// run command
	os.Stdin = InBufferFile
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return err
	}
	// run pipeline
	cNew := NewCmd(c, exec.Command(command, args...))
	output, err := cNew.CombinedOutput()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return err
	}
	fmt.Print(string(output))
	return nil
}

func pipeline(c *exec.Cmd, command string, args ...string) error {
	pipe := InSliceString([]string{"|"}, args)
	for pipe >= 0 {
		argsAfterPipe := args[pipe+1:]
		args = args[:pipe]
		err := run(c, command, args...)
		if err != nil {
			return err
		}
		args = argsAfterPipe
		pipe = InSliceString([]string{"|"}, args)
		ResetBuffer(false)
	}
	err := run(c, command, args[0:]...)
	if err != nil {
		return err
	}
	return nil
}

func Runcmd(c *exec.Cmd, command string, args ...string) error {
	defer ResetRedirection()
	defer ResetBuffer()
	return pipeline(c, command, args...)
}
