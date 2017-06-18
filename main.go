package main

import (
	"fmt"
	"os"

	"github.com/nrocco/tpm/cmd"
)

var (
	VERSION = "undefined"
)

func main() {
	if err := cmd.Execute(VERSION); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
