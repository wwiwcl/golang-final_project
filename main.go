package main
import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	str, _ := os.Getwd()
	fmt.Println(str)
	cmd := exec.Command("ls", "-l")
	out, _ := cmd.Output()
	fmt.Println(string(out))
}