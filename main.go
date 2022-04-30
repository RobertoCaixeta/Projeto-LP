package main

import (
	"fmt"
	"strings"

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

	linkOfPlayers := []string{}

	c.OnHTML("div > a", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, h *colly.HTMLElement) {
			if strings.Contains(h.Attr("href"), "/player") {
				linkOfPlayers = append(linkOfPlayers, h.Attr("href"))
			}
		})
		// listPlayers := e.ChildAttrs("a", "href")
		fmt.Println("Element: ", e)

	})

	c.Visit("https://www.sofascore.com/tournament/football/brazil/brasileiro-serie-a/325")
}
