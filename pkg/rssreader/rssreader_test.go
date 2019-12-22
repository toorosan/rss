package rssreader

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	urls := []string{"http://example.com/"}
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

	// ToDo: add test for supported version, when will be available.
}
