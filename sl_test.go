package main

import (
	"fmt"
	"seehuhn.de/go/ncurses" // https://pkg.go.dev/github.com/seehuhn/go-ncurses
)

// Smoke
func add_smoke(win Window, y int, x int) {
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

	var eraser = []string{
		"     ", "      ", "      ", "     ", "    ",
		"    " , "   "   , "   "   , "  "   , "  "  ,
		" "    , " "     , " "     , " "    , " "   ,
		" ",
	}

	dy []int = {  2,  1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 };
	dx []int = { -2, -1, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3 };

	// Unfinished
	
}

// Car
var HEIGHT, FUNNEL, LENGTH, PATTERNS int = 6, 4, 84, 6

var head = []string{
		"     ++      +------ ",
		"     ||      |+-+ |  ",
		"   /---------|| | |  ",
		"  + ========  +-+ |  ",
}

var head_wheel = [][]string{
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
}

var coal = []string{
	"____                 ",
	"|   \\@@@@@@@@@@@     ",
	"|    \\@@@@@@@@@@@@@_ ",
	"|                  | ",
	"|__________________| ",
	"   (O)       (O)     ",
}

var car = []string{
	"____________________ ",
	"|  ___ ___ ___ ___ | ",
	"|  |_| |_| |_| |_| | ",
	"|__________________| ",
	"|__________________| ",
	"   (O)        (O)    ",
}

var dell string = "                     "

func add_sl(win Window, x int) {
	// Get the size of the terminal
	TERM_LINES, TERM_WIDTH := ncurses.GetMaxYX()
	y = TERM_LINES / 2 - 3

	for i = 0; i <= HEIGHT; i++ {
		if i < 4 {
			win.MvAddStr(y + i, x, head[i])
		}
		else{
			win.MvAddStr(y + i, x, head_wheel[x % PATTERNS][i - 4])
		}

		win.MvAddStr(y + i, x + 21, coal[i])
		win.MvAddStr(y + i, x + 42, car[i])
		win.MvAddStr(y + i, x + 63, car[i])
    }

    add_smoke(y - 1, x + FUNNEL)
}

func animation(win Window){
	// Get the size of the terminal
	TERM_LINES, TERM_WIDTH := win.GetMaxYX()

	for i := TERM_WIDTH - 1; ; i--{
		add_sl(win, i)
	}
}

// Main
func main() {
	win := ncurses.Init()
	defer ncurses.EndWin()

	animation(win)
}
