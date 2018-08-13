package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var image []byte

// Preset image file
func init() {
	var err error
	image, err = ioutil.ReadFile("./image.png")
	if err != nil {
		panic(err)
	}
}

// Send HTML to browser
// Push image
func handlerHTML(w http.ResponseWriter, r *http.Request) {
	// Push if be able to cast http.Pusher.
	pusher, ok := w.(http.Pusher)
	if ok {
		log.Println("Connected HTTTP2")
		pusher.Push("/image", nil)
	}
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><body><img src="/image"></body></html>`)
}

// Send image gile to the browser
func handlerImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Write(image)
}

func main() {
	http.HandleFunc("/", handlerHTML)
	http.HandleFunc("/image", handlerImage)
	fmt.Println("start http listening :18433")
	err := http.ListenAndServeTLS(":18433", "../../ch06/cert/server.crt", "../../ch06/cert/server.key", nil)
	fmt.Println(err)
}
