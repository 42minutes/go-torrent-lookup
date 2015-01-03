package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
)

type Result struct {
	Name   string
	Manget map[string]string
}

func main() {
	term := "Big Bang Theory"
	searchUrl := fmt.Sprintf("https://torrentz.eu/search?q=%s", url.QueryEscape(term))
	fmt.Println("Parsing ", searchUrl)
	doc, err := goquery.NewDocument(searchUrl)
	if err == nil {
		doc.Find(".results dl dt").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Find("a").Attr("href")
			name := s.Find("a").Text()
			fmt.Println(name)
			results := ListResultPages("https://torrentz.eu" + link)
			for site, link := range results {
				fmt.Println(site)
				FindMagnet(link)
			}
		})
	}
}

func ListResultPages(url string) map[string]string {
	results := make(map[string]string)
	doc, err := goquery.NewDocument(url)
	if err == nil {
		doc.Find(".download dl").Each(func(i int, s *goquery.Selection) {
			dd := s.Find("dd").Text()
			if dd != "Sponsored Link" {
				link, _ := s.Find("dt a").Attr("href")
				name := s.Find("dt a span.u").Text()
				results[name] = link
				// fmt.Println(name, link)
			}
		})
	}
	return results
}

func FindMagnet(url string) map[string]string {
	doc, err := goquery.NewDocument(url)
	if err == nil {
		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			// fmt.Println(link)
			if strings.Contains(string(link), "magnet:") {
				fmt.Println(link)
			}
		})
	}
	return nil
}
