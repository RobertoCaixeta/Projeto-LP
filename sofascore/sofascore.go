package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"fmt"
	"github.com/gocolly/colly"
)

// Course stores information about a coursera course
type Player struct {
	Name       string
	Rating      string
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: coursera.org, www.coursera.org
		colly.AllowedDomains("sofascore.com", "www.sofascore.com"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./sofascore_cache"),
	)

	// Create another collector to scrape course details
	detailCollector := c.Clone()

	players := make([]Player, 0, 200)

	// On every a element which has href attribute call callback
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	// If attribute class is this long string return from callback
	// 	// As this a is irrelevant
	// 	if e.Attr("class") == "Button_1qxkboh-o_O-primary_cv02ee-o_O-md_28awn8-o_O-primaryLink_109aggg" {
	// 		return
	// 	}
	// 	link := e.Attr("href")
	// 	// If link start with browse or includes either signup or login return from callback
	// 	if !strings.HasPrefix(link, "/browse") || strings.Index(link, "=signup") > -1 || strings.Index(link, "=login") > -1 {
	// 		return
	// 	}
	// 	// start scaping the page under the link found
	// 	e.Request.Visit(link)
	// })

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	// On every a HTML element which has name attribute call callback
	c.OnHTML(`a[href]`, func(e *colly.HTMLElement) {
		// Activate detailCollector if the link contains "team/football/"
		playerURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Index(playerURL, "team/football/") != -1 {
			detailCollector.Visit(playerURL)
		}
	})

	// Extract details of the course
	detailCollector.OnHTML(`a[href]`, func(f *colly.HTMLElement) {
		log.Println("Estou no site", f.Request.URL)
		linkPlayer := f.Attr("href")
		if strings.HasPrefix(linkPlayer, "/player/"){
			name := f.ChildText(".sc-18688171-0")
			rating := f.ChildText(".sc-b0a04d91-0")
			// if name == "" {
			// 	log.Println("No name found", f.Request.URL)
			// }
			player := Player{
				Name:       name,
				Rating:     rating,
			}
			
			players = append(players, player)
		}else{
			fmt.Println("Não encontrou um link dentro desse site de jogador")
		}
	})

	// Start scraping on http://coursera.com/browse
	c.Visit("https://www.sofascore.com/tournament/football/brazil/brasileiro-serie-a/325")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(players)
}

	
	
	