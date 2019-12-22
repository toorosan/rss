package rssreader

import (
	"fmt"
	"testing"
)

func makeTestRSSBlob(version string) []byte {
	return []byte(fmt.Sprintf(`<rss version=%q>
    <channel>
        <title>Example News</title>
        <link>http://example.com/</link>
        <description>Example channel description.</description>
        <language>en-us</language>
        <pubDate>Tue, 10 Jun 2003 04:00:00 GMT</pubDate>
        <lastBuildDate>Tue, 10 Jun 2003 09:41:01 GMT</lastBuildDate>
        <docs>http://example.com/rss</docs>
        <generator>Weblog Editor 2.0</generator>
        <managingEditor>editor@example.com</managingEditor>
        <webMaster>webmaster@example.com</webMaster>
        <item>
            <title>Example item title</title>
            <link>http://example.com/</link>
            <description>Example description with <a href="http://example.com/">example link</a> inside.</description>
            <pubDate>Tue, 03 Jun 2003 09:39:21 GMT</pubDate>
            <guid>http://example.com/#421</guid>
        </item>
    </channel>
</rss>`, version))
}

func TestGetRSSVersion(t *testing.T) {
	expectedVersion := rssVerion("2.0")
	src := makeTestRSSBlob(string(expectedVersion))
	version, err := getRSSVersion([]byte(src))
	if err != nil {
		t.Fatalf("failed to parse source rss: %v", err)
	}

	if expectedVersion != version {
		t.Fatalf("failed to ensure rss version: %q != %q", expectedVersion, version)
	}
}

func TestGetRSSVersionFail(t *testing.T) {
	src := "some incorrect xml :)"
	_, err := getRSSVersion([]byte(src))
	if err == nil {
		t.Fatal("failed to ensure xml parse error")
	}
	expectedError := "failed to unmarshal xml: EOF"
	if expectedError != err.Error() {
		t.Fatalf("failed to ensure xml parse error is same as expected: %q != %q", expectedError, err.Error())
	}
}
