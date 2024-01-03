package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"cmdutil"
)

func main() {
	defer cmdutil.CloseFiles()
	c := exec.Command("ls")
	for cmdutil.Cmd_alive {
		c.Dir = cmdutil.DefaultWd
		wd := cmdutil.DefaultWd
		fmt.Fprintf(cmdutil.Stdout, "\033[%dm%s\033[0m", cmdutil.CwdColor, wd)
		fmt.Fprint(cmdutil.Stdout, "$ ")
		reader := bufio.NewReader(cmdutil.Stdin)
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Fprintln(cmdutil.Stderr, "InputError:", err)
			continue
		}
		args := strings.Split(string(input), "\"")
		tmp_build_args := []string{}
		tmp_args := []string{}
		if len(args) > 0 {
			for _, arg := range args {
				tmp_args = strings.Split(arg, " ")
				tmp_build_args = append(tmp_build_args, tmp_args...)
			}
			args = tmp_build_args
			err := cmdutil.Runcmd(c, args...)
			if err != nil {
				fmt.Fprintln(cmdutil.Stderr, err)
			}
		}
	}
}
