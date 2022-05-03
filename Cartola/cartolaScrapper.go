package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Posicao struct {
	id         int
	nome       string
	abreviacao string
}

type Clube struct {
	id            int
	Nome          string
	abreviacao    string
	nome_fantasia string
}

type Atleta struct {
	Nome             string
	RodadaID         int
	ClubeID          int
	PosicaoID        int
	PontosNum        float64
	PrecoNum         float64
	VariacaoNum      float64
	MediaNum         float64
	JogosNum         int
}

type Response struct {
	Atletas []struct {
		Scout               struct{}    `json:"scout"`
		AtletaID            int         `json:"atleta_id"`
		RodadaID            int         `json:"rodada_id"`
		ClubeID             int         `json:"clube_id"`
		PosicaoID           int         `json:"posicao_id"`
		StatusID            int         `json:"status_id"`
		PontosNum           float64     `json:"pontos_num"`
		PrecoNum            float64     `json:"preco_num"`
		VariacaoNum         float64     `json:"variacao_num"`
		MediaNum            float64     `json:"media_num"`
		JogosNum            int         `json:"jogos_num"`
		MinimoParaValorizar interface{} `json:"minimo_para_valorizar"`
		Slug                string      `json:"slug"`
		Apelido             string      `json:"apelido"`
		ApelidoAbreviado    string      `json:"apelido_abreviado"`
		Nome                string      `json:"nome"`
		Foto                string      `json:"foto"`
	} `json:"atletas"`
}

func getCartolaData() Response {
	res, err := http.Get("https://api.cartola.globo.com/atletas/mercado")

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	// read data
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Coun't unmarshal json")
	}

	return data
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func main() {
	// baseUrl := "https://api.cartola.globo.com/"

	data := getCartolaData()

	var atletas []Atleta
	// var clubes []Clube

	for _, atleta := range data.Atletas {
		var player Atleta
		player.Nome = atleta.Apelido
		player.RodadaID = atleta.RodadaID
		player.JogosNum = atleta.JogosNum
		player.MediaNum = atleta.MediaNum
		player.PrecoNum = atleta.PrecoNum
		player.ClubeID = atleta.ClubeID
		player.PosicaoID  = atleta.PosicaoID
		player.PontosNum = atleta.PontosNum
		player.VariacaoNum = atleta.VariacaoNum

		atletas = append(atletas, player)
	}

	file, _ := json.MarshalIndent(atletas, "", " ")
	_ = ioutil.WriteFile("cartola.json", file, 0644)
}
