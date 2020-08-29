package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Sites struct {
	Site string
	Host string
}

func decode() ([]Sites) {
	file, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println("error:", err)
	}
	var urls []Sites
	err1 := json.Unmarshal(file, &urls)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	return urls
}
