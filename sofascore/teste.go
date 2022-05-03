package main

import(
	"fmt"
	"net/http"
	"time"
	//"github.com/PuerkitoBio/goquery"
)

func getListing(listingURL string) []string {
	var links []string
	//HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	request, err := http.NewRequest("GET", listingURL, nil)
	if err != nil {
		fmt.Println(err)
	}

	//Setting headers
	request.Header.Set("pragma", "no-cache")
	request.Header.Set("cache-control", "no-cache")
	request.Header.Set("dnt", "1")
	request.Header.Set("upgrade-insecure-requests", "1")
	request.Header.Set("referer", "https://www.sofascore.com/")
	resp, err := client.Do(request)

	fmt.Println(resp.StatusCode)

	return links
}

func main() {
	m := getListing("https://www.sofascore.com/tournament/football/brazil/brasileiro-serie-a/325")
	fmt.Println(m)
}