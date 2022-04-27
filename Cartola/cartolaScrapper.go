package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type cartolaRes struct {
	clubes   *clube
	posicoes *posicao
	status   *status
	altetas  *atleta
}

type clube struct {
	ArrayValue []int
}

type posicao struct {
	ArrayValue []int
}

type status struct {
	ArrayValue []int
}

type atleta struct {
	ArrayValue []int
}

func getCartolaData() []byte {
	res, err := http.Get("https://api.cartola.globo.com/atletas/mercado")

	if err != nil {
		log.Fatalln(err)
	}

	// read data
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Failed to transform to json")
	}

	return jsonData
}

func main() {
	// baseUrl := "https://api.cartola.globo.com/"

	data := getCartolaData()
	fmt.Printf("Json data: %s\n", data)
}

// func main() {
// 	c := colly.NewCollector(
// 		colly.AllowedDomains("https://cartola.globo.com"),
// 		colly.AllowURLRevisit(),
// 	)

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Fazendo requisição para: ", r.URL)
// 	})

// 	c.OnError(func(r *colly.Response, err error) {
// 		fmt.Println("Algo deu errado menorzada: ", err)
// 	})

// 	c.OnResponse(func(r *colly.Response) {
// 		fmt.Println("Visitado", r.Request.URL)
// 	})

// 	c.OnHTML("div[id]", func(d *colly.HTMLElement) {
// 		fmt.Println("Nome do Jogador:", d.Attr("id"))
// 	})

// 	c.Visit("https://cartola.globo.com/#!/time")
// }
