package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

var root = "http://map.cam.ac.uk/colleges"

func checkError(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

func main() {

	resp, err := http.Get(root)
	checkError(err)

	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(io.Reader(resp.Body))
	checkError(err)

	list, err := os.Create("check-list.md")
	defer list.Close()
	checkError(err)

	var line string
	doc.Find("div.campl-map-container h3:first-child").Each(func(i int, col *goquery.Selection) {
		line = fmt.Sprintf("- [ ] %s\n", col.Text())
		_, err := list.WriteString(line)
		checkError(err)

		fmt.Printf("Got %d college(s)...\n", i+1)
	})
}
