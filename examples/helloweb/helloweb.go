package main

import "io"
import "net/http"

func HelloGoLab(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello GoLab!\n")
}

func main() {
	http.HandleFunc("/hello", HelloGoLab) //HandlerFunc(HelloGoLab) and attach to DefaultServeMux
	http.ListenAndServe(":8000", nil)     //New server around DefaultServeMux if 2nd param is nil
}
