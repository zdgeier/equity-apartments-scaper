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

var cont struct {
    Serv *httptest.Server
}

func TestMain(m *testing.M) {
    cont.Serv = newTestServer()
    defer cont.Serv.Close()

    os.Exit(m.Run())
}

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
    fmt.Printf("%s\n", cont.Serv.URL)
    resp, _ := http.Get(cont.Serv.URL)
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    if !strings.HasPrefix(string(body), "<!DOCTYPE html>") {
	t.Error("Root html file not valid")
    }
}

func TestApartments(t *testing.T) {
    // TODO: Write test
}

func TestApartmentURLs(t *testing.T) {
    expected := [...]string {
	"/arlington/rosslyn/1800-oak-apartments",
	"/arlington/pentagon-city/1401-joyce-on-pentagon-row-apartments",
	"/arlington/courthouse/2201-pershing-apartments",
	"/arlington/virginia-square/virginia-square-apartments",
	"/arlington/courthouse/sheffield-court-apartments",
	"/arlington/clarendon/the-reserve-at-clarendon-centre-apartments",
	"/arlington/courthouse/the-prime-at-arlington-courthouse-apartments",
	"/arlington/ballston/liberty-tower-apartments",
	"/arlington/clarendon/the-clarendon-apartments",
	"/arlington/columbia-pike/columbia-crossing-apartments",
	"/arlington/courthouse/courthouse-plaza-apartments",
	"/arlington/crystal-city/water-park-towers-apartments",
	"/arlington/courthouse/2201-wilson-apartments",
	"/arlington/pentagon-city/lofts-590-apartments",
	"/arlington/crystal-city/crystal-place-apartments",
    }

    eq := EquityQuery{searchUrl: cont.Serv.URL}

    urls := ApartmentURLs(eq)

    actual := [15]string{}
    copy(actual[:], urls)


    if actual != expected {
	t.Errorf("Error actual = %v, and Expected = %v.", actual, expected)
    }
}

func TestApartmentNames(t *testing.T) {
    eq := EquityQuery{searchUrl: cont.Serv.URL}

    names := ApartmentNames(eq)

    actual := [15]string{}
    copy(actual[:], names)

    expected := [...]string {
	"1800 Oak Apartments",
	"1401 Joyce on Pentagon Row Apartments",
	"2201 Pershing Apartments",
	"Virginia Square Apartments",
	"Sheffield Court Apartments",
	"The Reserve at Clarendon Centre Apartments",
	"The Prime at Arlington Courthouse Apartments",
	"Liberty Tower Apartments",
	"The Clarendon Apartments",
	"Columbia Crossing Apartments",
	"Courthouse Plaza Apartments",
	"Water Park Towers Apartments",
	"2201 Wilson Apartments",
	"Lofts 590 Apartments",
	"Crystal Place Apartments",
    }

    if actual != expected {
	t.Errorf("Error actual = %v, and Expected = %v.", actual, expected)
    }
}
