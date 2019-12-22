package rssreader

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

const rss20 rssVerion = "2.0"

// rss20XML describes base RSS2.0 feed structure.
type rss20XML struct {
	Description string      `xml:"channel>description"`
	Items       []rss20Item `xml:"channel>item"`
	Link        string      `xml:"channel>link"`
	Title       string      `xml:"channel>title"`
}

// rss20Item describes single raw item of RSS2.0 feed.
type rss20Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
}

// rss20Parser implements RSS parser for RSS2.0 protocol
type rss20Parser struct {
	blob []byte
}

func (r *rss20Parser) Version() rssVerion {
	return rss20
}

func (r *rss20Parser) Parse() ([]RssItem, error) {
	fail := func(err error) ([]RssItem, error) {
		return nil, err
	}
	raw := rss20XML{}
	err := xml.Unmarshal(r.blob, &raw)
	if err != nil {
		return fail(err)
	}
	result := make([]RssItem, len(raw.Items))
	for i, item := range raw.Items {
		pDate, err := http.ParseTime(item.PubDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date %q", item.PubDate)
		}
		result[i] = RssItem{
			Title:       item.Title,
			Source:      raw.Title,
			SourceURL:   raw.Link,
			Link:        item.Link,
			PubishDate:  pDate,
			Description: item.Description,
		}
	}

	return result, nil
}
