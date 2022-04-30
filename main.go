package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	crawler()
}

func crawler() {
	c := colly.NewCollector(
		colly.MaxDepth(10),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL.String())
	})

	c.OnHTML("div > a", func(e *colly.HTMLElement) {
		// listPlayers := e.ChildAttrs("a", "href")
		fmt.Println("Element: ", e)

	})

	c.Visit("https://www.sofascore.com/tournament/football/brazil/brasileiro-serie-a/325")
}
