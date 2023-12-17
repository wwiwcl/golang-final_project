/*source: https://github.com/mtoyoda/sl*/

package sl

import (
	"fmt"
	"os/exec"
)

func Sl(c *exec.Cmd, args ...string) error {
	fmt.Println("sl has not done yet") // TODO
	return nil
}

const D51HEIGHT int = 10
const D51FUNNEL int = 7
const D51LENGTH int = 83
const D51PATTERNS int = 6

const D51STR1 string = "      ====        ________                ___________ "
const D51STR2 string = "  _D _|  |_______/        \\__I_I_____===__|_________| "
const D51STR3 string = "   |(_)---  |   H\\________/ |   |        =|___ ___|   "
const D51STR4 string = "   /     |  |   H  |  |     |   |         ||_| |_||   "
const D51STR5 string = "  |      |  |   H  |__--------------------| [___] |   "
const D51STR6 string = "  | ________|___H__/__|_____/[][]~\\_______|       |   "
const D51STR7 string = "  |/ |   |-----------I_____I [][] []  D   |=======|__ "

const D51WHL11 string = "__/ =| o |=-~~\\  /~~\\  /~~\\  /~~\\ ____Y___________|__ "
const D51WHL12 string = " |/-=|___|=    ||    ||    ||    |_____/~\\___/        "
const D51WHL13 string = "  \\_/      \\O=====O=====O=====O_/      \\_/            "

const D51WHL21 string = "__/ =| o |=-~~\\  /~~\\  /~~\\  /~~\\ ____Y___________|__ "
const D51WHL22 string = " |/-=|___|=O=====O=====O=====O   |_____/~\\___/        "
const D51WHL23 string = "  \\_/      \\__/  \\__/  \\__/  \\__/      \\_/            "

const D51WHL31 string = "__/ =| o |=-O=====O=====O=====O \\ ____Y___________|__ "
const D51WHL32 string = " |/-=|___|=    ||    ||    ||    |_____/~\\___/        "
const D51WHL33 string = "  \\_/      \\__/  \\__/  \\__/  \\__/      \\_/            "

const D51WHL41 string = "__/ =| o |=-~O=====O=====O=====O\\ ____Y___________|__ "
const D51WHL42 string = " |/-=|___|=    ||    ||    ||    |_____/~\\___/        "
const D51WHL43 string = "  \\_/      \\__/  \\__/  \\__/  \\__/      \\_/            "

const D51WHL51 string = "__/ =| o |=-~~\\  /~~\\  /~~\\  /~~\\ ____Y___________|__ "
const D51WHL52 string = " |/-=|___|=   O=====O=====O=====O|_____/~\\___/        "
const D51WHL53 string = "  \\_/      \\__/  \\__/  \\__/  \\__/      \\_/            "

const D51WHL61 string = "__/ =| o |=-~~\\  /~~\\  /~~\\  /~~\\ ____Y___________|__ "
const D51WHL62 string = " |/-=|___|=    ||    ||    ||    |_____/~\\___/        "
const D51WHL63 string = "  \\_/      \\_O=====O=====O=====O/      \\_/            "

const D51DEL string = "                                                      "

const COAL01 string = "                              "
const COAL02 string = "                              "
const COAL03 string = "    _________________         "
const COAL04 string = "   _|                \\_____A  "
const COAL05 string = " =|                        |  "
const COAL06 string = " -|                        |  "
const COAL07 string = "__|________________________|_ "
const COAL08 string = "|__________________________|_ "
const COAL09 string = "   |_D__D__D_|  |_D__D__D_|   "
const COAL10 string = "    \\_/   \\_/    \\_/   \\_/    "

const COALDEL string = "                              "

const LOGOHEIGHT int = 6
const LOGOFUNNEL int = 4
const LOGOLENGTH int = 84
const LOGOPATTERNS int = 6

const LOGO1 string = "     ++      +------ "
const LOGO2 string = "     ||      |+-+ |  "
const LOGO3 string = "   /---------|| | |  "
const LOGO4 string = "  + ========  +-+ |  "

const LWHL11 string = " _|--O========O~\\-+  "
const LWHL12 string = "//// \\_/      \\_/    "

const LWHL21 string = " _|--/O========O\\-+  "
const LWHL22 string = "//// \\_/      \\_/    "

const LWHL31 string = " _|--/~O========O-+  "
const LWHL32 string = "//// \\_/      \\_/    "

const LWHL41 string = " _|--/~\\------/~\\-+  "
const LWHL42 string = "//// \\_O========O    "

const LWHL51 string = " _|--/~\\------/~\\-+  "
const LWHL52 string = "//// \\O========O/    "

const LWHL61 string = " _|--/~\\------/~\\-+  "
const LWHL62 string = "//// O========O_/    "

const LCOAL1 string = "____                 "
const LCOAL2 string = "|   \\@@@@@@@@@@@     "
const LCOAL3 string = "|    \\@@@@@@@@@@@@@_ "
const LCOAL4 string = "|                  | "
const LCOAL5 string = "|__________________| "
const LCOAL6 string = "   (O)       (O)     "

const LCAR1 string = "____________________ "
const LCAR2 string = "|  ___ ___ ___ ___ | "
const LCAR3 string = "|  |_| |_| |_| |_| | "
const LCAR4 string = "|__________________| "
const LCAR5 string = "|__________________| "
const LCAR6 string = "   (O)        (O)    "

const DELLN string = "                     "

const C51HEIGHT int = 11
const C51FUNNEL int = 7
const C51LENGTH int = 87
const C51PATTERNS int = 6

const C51DEL string = "                                                       "

const C51STR1 string = "        ___                                            "
const C51STR2 string = "       _|_|_  _     __       __             ___________"
const C51STR3 string = "    D__/   \\_(_)___|  |__H__|  |_____I_Ii_()|_________|"
const C51STR4 string = "     | `---'   |:: `--'  H  `--'         |  |___ ___|  "
const C51STR5 string = "    +|~~~~~~~~++::~~~~~~~H~~+=====+~~~~~~|~~||_| |_||  "
const C51STR6 string = "    ||        | ::       H  +=====+      |  |::  ...|  "
const C51STR7 string = "|    | _______|_::-----------------[][]-----|       |  "

const C51WH61 string = "| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__"
const C51WH62 string = "------'|oOo|==[]=-     ||      ||      |  ||=======_|__"
const C51WH63 string = "/~\\____|___|/~\\_|   O=======O=======O  |__|+-/~\\_|     "
const C51WH64 string = "\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       "

const C51WH51 string = "| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__"
const C51WH52 string = "------'|oOo|===[]=-    ||      ||      |  ||=======_|__"
const C51WH53 string = "/~\\____|___|/~\\_|    O=======O=======O |__|+-/~\\_|     "
const C51WH54 string = "\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       "

const C51WH41 string = "| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__"
const C51WH42 string = "------'|oOo|===[]=- O=======O=======O  |  ||=======_|__"
const C51WH43 string = "/~\\____|___|/~\\_|      ||      ||      |__|+-/~\\_|     "
const C51WH44 string = "\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       "

const C51WH31 string = "| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__"
const C51WH32 string = "------'|oOo|==[]=- O=======O=======O   |  ||=======_|__"
const C51WH33 string = "/~\\____|___|/~\\_|      ||      ||      |__|+-/~\\_|     "
const C51WH34 string = "\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       "

const C51WH21 string = "| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__"
const C51WH22 string = "------'|oOo|=[]=- O=======O=======O    |  ||=======_|__"
const C51WH23 string = "/~\\____|___|/~\\_|      ||      ||      |__|+-/~\\_|     "
const C51WH24 string = "\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       "

const C51WH11 string = "| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__"
const C51WH12 string = "------'|oOo|=[]=-      ||      ||      |  ||=======_|__"
const C51WH13 string = "/~\\____|___|/~\\_|  O=======O=======O   |__|+-/~\\_|     "
const C51WH14 string = "\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       "

const (
	Cols          = 80
	Rows          = 24
	LogoPatterns  = 6
	LogoHeight    = 6
	LogoLength    = 80
	D51Patterns   = 6
	D51Height     = 10
	D51Length     = 80
	C51Patterns   = 6
	C51Height     = 11
	C51Length     = 80
	SmokePatterns = 16
)
