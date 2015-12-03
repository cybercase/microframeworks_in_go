package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloGoLab(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	HelloGoLab(w, r)

	if w.Body.String() != "Hello GoLab!" {
		t.Fail()
	}
}
