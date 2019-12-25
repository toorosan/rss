package main

import "time"

// rssItem describes single RSS item json conversion schema.
type rssItem struct {
	Title       string    `json:"title"`
	Source      string    `json:"source"`
	SourceURL   string    `json:"sourceURL"`
	Link        string    `json:"link"`
	PubishDate  time.Time `json:"pubishDate"`
	Description string    `json:"description"`
}
