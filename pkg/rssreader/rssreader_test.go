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
	results, err := Parse(urls)
	if len(results) != 0 {
		t.Fatal("failed to ensure empty results for non-rss url")
	}
	if err == nil {
		t.Fatal("failed to ensure errors for non-rss url")
	}
	if _, ok := err.(*errorList); !ok {
		t.Fatal("failed to ensure error type is error list")
	}
	if err.Error() != "failed to get specific rss parser for url \"http://example.com\": failed to get blob data from url \"http://example.com\"" {
		t.Fatal("failed to ensure error text")
	}

	// test for unsupported version
	feedGetter = &mockedGetter{
		result: makeTestRSSBlob("unsupported-version"),
	}
	results, err = Parse(urls)
	if len(results) != 0 {
		t.Fatal("failed to ensure empty results for unsupported version of rss")
	}
	if err == nil {
		t.Fatal("failed to ensure errors for unsupported version of rss")
	}
	if _, ok := err.(*errorList); !ok {
		t.Fatal("failed to ensure error type is error list")
	}
	if err.Error() != "failed to get specific rss parser for url \"http://example.com\": rss version \"unsupported-version\" is not supported" {
		t.Fatal("failed to ensure error text")
	}

	// test for supported version
	feedGetter = &mockedGetter{
		result: makeTestRSSBlob(string(rss20)),
	}
	results, err = Parse(urls)
	if err != nil {
		t.Fatal("failed to ensure no errors if 3 same URLs were passed")
	}
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
	if results[0].PubishDate.String() != "2003-06-03 09:39:21 +0000 GMT" {
		t.Fatal("failed to ensure PubishDate field is same as expected")
	}
	if results[0].Description != "Example description with  inside." {
		t.Fatal("failed to ensure Description field is same as expected")
	}

	// test for supported version 2
	feedGetter = &mockedGetter{
		result: makeTestRSSBlob(string(rss20)),
	}
	results, err = Parse([]string{"http://example.com", "https://example.com"})
	if err != nil {
		t.Fatal("failed to ensure no errors are present")
	}

	if len(results) != 2 {
		t.Fatal("failed to ensure 2 results are present")
	}
	if results[0].String() != "Example News: 2003-06-03 09:39:21 +0000 GMT - \"Example item title\" - http://example.com/" {
		t.Fatal("failed to ensure first result is equal to the expectations")
	}
}
