package main 

import (
	"fmt"
	"github.com/fatih/color"
)

func request(url, username string) {
	finURL := url + username
	fmt.Println("Sending request to " + finURL)
}

func print(set, message string) {
	yellow := color.New(color.Bold, color.FgYellow).SprintFunc()
	red := color.New(color.Bold, color.FgRed).SprintFunc()
	blue := color.New(color.Bold, color.FgBlue).SprintFunc()
	switch set {
	case "prompt":
		fmt.Printf("%s %s\n", yellow("[PROMPT]"), message)
	case "error":
		fmt.Printf("%s %s\n", red("[ERROR]"), message)
	case "success":
		fmt.Printf("%s %s\n", blue("[SUCCESS]"), message)
	}
}

func logo(art string) {
	blue := color.New(color.Bold, color.FgBlue).PrintFunc()
	blue("\n", art)
}
