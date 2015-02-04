package torrentlookup

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var trackers []string = []string{
	"udp://open.demonii.com:1337/announce",
	"udp://tracker.publicbt.com:80/announce",
	"udp://tracker.openbittorrent.com:80/announce",
	"udp://tracker.istole.it:80",
	// "http://www.eddie4.nl:6969/announce",
	// "http://tracker.nwps.ws:6969/announce",
	// "http://bigfoot1942.sektori.org:6969/announce",
	// "http://9.rarbg.com:2710/announce",
	// "http://torrent-tracker.ru:80/announce.php",
	// "http://bttracker.crunchbanglinux.org:6969/announce",
	// "http://explodie.org:6969/announce",
	// "http://tracker.tfile.me/announce",
	// "http://tracker.best-torrents.net:6969/announce",
	// "http://tracker1.wasabii.com.tw:6969/announce",
	// "http://bt.careland.com.cn:6969/announce",
}

type Provider struct {
	Name        string
	SearchUrl   string
	NameQuery   string
	MagnetQuery string
	SeedsQuery  string
	Crawl       bool
}

var providers map[string]Provider = map[string]Provider{
	"kickass": {Name: "Kickass", SearchUrl: "https://kickass.so/usearch/%s", NameQuery: "td:nth-child(1) .torrentname div a", MagnetQuery: "td:nth-child(1) a.imagnet", SeedsQuery: "td:nth-child(5)", Crawl: false},
	// "torrentz": {Name: "Torrentz.eu", SearchUrl: "https://torrentz.eu/verified?f=%s", NameQuery: ".results dl dt a", MagnetQuery: ".results dl dt a", Crawl: false},
}

func SearchProvider(query string, providerKey string) (name, infohash string) {
	provider := providers[providerKey]
	searchUrl := fmt.Sprintf(provider.SearchUrl, url.QueryEscape(query))
	doc, err := goquery.NewDocument(searchUrl)
	if err == nil {
		seedsString := doc.Find(provider.SeedsQuery).First().Text()
		seeds, _ := strconv.Atoi(seedsString)
		name = doc.Find(provider.NameQuery).First().Text()
		magnet, _ := doc.Find(provider.MagnetQuery).First().Attr("href")
		if magnet != "" {
			infohash = getInfohashFromMagnet(magnet)
		}
		if seeds == 0 {
			name = ""
			infohash = ""
		}
	}
	return name, infohash
}

// TODO Allow Search to return multiple results for us to be able to check season/episode
func Search(query string) (name, infohash string) {
	for providerKey, _ := range providers {
		name, infohash = SearchProvider(query, providerKey)
	}
	return name, infohash
}

func getInfohashFromMagnet(magnet string) (infohash string) {
	re := regexp.MustCompile("([a-zA-Z0-9]{40})")
	// TODO Check for errors
	// TODO Need a beter regex
	return re.FindString(magnet)
}

func listResultPages(url string) map[string]string {
	results := make(map[string]string)
	doc, err := goquery.NewDocument(url)
	if err == nil {
		doc.Find(".download dl").Each(func(i int, s *goquery.Selection) {
			dd := s.Find("dd").Text()
			if dd != "Sponsored Link" {
				link, _ := s.Find("dt a").Attr("href")
				name := s.Find("dt a span.u").Text()
				results[name] = link
			}
		})
	}
	return results
}

func findMagnets(url string) []string {
	magnets := make([]string, 0)
	doc, err := goquery.NewDocument(url)
	if err == nil {
		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if strings.Contains(string(link), "magnet:") {
				magnets = append(magnets, link)
			}
		})
	}
	return nil
}

func FakeMagnet(infohash string) string {
	var magnetUrl string = fmt.Sprintf("magnet:?xt=urn:btih:%s", infohash)
	for _, tracker := range trackers {
		magnetUrl += fmt.Sprintf("&tr=%s", url.QueryEscape(tracker))
	}
	return magnetUrl
}
