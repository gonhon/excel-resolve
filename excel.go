package main

import (
	"fmt"
	"os"

	"github.com/gonhon/excel-resolve/cmd"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}()
	cmd.Execute()
}
