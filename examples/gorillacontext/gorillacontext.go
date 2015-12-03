package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
)

func UserMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "User", "GoLab")
		h.ServeHTTP(w, r)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name, _ := context.GetOk(r, "User")
	fmt.Fprintf(w, "Hello "+name.(string))
}

func main() {
	r := http.NewServeMux()
	r.Handle("/hello", MyMiddleware(http.HandlerFunc(Hello)))
	http.ListenAndServe(":8000", context.ClearHandler(r))
} //END
