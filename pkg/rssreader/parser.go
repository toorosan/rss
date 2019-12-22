package rssreader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
