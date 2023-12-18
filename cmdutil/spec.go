package cmdutil

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"

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
}

func runSpecCase(c *exec.Cmd, command string, args ...string) error {
	return command_keyword[command](c, args...)
}

func mkdir(c *exec.Cmd, args ...string) error {
	fmt.Println(len(args))
	return NewCmd(c, exec.Command("mkdir", "-p", args[0])).Run()
}

func chdir(c *exec.Cmd, args ...string) error {
	cwd, err := Getcwd(c)
	if err != nil {
		return err
	}
	c.Dir = joinPath(cwd, args[0])
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

func cat(c *exec.Cmd, args ...string) error {
	easteregg := inSliceString([]string{"-cat", "--cat"}, args)
	if easteregg >= 0 {
		args = append(args[:easteregg], args[easteregg+1:]...)
		fmt.Println(` 　　　　　　 ＿＿
　　　　　 ／＞　　フ
　　　　　| 　_　 _ |
　 　　　／` + "`" + ` ミ＿xノ
　　 　 /　　　 　 |
　　　 /　 ヽ　　 ﾉ
　 　 │　　|　|　|
　／￣|　　 |　|　|
　| (￣ヽ＿_ヽ_)__)
　＼二つ`)
	}
	output, err := NewCmd(c, exec.Command("cat", args...)).CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(output))
	return nil
}

func exit(c *exec.Cmd, args ...string) error {
	Cmd_alive = false
	return nil
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
