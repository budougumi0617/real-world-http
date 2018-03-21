package main

import (
	"log"
	"net/http"
)

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	http.HandleFunc("/cookie", cookieHandler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
