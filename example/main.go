package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	tinyspider "github.com/xuhe2/TinySpider"
)

func main() {
	spider := tinyspider.NewSpider()

	spider.AddTask(func(doc *goquery.Document) {
		fmt.Println(doc.Find("title").Text())
	})

	spider.Get("https://github.com/")
}
