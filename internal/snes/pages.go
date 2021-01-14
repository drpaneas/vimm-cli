package snes

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

const (
	// PageSNES is the main page for SNES roms and should always end with a trailing slash
	snesRootURL string = "https://vimm.net/vault/SNES/"
	abc         string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var snesPage []Page

// Page defines a page
// TODO move it centrally
type Page struct {
	URL           string
	NumberOfGames int
	HTML          *goquery.Document
}

// Step 1: Must be called first, because it sets the limit of the slice length
func constructSnesPagesLinks() {
	for _, rune := range abc {
		link := fmt.Sprintf("%s%c", snesRootURL, rune)
		snesPage = append(snesPage, Page{URL: link})
	}
}

// Just for debugging for Step 1
func printSnesPagesLinks() {
	for i := 0; i < len(snesPage); i++ {
		fmt.Println(snesPage[i].URL)
	}
}

// Step 2: Contact the server and fetch the HTML code for all pages
func fetchSnesPagesHTML() {
	for i := 0; i < len(snesPage); i++ {
		fetchSingleSnesPageHTML(i)
	}
}

// Step 2 helper: Fetch HTML per page
func fetchSingleSnesPageHTML(i int) {
	// getHTML retries 10 times. Terminate if all retries fail.
	doc, err := getHTML(snesPage[i].URL)
	if err != nil {
		log.Fatalf("Couldn't get HTML for %s\n", snesPage[i].URL)
	}
	// A valid HTML page is one that includes games, not just HTML code
	for try := 1; try <= 10; try++ {
		if fetchGamesFromSnesPage(doc) == 0 {
			fmt.Printf("\n\n\nPROBLEM WITH GAMES PER PAGE %d/10\n\n", try)
			doc, _ = getHTML(snesPage[i].URL)
		} else {
			break
		}
	}
}
