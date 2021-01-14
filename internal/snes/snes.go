package snes

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
)

// TODO: Move this to a separate package
func getHTML(page string) (doc *goquery.Document, err error) {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	standardClient := retryClient.StandardClient() // *http.Client

	// Request the HTML page.
	res, err := standardClient.Get(page)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc, err
}

// Test in the main
func RunIt() {
	// fetchDownloadLink("https://vimm.net/vault/1640")
	constructSnesPagesLinks()
	printSnesPagesLinks()
	fetchSnesPagesHTML()
	printGamesLinks()
	// downloadAllSnesRoms() -- does not work yet
}
