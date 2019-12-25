package rssreader

import (
	"encoding/xml"
	"fmt"
)

// rssVerion defines supported RSS protocol versions.
type rssVerion string

func getRSSVersion(blob []byte) (rssVerion, error) {
	var base baseRSSXML
	err := parseXML(blob, &base)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal xml: %v", err)
	}

	return base.Version, nil
}

// baseRSSXML describes base RSS XML structure.
type baseRSSXML struct {
	XMLName xml.Name  `xml:"rss"`
	Version rssVerion `xml:"version,attr"`
}
