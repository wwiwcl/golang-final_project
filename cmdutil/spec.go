package cmdutil

import (
	"fmt"
	"os/exec"
	"strconv"

	"sl"
)

var command_keyword = map[string]func(c *exec.Cmd, args ...string) error{
	"mkdir":     mkdir,
	"chdir":     chdir,
	"cd":        chdir,
	"exit":      exit,
	"sl":        sl.Sl,
	"cls":       cls,
	"starburst": sao,
	"sb":        sao,
	"sao":       sao,
	"cat":       cat,
	"rick":      neverGonnaGiveYouUp,
	"never":     neverGonnaGiveYouUp,
	"8ball":     magic8Ball,
	"color":     setCwdColor,
}

func runSpecCase(c *exec.Cmd, args ...string) error {
	return command_keyword[args[0]](c, args[1:]...)
}

func mkdir(c *exec.Cmd, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("mkdir: missing argument")
	}
	return NewCmd(c, exec.Command("mkdir", "-p", args[0])).Run()
}

func chdir(c *exec.Cmd, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("chdir: missing argument")
	}
	cwd, err := Getcwd(c)
	if err != nil {
		return err
	}
	c.Dir, err = makePath(c, args[0])
	if err != nil {
		return err
	}
	if NewCmd(c, exec.Command("ls")).Run() != nil {
		c.Dir = cwd
		return fmt.Errorf("directory %s does not exist", args[0])
	}
	return nil
}

func cls(c *exec.Cmd, args ...string) error {
	fmt.Println("\033[2J")
	return nil
}

func setCwdColor(c *exec.Cmd, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("color: missing argument")
	}
	tmpCwdColor, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	if ((tmpCwdColor >= 30) && (tmpCwdColor <= 37)) || ((tmpCwdColor >= 90) && (tmpCwdColor <= 97)) {
		CwdColor = tmpCwdColor
	} else {
		return fmt.Errorf("color: invalid color code")
	}
	return nil
}

func exit(c *exec.Cmd, args ...string) error {
	Cmd_alive = false
	return nil
}
