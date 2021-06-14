package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonvat.com/")
	if err != nil {
		log.Fatal("The request failed!")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("failed to read body: ", err)
	}

	fmt.Println(string(bytes))
}
