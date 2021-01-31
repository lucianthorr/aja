package main

import (
	"fmt"

	"github.com/lucianthorr/aja/cmd"
)

func main() {

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
