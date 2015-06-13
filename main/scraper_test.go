package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/wayn3h0/go-decimal"
	"io"
	"os"
	"testing"
)

type CannedDtkDownloader struct{}

func (d *CannedDtkDownloader) DownloadUrl(url string) (r io.Reader, err error) {
	return os.Open("dtk.html")
}

func TestScraperGetsAllCards(t *testing.T) {
	downloader := &CannedDtkDownloader{}
	scraper := &Scraper{Downloader: downloader}

	set := scraper.ScrapeSet("dtk")

	cardCount := len(set.Cards)
	assert.Equal(t, 248, cardCount, "There are 248 cards in DTK (excluding basic lands)")
}

func TestPriceIsParsedCorrectlyForDeathmistRaptor(t *testing.T) {
	downloader := &CannedDtkDownloader{}
	scraper := &Scraper{Downloader: downloader}

	set := scraper.ScrapeSet("dtk")

	for _, card := range set.Cards {
		if card.Name == "Deathmist Raptor" {
			assert.Equal(
				t,
				decimal.New(19.89),
				card.Price,
				"The card's price should be parsed correctly")
			return
		}
	}
	assert.Fail(t, "Deathmist Raptor not found")
}
