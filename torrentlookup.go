package torrentlookup

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// Provider -
type Provider struct {
	Name           string
	SearchURL      string
	RowQuery       string
	NameSubQuery   string
	MagnetSubQuery string
	SeedsSubQuery  string
}

// Search allows finding magnet links in the provider
func (provider *Provider) Search(query string) (name, infohash string, err error) {
	searchURL := fmt.Sprintf(provider.SearchURL, url.QueryEscape(query))
	doc, err := goquery.NewDocument(searchURL)
	if err != nil {
		return
	}
	doc.Find(provider.RowQuery).EachWithBreak(func(i int, s *goquery.Selection) bool {
		seeds, _ := strconv.Atoi(s.Find(provider.SeedsSubQuery).First().Text())
		name = s.Find(provider.NameSubQuery).First().Text()
		magnet, _ := s.Find(provider.MagnetSubQuery).First().Attr("href")
		if magnet != "" {
			infohash = getInfohashFromMagnet(magnet)
		}
		if seeds > 0 && infohash != "" {
			return false
		}
		return true
	})
	return name, infohash, nil
}

func getInfohashFromMagnet(magnet string) (infohash string) {
	re := regexp.MustCompile("([a-zA-Z0-9]{40})")
	// TODO Check for errors
	// TODO Need a beter regex
	return re.FindString(magnet)
}
