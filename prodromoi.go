package main

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
	
	print("error", "bruh that didnt work")
	print("prompt", "type something")
	print("success", "yay something worked for once!")
	print("running", "something is running")
	print("info", "important information")

	if validate("github.com/", "Akshay-Rohatgi") == true {
		print("success", "username found!")
	}
}