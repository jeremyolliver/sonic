package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/text"
	"strings"
)

func InfoCommand(instanceidentifier string, fullOutput bool) {
	fmt.Println("Searching for: " + instanceidentifier + "...")

	if strings.HasPrefix(instanceidentifier, "i-") {
		DescribeEC2Instance(instanceidentifier, fullOutput)
	} else if strings.HasPrefix(instanceidentifier, "mi-") {
		fmt.Println(text.Colors{text.FgRed}.Sprint("TODO: querying managed instances via SSM is not yet supported"))
	} else {
		fmt.Println(text.Colors{text.FgRed}.Sprint("Error: Unsupported query format: " + instanceidentifier))
	}
}
