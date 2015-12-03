package main

import (
	"io"
	"net/http"

	"github.com/bmizerany/pat"
)

// START OMIT
func Hello(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	io.WriteString(w, "Hello, "+params.Get(":name"))
}

func main() {
	m := pat.New()
	m.Get("/:name", http.HandlerFunc(Hello))
	http.ListenAndServe(":8000", m)
}

// END OMIT
