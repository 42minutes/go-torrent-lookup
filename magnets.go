package torrentlookup

import (
	"fmt"
	"net/url"
)

var trackers = []string{
	"udp://open.demonii.com:1337/announce",
	"udp://tracker.publicbt.com:80/announce",
	"udp://tracker.openbittorrent.com:80/announce",
	"udp://tracker.istole.it:80",
	"http://www.eddie4.nl:6969/announce",
	"http://tracker.nwps.ws:6969/announce",
	"http://bigfoot1942.sektori.org:6969/announce",
	"http://9.rarbg.com:2710/announce",
	"http://torrent-tracker.ru:80/announce.php",
	"http://bttracker.crunchbanglinux.org:6969/announce",
	"http://explodie.org:6969/announce",
	"http://tracker.tfile.me/announce",
	"http://tracker.best-torrents.net:6969/announce",
	"http://tracker1.wasabii.com.tw:6969/announce",
	"http://bt.careland.com.cn:6969/announce",
}

// CreateFakeMagnet will create a magnet link using a given hash and a predefined list
// of trackets that may or may not have this magnet.
func CreateFakeMagnet(infohash string) string {
	magnetURL := fmt.Sprintf("magnet:?xt=urn:btih:%s", infohash)
	for _, tracker := range trackers {
		magnetURL += fmt.Sprintf("&tr=%s", url.QueryEscape(tracker))
	}
	return magnetURL
}
