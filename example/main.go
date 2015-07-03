package main

import (
	"fmt"
	"github.com/northbright/basicauth"
	"io"
	"net/http"
)

type MyHandler struct{}

var (
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
)

func hello(w http.ResponseWriter, r *http.Request) {
	// Use default argument(default Basic realm string).
	ba := basicauth.New("admin", "admin")
	if ba.IsOK(w, r) {
		io.WriteString(w, "/: you are in.")
	} else {
		io.WriteString(w, "/: 401: unauthorized.")
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	// Use customized argument(Basis realm string)
	ba := basicauth.NewWithArgs("john", "123456", basicauth.Arguments{"Yo!"})
	if ba.IsOK(w, r) {
		io.WriteString(w, "/welcome: you're in.")
	} else {
		io.WriteString(w, "/welcome: 401: unauthorized.")
	}
}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	if h, ok := mux[url]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "URL can not found: "+r.URL.String())
}

func main() {
	mux["/"] = hello
	mux["/welcome"] = welcome

	server := http.Server{
		Addr:    ":8080",
		Handler: &MyHandler{},
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("ListenAndServe: %v", err)
	}

	// Output:
}
