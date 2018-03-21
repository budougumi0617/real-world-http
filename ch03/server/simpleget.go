package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	for k, v := range req.Header {
		fmt.Println(k, v)
	}
	fmt.Printf("HEAD elements= %+v\n", req.Header)
	b, _ := ioutil.ReadAll(req.Body)
	fmt.Printf("BODY = %+v\n", string(b))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}
