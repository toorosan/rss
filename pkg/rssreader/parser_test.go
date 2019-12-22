package rssreader

import (
	"fmt"
	"net/url"
	"testing"
)

func TestNewParserForUnsupportedVersion(t *testing.T) {
	u, err := url.Parse("http://example.com/")
	if err != nil {
		t.Fatalf("failed to prepare test: %v", err)
	}
	expectedVersion := "unsupported-version"
	p, err := newParser(*u, &mockedGetter{
		result:        makeTestRSSBlob(expectedVersion),
		expectedError: nil,
	})
	if err == nil {
		t.Fatalf("failed to ensure unsupported rss version is handled properly")
	}
	if p != nil {
		t.Fatalf("failed to ensure parser is not returned in case of unsupported rss version")
	}
	if err.Error() != fmt.Sprintf("rss version %q is not supported", expectedVersion) {
		t.Fatalf("failed to ensure error text for unsupported rss version")
	}
}

type mockedGetter struct {
	result        []byte
	expectedError error
}

func (mg *mockedGetter) Get(_ url.URL) ([]byte, error) {
	return mg.result, mg.expectedError
}
