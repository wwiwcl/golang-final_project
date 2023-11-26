package main

import (
	"fmt"
	"os"
)

func main() {
	var cmd string
	fmt.Scanln(&cmd)
	str, _ := os.Getwd()
	fmt.Println(str)
}
