package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main (){
	c := colly.NewCollector(colly.AllowedDomains("www.sofascore.com", "sofascore.com") )

	c.OnHTML("span.styles__PlayerName-sc-1loq6tv-14", func(h *colly.HTMLElement) {
		fmt.Println("Nome do Jogador:",h.Text)
	})


	c.OnHTML("div.styles__Container-sc-3ao04p-4 table.styles__StatisticsGroupTable-sc-3ao04p-6 tr td", func(h *colly.HTMLElement) {
		selection := h.DOM
		
		fmt.Println(selection.Find("td").Text())
		fmt.Println("atributo:",h.Text)
	})

	c.Visit("https://www.sofascore.com/player/hulk/34705")
}