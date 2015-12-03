package main

import (
	"fmt"
	"io"
	"net/http"
)

type GZippable interface {
	Flush()
}

type GZipResponseWriter struct {
	http.ResponseWriter
}

func (grw *GZipResponseWriter) Flush() {
	fmt.Println("Flushing")
}

func gZipBefore(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	fmt.Println("testMiddleware: before")
	return &GZipResponseWriter{ResponseWriter: w}, r
}

func gZipAfter(w http.ResponseWriter, r *http.Request) {
	grw, ok := w.(*GZipResponseWriter)
	if ok {
		grw.Flush()
	}
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, r *http.Request) {
	w, r = gZipBefore(w, r)
	defer gZipAfter(w, r)

	// w, r = sessionBefore(w, r)
	// defer sessionAfter(w, r)

	fmt.Println("Writing to rw")
	io.WriteString(w, "hello, world!\n")

}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8000", nil)
}
