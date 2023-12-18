package cmdutil

import (
	"fmt"
	"os"
	"os/exec"
)

func run(c *exec.Cmd, command string, args ...string) error {
	// redirection stdin
	redirectin, filein, err := checkRedirection(0, &args)
	if err != nil {
		return err
	}
	if redirectin {
		resetRedirection(0)
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
		resetRedirection(1)
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
		resetRedirection(2)
		for i := 0; i < len(fileerr); i++ {
			if fileerr[i] != Stdin && fileerr[i] != Stdout && fileerr[i] != Stderr {
				defer fileerr[i].Close()
			}
			redirection(2, fileerr[i])
		}
	}
	defer outputsAfterRun()
	// run specific
	if inSliceString([]string{command}, keysOfStringMap(command_keyword)) >= 0 {
		return runSpecCase(c, command, args...)
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
	pipe := inSliceString([]string{"|"}, args)
	for pipe >= 0 {
		argsAfterPipe := args[pipe+1:]
		args = args[:pipe]
		err := run(c, command, args...)
		if err != nil {
			return err
		}
		args = argsAfterPipe
		pipe = inSliceString([]string{"|"}, args)
		resetBuffer(false)
	}
	err := run(c, command, args[0:]...)
	if err != nil {
		return err
	}
	return nil
}

func pipelinePass() error {
	_, err := InBufferFile.Write(contents)
	InBufferFile.Sync()
	if err != nil {
		return err
	}
	return nil
}
