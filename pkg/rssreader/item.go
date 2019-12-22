package rssreader

import (
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
