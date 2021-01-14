package snes

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

var SnesRom []Rom // Build an array that will have numberOfGames elements from type Rom

// TODO move it centrally
// Host without trailing slash at the end
const Host string = "https://vimm.net"

// Rom defines a typical game card
// has to start with capital letter or the Marshal won't work in JSON
type Rom struct {
	Title        string `json:"title"`
	Link         string `json:"link"`
	DownloadLink string `json:"download_link"`
	Filename     string `json:"filename"`
	Image        string `json:"image"`
	Quality      string `json:"quality"`
	Hack         string `json:"hack"`
	Console      string `json:"console"`
	Filesize     string `json:"filesize"`
	Genre        string `json:"genre"`
	Downloads    string `json:"downloads"`
	Rating       string `json:"rating"`
	Players      int    `json:"players"`
	Year         int    `json:"year"`
	Publisher    string `json:"publisher"`
	Manual       string `json:"manual"`
}

// https://download.vimm.net/download/?mediaId=983

// Step 3 (used in 2), fetch all game links per page
// also returns how many they are
func fetchGamesFromSnesPage(doc *goquery.Document) int {
	counter := 0
	doc.Find("#innerMain > table.rounded.centered.cellpadding1.hovertable > tbody > tr > td").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			link, ok := s.Attr("href")
			if ok {
				if !(strings.Contains(link, "rating") || strings.Contains(link, "manual")) {
					SnesRom = append(SnesRom, Rom{Link: Host + link})
					fmt.Println(s.Text())
					counter++
				}
			}
		})
	})
	return counter
}

// Just for debugging for Step 3
func printGamesLinks() {
	counter := 0
	for i := 0; i < len(SnesRom); i++ {
		fmt.Println(SnesRom[i].Link)
		counter++
	}
	fmt.Printf("There are %d games for snes", counter)
}

func fetchDownloadLink(romLink string) {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(romLink),
		page.SetDownloadBehavior(page.SetDownloadBehaviorBehaviorAllow).WithDownloadPath("."),
		chromedp.Click(`#download_form > button`),
		chromedp.Sleep(20*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("success")

}

func downloadAllSnesRoms() {
	for i := 0; i < len(SnesRom); i++ {
		fetchDownloadLink(SnesRom[i].Link)
	}
}
