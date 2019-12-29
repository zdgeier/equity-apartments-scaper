package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type EquityQuery struct {
    searchUrl string
}

func ApartmentNames(e EquityQuery) {
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

	c.Visit(e.searchUrl)
}

func main() {
    myquery := EquityQuery{searchUrl: "https://www.equityapartments.com/arlington-apartments"}
    ApartmentNames(myquery)
}
