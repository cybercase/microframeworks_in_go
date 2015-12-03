package main

import "io"
import "net/http"

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Golab!")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/hello", &MyHandler{})

	server := http.Server{Addr: ":8000", Handler: mux}
	server.ListenAndServe()
}
