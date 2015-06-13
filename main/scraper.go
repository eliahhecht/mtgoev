package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/wayn3h0/go-decimal"
	"io"
	"strings"
)

type Downloader interface {
	DownloadUrl(url string) (io.Reader, error)
}

type Scraper struct {
	Downloader Downloader
}

func parseCards(tBody *goquery.Selection) []Card {
	if tBody == nil {
		panic("Can't parse a non-existent tbody")
	}
	allCards := make([]Card, 0, 0)
	tBody.Find("tr").Each(func(i int, s *goquery.Selection) {
		card := parseCard(s)
		if !(card.Rarity == RARITY_BASIC) {
			allCards = append(allCards, card)
		}
	})
	return allCards
}

func parseCard(s *goquery.Selection) Card {
	tds := s.Find("td")

	rarityText := tds.Eq(2).Text()
	rarity := parseRarityText(rarityText)

	name := tds.Eq(0).Text()

	priceText := tds.Eq(3).Text()
	price, _ := decimal.Parse(strings.TrimSpace(priceText))

	card := Card{Name: name, Rarity: rarity, Price: price}
	return card
}

func parseRarityText(s string) int {
	switch s {
	case "Common":
		return RARITY_C
	case "Uncommon":
		return RARITY_U
	case "Rare":
		return RARITY_R
	case "Mythic":
		return RARITY_M
	default:
		return RARITY_BASIC
	}
}

func (scraper *Scraper) ScrapeSet(setName string) *Set {
	r, _ := scraper.Downloader.DownloadUrl("")
	doc, _ := goquery.NewDocumentFromReader(r) // todo: error handling
	lastTBody := doc.Find("tbody").Last()
	cards := parseCards(lastTBody)
	return &Set{Cards: cards}
}
