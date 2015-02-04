package torrentlookup

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
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

func Search(term string, deepCrawl bool) {
	searchUrl := fmt.Sprintf("https://torrentz.eu/verified?f=%s", url.QueryEscape(term))
	fmt.Println("Parsing ", searchUrl)

	doc, err := goquery.NewDocument(searchUrl)
	if err == nil {
		doc.Find(".results dl dt").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Find("a").Attr("href")
			name := s.Find("a").Text()
			fmt.Println(name)
			if deepCrawl == true {
				results := listResultPages("https://torrentz.eu" + link)
				for site, link := range results {
					fmt.Println(site)
					findMagnet(link)
				}
			} else {
				infohash := strings.Trim(link, "/")
				magnetUrl := fakeMagnet(infohash)
				fmt.Println(magnetUrl)
			}
		})
	}
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
				// fmt.Println(name, link)
			}
		})
	}
	return results
}

func findMagnet(url string) map[string]string {
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

func fakeMagnet(infohash string) string {
	var magnetUrl string = fmt.Sprintf("magnet:?xt=urn:btih:%s", infohash)
	for _, tracker := range trackers {
		magnetUrl += fmt.Sprintf("&tr=%s", url.QueryEscape(tracker))
	}
	return magnetUrl
}
