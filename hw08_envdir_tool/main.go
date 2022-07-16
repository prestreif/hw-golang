package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		_ = fmt.Errorf("Error: %v\r\n", "need two or more arguments")
	}

	envs, err := ReadDir(os.Args[1])
	if err != nil {
		_ = fmt.Errorf("Error: %v\r\n", err)
	}

	os.Exit(RunCmd(os.Args[2:], envs))
}
