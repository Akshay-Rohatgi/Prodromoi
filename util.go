package main 

import (
	"fmt"
	"github.com/fatih/color"
)


func print(set, message string) {
	yellow := color.New(color.Bold, color.FgYellow).SprintFunc()
	red := color.New(color.Bold, color.FgRed).SprintFunc()
	blue := color.New(color.Bold, color.FgBlue).SprintFunc()
	green := color.New(color.Bold, color.FgGreen).SprintFunc()
	white := color.New(color.Bold, color.FgWhite).SprintFunc()
	
	switch set {
	case "prompt":
		fmt.Printf("%s %s\n", yellow("[PROMPT]"), message)
	case "error":
		fmt.Printf("%s %s\n", red("[ERROR]"), message)
	case "success":
		fmt.Printf("%s %s\n", blue("[SUCCESS]"), message)
	case "running":
		fmt.Printf("%s %s\n", green("[RUNNING]"), message)
	case "info":
		fmt.Printf("%s %s\n", white("[INFO]"), message)
	}
}

func logo(art string) {
	blue := color.New(color.Bold, color.FgBlue).PrintFunc()
	blue("\n", art)
}
