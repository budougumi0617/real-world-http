package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	reader, err := strings.NewReader("テキスト")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
