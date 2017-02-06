package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandleRequest(t *testing.T) {
    cases := []struct {
        w *httptest.ResponseRecorder
        r *http.Request
        want string
    }{
        {httptest.NewRecorder(), httptest.NewRequest("GET", "http://localhost:8080/culo", nil), "Hi there, Caca culo!"},
        {httptest.NewRecorder(), httptest.NewRequest("GET", "http://localhost:8080/5", nil), "Hi there, Caca 5!"},
        {httptest.NewRecorder(), httptest.NewRequest("GET", "http://localhost:8080/", nil), "Hi there, Caca !"},
    }

    for _, c := range cases {
        HandleRequest(c.w, c.r)
        if c.w.Body.String() != c.want || c.w.Code != 200 {
            t.Errorf("Handle(%q) => %q, want %q", c.r.URL, c.w.Body.String(), c.want)
        }
    }
}

