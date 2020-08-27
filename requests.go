package main

import (
	"net/http"
)

func request(url, username string) (*http.Response) {
	finURL := "https://" + url + username
	print("running", "Sending request to " + finURL)

	resp, err := http.Get(finURL)
    if err != nil {
        print("error", "There was an error sending the GET request")
	}
	defer resp.Body.Close()
	return resp
}

func validate(url, username string) (bool) {
	bruh := request("github.com/", "Akshay-Rohatgi")
	if bruh.StatusCode == 200 {
		return true
	}
	return false
}