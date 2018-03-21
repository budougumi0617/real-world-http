package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func cookieHandler(w http.ResponseWriter, req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))

	// Add cookie
	c := &http.Cookie{
		Name:  "cookie_name",
		Value: "cookie_value",
	}
	http.SetCookie(w, c)
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}
