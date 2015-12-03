package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func UserMiddleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c.Env["User"] = "GoLab"
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func Hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.Env["User"])
}

func main() {
	goji.Use(UserMiddleware)
	goji.Get("/hello", Hello)
	goji.Serve()
} //END
