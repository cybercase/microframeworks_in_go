package test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	http.HandleFunc("/hello", HelloGoLab)

	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/hello")
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if string(body) != "Hello GoLab!" {
		t.Fail()
	}
}
