package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// Get status by string
	log.Println(resp.Status)
	// Get status by number
	log.Println(resp.StatusCode)
	// Get headers
	log.Println("Headers:", resp.Header)
	// Get specified element from header
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
	log.Println(string(body))
}
