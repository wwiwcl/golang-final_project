package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var cmd_alive bool = true

func newCmd(c1 *exec.Cmd, c2 *exec.Cmd) *exec.Cmd {
	// run command in c2 with cwd of c1
	cmd := exec.Command(c2.Path, c2.Args[1:]...)
	cmd.Dir = getcwd(c1)
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

func InSliceString(e string, slice []string) bool {
	for _, s := range slice {
		if s == e {
			return true
		}
	}
	return false
}

func keysOfStringMap(inputMap map[string]func(c *exec.Cmd, args ...string) error) []string {
	keys := make([]string, 0, len(inputMap))
	for key := range inputMap {
		keys = append(keys, key)
	}
	return keys
}

func joinPath(pathA string, pathB string) string {
	return filepath.Join(pathA, pathB)
}

func isAbsPath(path string) bool {
	return filepath.IsAbs(path)
}

func pathExists(c *exec.Cmd, path string) bool {
	return newCmd(c, exec.Command("ls", path)).Run() == nil
}

func getcwd(args ...*exec.Cmd) string {
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

func mkdir(c *exec.Cmd, args ...string) error {
	fmt.Println(len(args))
	return newCmd(c, exec.Command("mkdir", "-p", args[0])).Run()
}

func chdir(c *exec.Cmd, args ...string) error {
	cwd := getcwd(c)
	c.Dir = joinPath(cwd, args[0])
	if newCmd(c, exec.Command("ls")).Run() != nil {
		c.Dir = cwd
		return fmt.Errorf("directory %s does not exist", args[0])
	}
	return nil
}

func sl(c *exec.Cmd, args ...string) error {
	fmt.Println("sl") // TODO
	return nil
}

func cls(c *exec.Cmd, args ...string) error {
	fmt.Println("\033[2J")
	return nil
}

func exit(c *exec.Cmd, args ...string) error {
	cmd_alive = false
	return nil
}

var command_keyword = map[string]func(c *exec.Cmd, args ...string) error{
	"mkdir":     mkdir,
	"chdir":     chdir,
	"exit":      exit,
	"sl":        sl,
	"cls":       cls,
	"starburst": sao,
	"sb":        sao,
	"sao":       sao,
}

// /////////////////////wip///////////////////////
func runSpecCase(c *exec.Cmd, command string, args ...string) error {
	return command_keyword[command](c, args...)
}

//////////////////////endwip//////////////////////

func runcmd(c *exec.Cmd, command string, args ...string) error {
	if InSliceString(command, keysOfStringMap(command_keyword)) {
		return runSpecCase(c, command, args...)
	}
	cNew := newCmd(c, exec.Command(command, args...))
	output, err := cNew.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Printf(string(output))
	return nil
}

func main() {
	c := exec.Command("ls")
	for cmd_alive {
		getcwd(c)
		//fmt.Println(c.Dir)
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("InputError:", err)
			continue
		}
		args := strings.Split(string(input), " ")
		if len(args) > 0 {
			runcmd(c, args[0], args[1:]...)
		}
	}
}

func sao(c *exec.Cmd, args ...string) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch r.Intn(10) {
	case 0:
		fmt.Println("克萊因表示：那到底是什麼技能啊？")
	case 1:
		fmt.Println("摸......摸頭還要哭")
	case 2:
		fmt.Println("騙人的吧")
	case 3:
		fmt.Println("這就是等級制的MMO的不合理之處")
	case 4:
		fmt.Println("Tips: 這條command也是Switch")
	case 5:
		fmt.Println("拜託你們\n先幫我撐個十秒左右就好")
	case 6:
		fmt.Println("這不是很戲劇化的發展嗎")
	case 7:
		fmt.Println("令 人 晶 彥")
	case 8:
		fmt.Println("是隱藏寶箱，好耶")
	case 9:
		fmt.Println(`1.右劍由左向右揮
2.左劍刺入後從右上帶出，同時跳起
3.轉身，右劍由左向右揮
4.左劍由左向右揮
5.轉身，右手在上，左手在下，作平行狀，由左向右揮，落地
6.跳起，左右手同時舉起，劃X字形，再以相反方向再劃一次X形
7.克萊因表示：那到底是什麼技能啊？
8.跳起，右手突刺，但因遭對手揮拳攻擊臉部而中斷，落地
9.心中默喊：要更快
10.跳起，雙手同時突刺，左劍從左下帶出，右劍從右上帶出，落地
11.跳起，雙劍突刺，左劍揮空後立刻向右傾斜，面向左邊
12.對手突刺，以雙劍進行格擋
13.雙手反手拿劍
14.向左翻滾1圈，翻滾同時雙劍對對手進行迴轉攻擊
15.翻滾完畢後，雙手正手拿劍，置於頭部兩側
16.右劍由左上往右下揮，同時左劍由右下往左上揮
17.左劍由左上往右下揮，右劍由右下往左上揮
18.雙劍置於頭部兩側集氣，因對手揮拳攻擊臉部而中斷
19.心中默唸：還要更快
20.右劍由右上往左下揮
21.對手由左上往右下揮劍，此時彎腰閃避並向左旋轉
23.旋轉完畢同時左劍由右往左揮
24.對手由右往左揮劍，此時右劍由右上往坐下格擋
25.左劍由下往上揮
26.右劍由右上往正下方揮
27.對手由上往下揮劍，此手置右劍在上，左劍在下，作平行狀，由左向右揮
28.左劍在上，右劍在下，作平行狀，由右向左揮
29.左劍在上，右劍在下，作平行狀，由左向右揮
30.右劍突刺，但被對手以左手徒手抓住
31.左劍突刺，結束`)
	}
	return nil
}
