package main

import (
	"os"
)

func main() {

	logo(`
██▓███   ██▀███   ▒█████  ▓█████▄  ██▀███   ▒█████   ███▄ ▄███▓ ▒█████   ██▓
▓██░  ██▒▓██ ▒ ██▒▒██▒  ██▒▒██▀ ██▌▓██ ▒ ██▒▒██▒  ██▒▓██▒▀█▀ ██▒▒██▒  ██▒▓██▒
▓██░ ██▓▒▓██ ░▄█ ▒▒██░  ██▒░██   █▌▓██ ░▄█ ▒▒██░  ██▒▓██    ▓██░▒██░  ██▒▒██▒
▒██▄█▓▒ ▒▒██▀▀█▄  ▒██   ██░░▓█▄   ▌▒██▀▀█▄  ▒██   ██░▒██    ▒██ ▒██   ██░░██░
▒██▒ ░  ░░██▓ ▒██▒░ ████▓▒░░▒████▓ ░██▓ ▒██▒░ ████▓▒░▒██▒   ░██▒░ ████▓▒░░██░
▒▓▒░ ░  ░░ ▒▓ ░▒▓░░ ▒░▒░▒░  ▒▒▓  ▒ ░ ▒▓ ░▒▓░░ ▒░▒░▒░ ░ ▒░   ░  ░░ ▒░▒░▒░ ░▓  
░▒ ░       ░▒ ░ ▒░  ░ ▒ ▒░  ░ ▒  ▒   ░▒ ░ ▒░  ░ ▒ ▒░ ░  ░      ░  ░ ▒ ▒░  ▒ ░
░░         ░░   ░ ░ ░ ░ ▒   ░ ░  ░   ░░   ░ ░ ░ ░ ▒  ░      ░   ░ ░ ░ ▒   ▒ ░
			░         ░ ░     ░       ░         ░ ░         ░       ░ ░   ░  
							░                                                																						  
`)
	args := os.Args
	args = args[1:]
	if len(args) < 1 {
		print("error", "No username specified")
		os.Exit(1)
	}

	username := args[0]
	for _, siteInfo := range decode() {
		print("running", "Checking for user " + username + " at " + siteInfo.Site)
		check := validate(siteInfo.Host, username)
		if check {
			print("success", "Found a user with the username " + username + " at " + siteInfo.Site)
		} else {
			print("error", "Did not find a user with the username " + username + " at " + siteInfo.Site)
		}
	}

}
