package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	getData()
}

func getData() {
	url := "https://www.google.com/search?q=go+tutorials&gl=us&hl=en"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	c := 0
	doc.Find("div.g").Each(func(i int, result *goquery.Selection) {
		title := result.Find("h3").First().Text()
		link, _ := result.Find("a").First().Attr("href")
		snippet := result.Find(".VwiC3b").First().Text()

		fmt.Printf("Title: %s\n", title)
		fmt.Printf("Link: %s\n", link)
		fmt.Printf("Snippet: %s\n", snippet)
		fmt.Printf("Position: %d\n", c+1)
		fmt.Println()

		c++
	})
}
