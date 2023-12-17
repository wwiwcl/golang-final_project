package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"cmdutil"
)

func main() {
	c := exec.Command("ls")
	for cmdutil.Cmd_alive {
		fmt.Print(cmdutil.Getcwd(c))
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("InputError:", err)
			continue
		}
		args := strings.Split(string(input), " ")
		if len(args) > 0 {
			cmdutil.Runcmd(c, args[0], args[1:]...)
		}
	}
}
