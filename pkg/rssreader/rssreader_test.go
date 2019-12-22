package rssreader

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	urls := []string{"http://example.com", "http://example.com", "http://example.com"}
	// test for unexpected getter error
	feedGetter = &mockedGetter{
		expectedError: fmt.Errorf("failed to get rss feed"),
	}
	results := Parse(urls)
	if len(results) != 0 {
		t.Fatal("failed to ensure empty results for non-rss url")
	}
	// test for unsupported version
	feedGetter = &mockedGetter{
		result: makeTestRSSBlob("unsupported-version"),
	}
	results = Parse(urls)
	if len(results) != 0 {
		t.Fatal("failed to ensure empty results for unsupported version of rss")
	}

	// test for supported version
	feedGetter = &mockedGetter{
		result: makeTestRSSBlob(string(rss20)),
	}
	results = Parse(urls)
	if len(results) != 1 {
		t.Fatal("failed to ensure result is exactly one even if 3 URLs were passed, as they are the same")
	}
	if results[0].Title != "Example item title" {
		t.Fatal("failed to ensure Title field is same as expected")
	}
	if results[0].Source != "Example News" {
		t.Fatal("failed to ensure Source field is same as expected")
	}
	if results[0].SourceURL != "http://example.com/" {
		t.Fatal("failed to ensure SourceURL field is same as expected")
	}
	if results[0].Link != "http://example.com/" {
		t.Fatal("failed to ensure Link field is same as expected")
	}
	if results[0].PubishDate.String() != "2003-06-03 09:39:21 +0000 UTC" {
		t.Fatal("failed to ensure PubishDate field is same as expected")
	}
	if results[0].Description != "Example description with  inside." {
		t.Fatal("failed to ensure Description field is same as expected")
	}

	// test for supported version 2
	feedGetter = &mockedGetter{
		result: makeTestRSSBlob(string(rss20)),
	}
	results = Parse([]string{"http://example.com", "https://example.com"})
	if len(results) != 2 {
		t.Fatal("failed to ensure 2 results are present")
	}
}
