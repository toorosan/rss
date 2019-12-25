package rssreader

import (
	"fmt"
	"time"
)

// RssItem describes single RSS item.
type RssItem struct {
	Title       string
	Source      string
	SourceURL   string
	Link        string
	PubishDate  time.Time
	Description string
}

func (ri RssItem) String() string {
	return fmt.Sprintf("%s: %s - %q - %s", ri.Source, ri.PubishDate, ri.Title, ri.Link)
}
