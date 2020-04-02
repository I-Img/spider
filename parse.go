package main

import (
	"encoding/xml"
	"fmt"
)

type rsp struct {
	XMLName xml.Name `xml:"rsp"`
	Text    string   `xml:",chardata"`
	Stat    string   `xml:"stat,attr"`
	Photos  struct {
		Text    string `xml:",chardata"`
		Page    string `xml:"page,attr"`
		Pages   string `xml:"pages,attr"`
		Perpage string `xml:"perpage,attr"`
		Total   string `xml:"total,attr"`
		Photo   []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Owner    string `xml:"owner,attr"`
			Secret   string `xml:"secret,attr"`
			Server   string `xml:"server,attr"`
			Farm     string `xml:"farm,attr"`
			Title    string `xml:"title,attr"`
			Ispublic string `xml:"ispublic,attr"`
			Isfriend string `xml:"isfriend,attr"`
			Isfamily string `xml:"isfamily,attr"`
		} `xml:"photo"`
	} `xml:"photos"`
}

func parse(data []byte) (ids []string, err error) {

	var r rsp

	err = xml.Unmarshal(data, &r)
	if err != nil {
		return
	}

	for _, i := range r.Photos.Photo {
		ids = append(ids, i.ID)
	}

	return
}

type size struct {
	XMLName xml.Name `xml:"rsp"`
	Text    string   `xml:",chardata"`
	Stat    string   `xml:"stat,attr"`
	Sizes   struct {
		Text        string `xml:",chardata"`
		Canblog     string `xml:"canblog,attr"`
		Canprint    string `xml:"canprint,attr"`
		Candownload string `xml:"candownload,attr"`
		Size        []struct {
			Text   string `xml:",chardata"`
			Label  string `xml:"label,attr"`
			Width  string `xml:"width,attr"`
			Height string `xml:"height,attr"`
			Source string `xml:"source,attr"`
			URL    string `xml:"url,attr"`
			Media  string `xml:"media,attr"`
		} `xml:"size"`
	} `xml:"sizes"`
}

func parseSize(data []byte) (url []string, err error) {
	var s size

	err = xml.Unmarshal(data, &s)
	if err != nil {
		return
	}

	for _, i := range s.Sizes.Size {
		if i.Label == "Thumbnail" {
			url = append(url, fmt.Sprintf("Thumbnail|%s", i.Source))
		}

		if i.Label == "Small 320" {
			url = append(url, fmt.Sprintf("Small|%s", i.Source))
		}

		if len(url) == 2 {
			break
		}
	}
	return
}
