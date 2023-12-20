package cmdutil

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// func run(c *exec.Cmd, args ...string) error {
// 	defer resetRedirection()
// 	// redirection stdin
// 	redirectin, filein, err := checkRedirection(c, 0, &args)
// 	if err != nil {
// 		return err
// 	}
// 	if redirectin {
// 		resetRedirection(0)
// 		for i := 0; i < len(filein); i++ {
// 			if filein[i] != Stdin && filein[i] != Stdout && filein[i] != Stderr {
// 				defer filein[i].Close()
// 			}
// 			redirection(0, filein[i])
// 		}
// 	}
// 	// redirection stdout
// 	redirectout, fileout, err := checkRedirection(c, 1, &args)
// 	if err != nil {
// 		return err
// 	}
// 	if redirectout {
// 		resetRedirection(1)
// 		for i := 0; i < len(fileout); i++ {
// 			if fileout[i] != Stdin && fileout[i] != Stdout && fileout[i] != Stderr {
// 				defer fileout[i].Close()
// 			}
// 			redirection(1, fileout[i])
// 		}
// 	}
// 	// redirection stderr
// 	redirecterr, fileerr, err := checkRedirection(c, 2, &args)
// 	if err != nil {
// 		return err
// 	}
// 	if redirecterr {
// 		resetRedirection(2)
// 		for i := 0; i < len(fileerr); i++ {
// 			if fileerr[i] != Stdin && fileerr[i] != Stdout && fileerr[i] != Stderr {
// 				defer fileerr[i].Close()
// 			}
// 			redirection(2, fileerr[i])
// 		}
// 	}
// 	if pipe >= 0 {
// 		if Out[0] == Stdout {
// 			if len(Out) == 1 {
// 				Out = []*os.File{}
// 			} else {
// 				Out = Out[1:]
// 			}
// 		}
// 		if Err[0] == Stderr {
// 			if len(Err) == 1 {
// 				Err = []*os.File{}
// 			} else {
// 				Err = Err[1:]
// 			}
// 		}
// 	}
// 	command := args[0]
// 	if len(args) > 1 {
// 		args = args[1:]
// 	} else {
// 		args = []string{}
// 	}
// 	defer outputsAfterRun()
// 	// run specific
// 	if inSliceString([]string{command}, keysOfStringMap(command_keyword)) >= 0 {
// 		return runSpecCase(c, command, args...)
// 	}
// 	// run command
// 	os.Stdin = InBufferFile
// 	if err != nil {
// 		fmt.Fprint(os.Stderr, err.Error())
// 		return err
// 	}
// 	// run pipeline
// 	cNew := NewCmd(c, exec.Command(command, args...))
// 	output, err := cNew.CombinedOutput()
// 	if err != nil {
// 		fmt.Fprint(os.Stderr, err.Error())
// 		return err
// 	}
// 	fmt.Print(string(output))
// 	return nil
// }

func buildRun(c *exec.Cmd, args ...string) error {
	// redirection stdin
	var inRedirection io.Reader
	if len(cmdPipeline) == 0 {
		inRedirection = io.Reader(os.Stdin)
	} else {
		pipeStdout, err := cmdPipeline[len(cmdPipeline)-1].StdoutPipe()
		if err != nil {
			return err
		}
		inRedirection = io.Reader(pipeStdout)
	}
	redirectin, filein, err := checkRedirection(c, 0, &args)
	if err != nil {
		return err
	}
	if redirectin {
		files := make([]io.Reader, len(filein)+1)
		files[0] = inRedirection
		for i := 0; i < len(filein); i++ {
			if filein[i] != Stdin && filein[i] != Stdout && filein[i] != Stderr {
				FilesToClose = append(FilesToClose, filein[i])
			}
			files[i+1] = filein[i]
		}
		inRedirection = io.MultiReader(files...)
	}

	// redirection stdout
	outRedirection := io.MultiWriter(os.Stdout)
	redirectout, fileout, err := checkRedirection(c, 1, &args)
	if err != nil {
		return err
	}
	if redirectout {
		files := make([]io.Writer, len(filein))
		for i := 0; i < len(fileout); i++ {
			if fileout[i] != Stdin && fileout[i] != Stdout && fileout[i] != Stderr {
				FilesToClose = append(FilesToClose, fileout[i])
			}
			files[i] = fileout[i]
		}
		outRedirection = io.MultiWriter(files...)
	}

	// redirection stderr
	errRedirection := io.MultiWriter(os.Stderr)
	redirecterr, fileerr, err := checkRedirection(c, 2, &args)
	if err != nil {
		return err
	}
	if redirecterr {
		files := make([]io.Writer, len(filein))
		for i := 0; i < len(fileerr); i++ {
			if fileerr[i] != Stdin && fileerr[i] != Stdout && fileerr[i] != Stderr {
				FilesToClose = append(FilesToClose, fileerr[i])
			}
			files[i] = fileerr[i]
		}
		errRedirection = io.MultiWriter(files...)
	}
	if pipe >= 0 {
		if Out[0] == Stdout {
			if len(Out) == 1 {
				Out = []*os.File{}
			} else {
				Out = Out[1:]
			}
		}
		if Err[0] == Stderr {
			if len(Err) == 1 {
				Err = []*os.File{}
			} else {
				Err = Err[1:]
			}
		}
	}
	command := args[0]
	if len(args) > 1 {
		args = args[1:]
	} else {
		args = []string{}
	}
	// run command
	os.Stdin = InBufferFile
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return err
	}
	// run pipeline
	cmdPipeline = append(cmdPipeline, c)
	cmdPipeline[len(cmdPipeline)-1] = NewCmd(c, exec.Command(command, args...))
	cmdPipeline[len(cmdPipeline)-1].Stdin = inRedirection
	cmdPipeline[len(cmdPipeline)-1].Stdout = outRedirection
	cmdPipeline[len(cmdPipeline)-1].Stderr = errRedirection
	// output, err := cNew.CombinedOutput()
	// if err != nil {
	// 	fmt.Fprint(os.Stderr, err.Error())
	// 	return err
	// }
	// fmt.Print(string(output))
	return nil
}

func runCase(c *exec.Cmd) error {
	if inSliceString([]string{c.Args[0]}, keysOfStringMap(command_keyword)) >= 0 {
		return runSpecCase(c, c.Args...)
	}
	return c.Run()
}

func runAll() error {
	tasks := make(chan bool, len(cmdPipeline))
	errors := make(chan error, len(cmdPipeline))
	for i := 0; i < len(cmdPipeline); i++ {
		go func(i int) {
			tasks <- true
			err := runCase(cmdPipeline[i])
			if err != nil {
				println(err.Error())
				errors <- err
			}
			<-tasks
		}(i)
	}
	return nil
}

func pipeline(c *exec.Cmd, args ...string) error {
	cmdPipeline = []*exec.Cmd{}
	pipe = inSliceString([]string{"|"}, args)
	for pipe >= 0 {
		if len(args) == pipe+1 {
			println(fmt.Errorf("pipeline: no commands after pipe"))
			args = args[:pipe]
			break
		}
		argsAfterPipe := args[pipe+1:]
		args = args[:pipe]
		err := buildRun(c, args...)
		if err != nil {
			return err
		}
		args = argsAfterPipe
		pipe = inSliceString([]string{"|"}, args)
	}
	err := buildRun(c, args...)
	if err != nil {
		return err
	}
	err = runAll()
	if err != nil {
		return err
	}
	return nil
}
