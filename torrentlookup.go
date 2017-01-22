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

// Torrent -
type Torrent struct {
	Name     string
	Infohash string
}

// Search allows finding magnet links in the provider
func (provider *Provider) Search(query string) ([]*Torrent, error) {
	results := []*Torrent{}
	searchURL := fmt.Sprintf(provider.SearchURL, url.QueryEscape(query))
	doc, err := goquery.NewDocument(searchURL)
	if err != nil {
		return nil, err
	}
	doc.Find(provider.RowQuery).Each(func(i int, s *goquery.Selection) {
		var infohash, name string
		seeds, _ := strconv.Atoi(s.Find(provider.SeedsSubQuery).First().Text())
		name = s.Find(provider.NameSubQuery).First().Text()
		magnet, _ := s.Find(provider.MagnetSubQuery).First().Attr("href")
		if magnet != "" {
			infohash = getInfohashFromMagnet(magnet)
		}
		if seeds > 0 && infohash != "" {
			tor := Torrent{
				Name:     name,
				Infohash: infohash,
			}
			results = append(results, &tor)
		}

	})
	return results, nil
}

func getInfohashFromMagnet(magnet string) (infohash string) {
	re := regexp.MustCompile("([a-zA-Z0-9]{40})")
	// TODO Check for errors
	// TODO Need a beter regex
	return re.FindString(magnet)
}
