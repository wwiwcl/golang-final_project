package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"cmdutil"
)

func main() {
	defer cmdutil.CloseBuffer()
	c := exec.Command("ls")
	for cmdutil.Cmd_alive {
		wd, err := cmdutil.Getcwd(c)
		if err != nil {
			fmt.Fprintln(cmdutil.Stderr, err)
		}
		fmt.Fprintf(cmdutil.Stdout, "\033[%dm%s\033[0m", cmdutil.CwdColor, wd)
		fmt.Fprint(cmdutil.Stdout, "$ ")
		reader := bufio.NewReader(cmdutil.Stdin)
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Fprintln(cmdutil.Stderr, "InputError:", err)
			continue
		}
		args := strings.Split(string(input), " ")
		if len(args) > 0 {
			err := cmdutil.Runcmd(c, args...)
			if err != nil {
				fmt.Fprintln(cmdutil.Stderr, err)
			}
		}
	}
}
