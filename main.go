package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Invalid args. Please pass one file name.")
		return
	}

	cmd := exec.Command("tree-sitter", "highlight", os.Args[1])
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}
