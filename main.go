package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/webdav"
	"net/http"
)

func main() {
	var address string
	flag.StringVar(&address, "a", "localhost:8080", "Address to listen to.")
	flag.Parse()

	handler := &webdav.Handler{
		FileSystem: webdav.Dir("."),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, e error) {
			fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		},
	}

	http.ListenAndServe(address, handler)
}
