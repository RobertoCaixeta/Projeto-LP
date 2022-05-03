package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
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
/*func formacao(numZagueiro, numLateral, numMeio, numAtacante int, atletas []Atleta) map[string][]string{
	// tecnico
	i := 0
	for _, atleta := range atletas{
		for i < numZagueiro, i++{
		// enquanto tiver zagueiro 
			atleta[0]
			fmt.Println("Tecnico:", zagueiro)
		}
	}	
	resultado := make(map[string][]string)
	resultado["zagueiro"] = ["zagueiro1"]
	return resultado
	zagueiros := 
	numGoleiro, numTecnico := 1, 1
	for _, atleta := range atletas{
		// se ele ja pegou todas as posicoes
		if zagueiro == 0 && lateral == 0 && meio == 0 && ataque == 0{
			return resultado
		}

		if atleta.PosicaoID == 6 && tecnico != 0{
			tecnico -= 1
		}
		if atleta.PosicaoID == 6{
			return atleta.Nome
		}
	

}*/

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

func getAtletaId(data []Atleta, id int) []Atleta{
	var s []Atleta

	for _, atleta := range data{
		var player Atleta
		player.Nome = atleta.Nome
		player.RodadaID = atleta.RodadaID
		player.JogosNum = atleta.JogosNum
		player.MediaNum = atleta.MediaNum
		player.PrecoNum = atleta.PrecoNum
		player.ClubeID = atleta.ClubeID
		player.PosicaoID  = atleta.PosicaoID
		player.PontosNum = atleta.PontosNum
		player.VariacaoNum = atleta.VariacaoNum

		if player.PosicaoID == id {
			s = append(s, player)
		}
		
	}
	return s
}



func main() {
	// baseUrl := "https://api.cartola.globo.com/"

	data := getCartolaData()

	var atletas []Atleta

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

	sort.SliceStable(atletas, 
		func (i, j int) bool{
			if atletas[i].PosicaoID != atletas[j].PosicaoID{
				return atletas[i].PosicaoID > atletas[j].PosicaoID
			}
			return atletas[i].PontosNum > atletas[j].PontosNum
	})

	fmt.Println(len(atletas))

	goleiros := getAtletaId(atletas, 1)
	laterais := getAtletaId(atletas, 2)
	zagueiros := getAtletaId(atletas, 3)
	meias := getAtletaId(atletas, 4)
	atacantes := getAtletaId(atletas, 5)
	tecnicos := getAtletaId(atletas, 6)
	
	fmt.Println("Numero de Goleiros:", len(goleiros))
	fmt.Println("Numero de Laterais:", len(laterais))
	fmt.Println("Numero de Zagueiros:", len(zagueiros))
	fmt.Println("Numero de Meias:", len(meias))
	fmt.Println("Numero de Atacantes:", len(atacantes))
	fmt.Println("Numero de Tecnicos:", len(tecnicos))
	fmt.Println(len(goleiros) + len(laterais) + len(zagueiros) + len(meias) + len(atacantes) + len(tecnicos))
	
	// file, _ := json.MarshalIndent(atletas, "", " ")
	// _ = ioutil.WriteFile("cartola.json", file, 0644)
}
