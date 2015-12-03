package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// START OMIT
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	io.WriteString(w, "Hello "+ps.ByName("name"))
}

func main() {
	m := httprouter.New()
	m.GET("/:name", Hello)
	http.ListenAndServe(":8000", m)
}

// END OMIT
