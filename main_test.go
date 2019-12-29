package main

import (
    "strings"
    "log"
    "os"
    "io/ioutil"
    "fmt"
    "testing"
    "net/http"
    "net/http/httptest"
)

func newTestServer() *httptest.Server {
    mux := http.NewServeMux()

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	file, err := os.Open("arlington-apartments.html")
	if err != nil {
	    log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	w.Write(b)
    })

    return httptest.NewServer(mux)
}

func TestServer(t *testing.T) {
    ts := newTestServer()
    defer ts.Close()

    fmt.Printf("%s\n", ts.URL)
    resp, _ := http.Get(ts.URL)
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    if !strings.HasPrefix(string(body), "<!DOCTYPE html>") {
	t.Error("Root html file not valid")
    }
}

func TestApartmentNames(t *testing.T) {
    ts := newTestServer()
    defer ts.Close()

    eq := EquityQuery{searchUrl: ts.URL}

    ApartmentNames(eq)
}
