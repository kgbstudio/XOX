```go
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Items       []Item    `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func getRss(url string) (*Rss, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rss Rss
	err = xml.Unmarshal(bytes, &rss)
	if err != nil {
		return nil, err
	}

	return &rss, nil
}

func main() {
	url := "https://example.com/rss"
	rss, err := getRss(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range rss.Channel.Items {
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Link: %s\n", item.Link)
		fmt.Printf("Description: %s\n", item.Description)
		fmt.Printf("PubDate: %s\n", item.PubDate)
		fmt.Println()
	}
}