package main

import (
	"time" // time.Sleep()
	"seehuhn.de/go/ncurses" // https://pkg.go.dev/github.com/seehuhn/go-ncurses
)

// parameters
var ACCIDENT, FLY, L, C bool = false, false, false, false
// -a, -F, -l, -c

// Car type: LOGO
var LOGO = struct {
	HEIGHT, FUNNEL, LENGTH, PATTERNS int

	head		[]string
	head_wheel	[][]string
	coal		[]string
	car			[]string
}{
	HEIGHT:		6,
	FUNNEL:		4,
	LENGTH:		84,
	PATTERNS:	6,

	head: []string{
		"     ++      +------ ",
		"     ||      |+-+ |  ",
		"   /---------|| | |  ",
		"  + ========  +-+ |  ",
	},

	head_wheel: [][]string{
		{
			" _|--O========O~\\-+  ",
			"//// \\_/      \\_/    ",
		},

		{
			" _|--/O========O\\-+  ",
			"//// \\_/      \\_/    ",
		},

		{
			" _|--/~O========O-+  ",
			"//// \\_/      \\_/    ",
		},

		{
			" _|--/~\\------/~\\-+  ",
			"//// \\_O========O    ",
		},

		{
			" _|--/~\\------/~\\-+  ",
			"//// \\O========O/    ",
		},

		{
			" _|--/~\\------/~\\-+  ",
			"//// O========O_/    ",
		},
	},

	coal: []string{
		"____                 ",
		"|   \\@@@@@@@@@@@     ",
		"|    \\@@@@@@@@@@@@@_ ",
		"|                  | ",
		"|__________________| ",
		"   (O)       (O)     ",
	},

	car: []string{
		"____________________ ",
		"|  ___ ___ ___ ___ | ",
		"|  |_| |_| |_| |_| | ",
		"|__________________| ",
		"|__________________| ",
		"   (O)        (O)    ",
	},
}

// Car type: C51
var C51 = struct {
	HEIGHT, FUNNEL, LENGTH, PATTERNS int

	head       []string
	head_wheel [][]string
	coal       []string
}{
	HEIGHT:   11,
	FUNNEL:   7,
	LENGTH:   87,
	PATTERNS: 6,

	head: []string{
		"        ___                                            ",
		"       _|_|_  _     __       __             ___________",
		"    D__/   \\_(_)___|  |__H__|  |_____I_Ii_()|_________|",
		"     | `---'   |:: `--'  H  `--'         |  |___ ___|  ",
		"    +|~~~~~~~~++::~~~~~~~H~~+=====+~~~~~~|~~||_| |_||  ",
		"    ||        | ::       H  +=====+      |  |::  ...|  ",
		"|    | _______|_::-----------------[][]-----|       |  ",
	},

	head_wheel: [][]string{
		{
			"| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__",
			"------'|oOo|=[]=-      ||      ||      |  ||=======_|__",
			"/~\\____|___|/~\\_|  O=======O=======O   |__|+-/~\\_|     ",
			"\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       ",
		},

		{
			"| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__",
			"------'|oOo|=[]=- O=======O=======O    |  ||=======_|__",
			"/~\\____|___|/~\\_|      ||      ||      |__|+-/~\\_|     ",
			"\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       ",
		},

		{
			"| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__",
			"------'|oOo|==[]=- O=======O=======O   |  ||=======_|__",
			"/~\\____|___|/~\\_|      ||      ||      |__|+-/~\\_|     ",
			"\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       ",
		},

		{
			"| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__",
			"------'|oOo|===[]=- O=======O=======O  |  ||=======_|__",
			"/~\\____|___|/~\\_|      ||      ||      |__|+-/~\\_|     ",
			"\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       ",
		},

		{
			"| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__",
			"------'|oOo|===[]=-    ||      ||      |  ||=======_|__",
			"/~\\____|___|/~\\_|    O=======O=======O |__|+-/~\\_|     ",
			"\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       ",
		},

		{
			"| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__",
			"------'|oOo|==[]=-     ||      ||      |  ||=======_|__",
			"/~\\____|___|/~\\_|   O=======O=======O  |__|+-/~\\_|     ",
			"\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       ",
		},
	},

	coal: []string{
		"                              ",
		"                              ",
		"                              ",
		"    _________________         ",
		"   _|                \\_____A  ",
		" =|                        |  ",
		" -|                        |  ",
		"__|________________________|_ ",
		"|__________________________|_ ",
		"   |_D__D__D_|  |_D__D__D_|   ",
		"    \\_/   \\_/    \\_/   \\_/    ",
	},
}

// Car type: D51
var D51 = struct {
	HEIGHT, FUNNEL, LENGTH, PATTERNS int

	head		[]string
	head_wheel	[][]string
	coal		[]string
}{
	HEIGHT:		10,
	FUNNEL:		7,
	LENGTH:		83,
	PATTERNS:	6,

	head: []string{
		"      ====        ________                ___________ ",
		"  _D _|  |_______/        \\__I_I_____===__|_________| ",
		"   |(_)---  |   H\\________/ |   |        =|___ ___|   ",
		"   /     |  |   H  |  |     |   |         ||_| |_||   ",
		"  |      |  |   H  |__--------------------| [___] |   ",
		"  | ________|___H__/__|_____/[][]~\\_______|       |   ",
		"  |/ |   |-----------I_____I [][] []  D   |=======|__ ",
	},

	head_wheel: [][]string{
		{
			"__/ =| o |=-~~\\  /~~\\  /~~\\  /~~\\ ____Y___________|__ ",
			" |/-=|___|=    ||    ||    ||    |_____/~\\___/        ",
			"  \\_/      \\O=====O=====O=====O_/      \\_/            ",
		},

		{
			"__/ =| o |=-~~\\  /~~\\  /~~\\  /~~\\ ____Y___________|__ ",
			" |/-=|___|=O=====O=====O=====O   |_____/~\\___/        ",
			"  \\_/      \\__/  \\__/  \\__/  \\__/      \\_/            ",
		},

		{
			"__/ =| o |=-O=====O=====O=====O \\ ____Y___________|__ ",
			" |/-=|___|=    ||    ||    ||    |_____/~\\___/        ",
			"  \\_/      \\__/  \\__/  \\__/  \\__/      \\_/            ",
		},

		{
			"__/ =| o |=-~O=====O=====O=====O\\ ____Y___________|__ ",
			" |/-=|___|=    ||    ||    ||    |_____/~\\___/        ",
			"  \\_/      \\__/  \\__/  \\__/  \\__/      \\_/            ",
		},

		{
			"__/ =| o |=-~~\\  /~~\\  /~~\\  /~~\\ ____Y___________|__ ",
			" |/-=|___|=   O=====O=====O=====O|_____/~\\___/        ",
			"  \\_/      \\__/  \\__/  \\__/  \\__/      \\_/            ",
		},

		{
			"__/ =| o |=-~~\\  /~~\\  /~~\\  /~~\\ ____Y___________|__ ",
			" |/-=|___|=    ||    ||    ||    |_____/~\\___/        ",
			"  \\_/      \\_O=====O=====O=====O/      \\_/            ",
		},
	},

	coal: []string{
		"                              ",
		"                              ",
		"    _________________         ",
		"   _|                \\_____A  ",
		" =|                        |  ",
		" -|                        |  ",
		"__|________________________|_ ",
		"|__________________________|_ ",
		"   |_D__D__D_|  |_D__D__D_|   ",
		"    \\_/   \\_/    \\_/   \\_/    ",
	},
}

func my_MvAddStr(win *ncurses.Window, y int, x int, str string) {
	TERM_LINES, TERM_WIDTH := win.GetMaxYX()

    if y >= TERM_LINES || y < 0 || x >= TERM_WIDTH {
        return
    }else if x < 0 && len(str) + x >= 0 {
		win.MvAddStr(y, 0, str[0 - x:])
	}else if x + len(str) >= TERM_WIDTH && x < TERM_WIDTH {
		win.MvAddStr(y, x, str[:TERM_WIDTH - x])
	}else if x >= 0 && x < TERM_WIDTH {
		win.MvAddStr(y, x, str[:])
	}
}

// Smoke
func add_smoke(win *ncurses.Window, y int, x int) {
	var smoke = [][]string{
		{	"(   )", "(    )", "(    )", "(   )", "(  )",
			"(  )" , "( )"   , "( )"   , "()"   , "()"  ,
			"O"    , "O"     , "O"     , "O"    , "O"   ,
			" ",											},
		{	"(@@@)", "(@@@@)", "(@@@@)", "(@@@)", "(@@)",
			"(@@)" , "(@)"   , "(@)"   , "@@"   , "@@"  ,
			"@"    , "@"     , "@"     , "@"    , "@"   ,
			" ",											},
	}

	var dy = []int{  2,  1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 };
	var dx = []int{ -2, -1, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3 };


	var tempY, tempX, pattern int = y, x - (x % 4), (x / 2) % 2
	if pattern < 0 {
		pattern = - pattern
	}

	for i := 0; i < 16; i++{
		my_MvAddStr(win, tempY, tempX, smoke[pattern][i])

		tempY -= dy[i]
		tempX += dx[i] + len(smoke[pattern][i])

		pattern = 1 - pattern
	}
}

// Output a single frame
func add_LOGO(win *ncurses.Window, x int) {
	// Get the size of the terminal
	TERM_LINES, _ := win.GetMaxYX() // _ = TERM_WIDTH
	var y = TERM_LINES / 2 - 3

	for i := 0; i < LOGO.HEIGHT; i++ {
		if i < 4 {
			my_MvAddStr(win, y + i, x, LOGO.head[i])
		}else {
			my_MvAddStr(win, y + i, x, LOGO.head_wheel[(LOGO.LENGTH + x) % LOGO.PATTERNS][i - 4])
		}

		my_MvAddStr(win, y + i, x + 21, LOGO.coal[i])
		my_MvAddStr(win, y + i, x + 42, LOGO.car[i])
		my_MvAddStr(win, y + i, x + 63, LOGO.car[i])
	}

	if ACCIDENT == true {
		add_man(win, y + 1, x + 14)

		add_man(win, y + 1, x + 45)
		add_man(win, y + 1, x + 53)

		add_man(win, y + 1, x + 66)
		add_man(win, y + 1, x + 74)
	}

	add_smoke(win, y - 1, x + LOGO.FUNNEL)
}

func add_D51(win *ncurses.Window, x int) {
	// Get the size of the terminal
	TERM_LINES, _ := win.GetMaxYX() // _ = TERM_WIDTH
	var y = TERM_LINES / 2 - 5

	for i := 0; i < D51.HEIGHT; i++ {
		if i < 7 {
			my_MvAddStr(win, y + i, x, D51.head[i])
		}else {
			my_MvAddStr(win, y + i, x, D51.head_wheel[(D51.LENGTH + x) % D51.PATTERNS][i - 7])
		}

		my_MvAddStr(win, y + i, x + 53, D51.coal[i])
	}
	
	if ACCIDENT == true {
        add_man(win, y + 2, x + 43)
        add_man(win, y + 2, x + 47)
	}

	add_smoke(win, y - 1, x + D51.FUNNEL)
}

func add_C51(win *ncurses.Window, x int){
	// Get the size of the terminal
	TERM_LINES, _ := win.GetMaxYX() // _ = TERM_WIDTH
	var y = TERM_LINES / 2 - 5

	for i := 0; i < C51.HEIGHT; i++ {
		if i < 7 {
			my_MvAddStr(win, y + i, x, C51.head[i])
		}else {
			my_MvAddStr(win, y + i, x, C51.head_wheel[(C51.LENGTH + x) % C51.PATTERNS][i - 7])
		}

		my_MvAddStr(win, y + i, x + 55, C51.coal[i])
	}
	
	if ACCIDENT == true {
        add_man(win, y + 3, x + 45)
        add_man(win, y + 3, x + 49)
	}

	add_smoke(win, y - 1, x + C51.FUNNEL)
}

func add_man(win *ncurses.Window, y int, x int) {
	var man = [][]string{
		{
			"",
			"(O)",
		}, 

		{
			"Help!",
			"\\O/",
		},
	};

	for i := 0; i < 2; i++{
        my_MvAddStr(win, y + i, x, man[(LOGO.LENGTH + x) / 12 % 2][i]);
    }
}

func animation(){
	win := ncurses.Init()
	defer ncurses.EndWin()

	// Get the size of the terminal
	_, TERM_WIDTH := win.GetMaxYX() // _ = TERM_LINES

	// Set the cursor's visibility to 0
	_, _ = ncurses.CursSet(ncurses.CursorOff)

	for i := TERM_WIDTH - 1; ; i--{
		win.Erase()
		
		if L == true {
			if i + LOGO.LENGTH < 0 {
				break
			}
			add_LOGO(win, i)
		}else if C == true {
			if i + C51.LENGTH < 0 {
				break
			}
			add_C51(win, i)
		}else {
			if i + D51.LENGTH < 0 {
				break
			}
			add_D51(win, i)
		}

		win.Refresh()

		time.Sleep(50 * time.Millisecond) // 0.05s, fps = 20
	}

	win.Refresh()

	time.Sleep(100 * time.Millisecond)
}

// Main
func main() {
	animation()
}

/*

// ---------- NOTES ---------- //

sl source code: https://github.com/mtoyoda/sl/blob/master/sl.c

default car type: D51

!!! seehuhn.de/go/ncurses sucks :P !!!

	in: ~~\go\pkg\mod\seehuhn.de\go\ncurses@v0.2.0\keys.go

	line 171:
		Original:	C.KEY_EVENT:     KeyEvent,
		Fixed:		// C.KEY_EVENT:     KeyEvent,

*/
