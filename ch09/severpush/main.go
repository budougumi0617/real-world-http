package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var image []byte

// Preset image file
func init() {
	var err error
	image, err = ioutil.ReadFile("./image.ong")
	if err != nil {
		panic(err)
	}
}

// Send HTML to browser
// Push image
func handlerHTML(w http.ResponseWriter, r *http.Request) {
	pusher, ok := w.(http.Pusher)
	if ok {
		pusher.Push("/image", nil)
	}
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><body><img src="/image"></body></html>`)
}

func main() {

}
