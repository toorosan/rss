package rssreader

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html/charset"
)

// parser describes parser for certain RSS feed version.
type parser interface {
	Version() rssVerion
	// Parse prepares array of RSSItems from blob XML data.
	Parse() ([]RssItem, error)
}

// getter describes interface for some client allowing to get web page contents by it's url.
type getter interface {
	Get(u url.URL) ([]byte, error)
}

func newParser(u url.URL, g getter) (parser, error) {
	blobXML, err := g.Get(u)
	if err != nil {
		return nil, fmt.Errorf("failed to get blob data from url %q", u.String())
	}

	version, err := getRSSVersion(blobXML)
	if err != nil {
		return nil, fmt.Errorf("failed to get rss version: %v", err)
	}

	switch version {
	case rss20:
		return &rss20Parser{blob: blobXML}, nil
	default:
		return nil, fmt.Errorf("rss version %q is not supported", version)
	}
}

type webGetter struct{}

func (w *webGetter) Get(u url.URL) ([]byte, error) {
	var resp *http.Response
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get")
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// parseXML is used instead of regular `xml.Unmarshal`
// to fix XML decoding error: "xml: encoding \"XXX\" declared but Decoder.CharsetReader is nil".
func parseXML(xmlDoc []byte, target interface{}) error {
	reader := bytes.NewReader(xmlDoc)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(target)
	if err != nil {
		return fmt.Errorf("failed to decode target: %v", err)
	}

	return nil
}

// parseTime is a handy function to parse most of known date-time formats.
func parseTime(src string) (t time.Time, err error) {
	knownTimeLayouts := []string{time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}
	for _, layout := range knownTimeLayouts {
		t, err = time.Parse(layout, src)
		if err == nil {
			return
		}
	}

	return
}
