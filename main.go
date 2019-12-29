package main

import (
	"strings"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/PuerkitoBio/goquery"
)

type EquityQuery struct {
    baseUrl string
    searchUrl string
}

func Apartments(e EquityQuery) {
	urls := ApartmentURLs(e)

	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("#bedroom-type-2", func(e *colly.HTMLElement) {
		e.DOM.Find(".row.unit").Each(func(i int, s *goquery.Selection) {
			available := strings.TrimSpace(s.Find("p:contains('Available')").Text())
			floor := strings.TrimSpace(s.Find("span:contains('Floor')").Text())
			sqft := strings.TrimSpace(s.Find("span:contains('sq.')").Text())
			price := s.Find(".pricing").Text()
			fmt.Println("Price " + price + " " + available + " " + sqft + " " + floor) //+ date.Text())
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	for _, url := range urls {
	    c.Visit(e.baseUrl + url)
	}
}

func ApartmentURLs(e EquityQuery) []string {
    // Instantiate default collector
    c := colly.NewCollector()
    var urls []string

    // On every a element which has href attribute call callback
    c.OnHTML(".property", func(e *colly.HTMLElement) {
	name := e.DOM.Find("h3[itemprop='name']")
	link, _ := name.Find("a[href]").Attr("href")
	urls = append(urls, link)
    })

    // Before making a request print "Visiting ..."
    c.OnRequest(func(r *colly.Request) {
	    fmt.Println("Visiting", r.URL.String())
    })

    c.Visit(e.searchUrl)

    return urls
}

func ApartmentNames(e EquityQuery) []string {
    // Instantiate default collector
    c := colly.NewCollector()
    var l []string

    // On every a element which has href attribute call callback
    c.OnHTML(".property", func(e *colly.HTMLElement) {
	name := e.DOM.Find("h3[itemprop='name']")
	l = append(l, strings.TrimSpace(name.Text()))
    })

    c.Visit(e.searchUrl)

    return l
}

func main() {
    myquery := EquityQuery{
	baseUrl: "https://www.equityapartments.com",
	searchUrl: "https://www.equityapartments.com/arlington-apartments",
    }
    Apartments(myquery)
}
