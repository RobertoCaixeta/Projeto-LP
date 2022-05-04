package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"
	//"strconv"
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

func formacao(data []Atleta, numLateral, numZagueiro, numMeia, numAtacante int){
	//c := make(chan string)
	getAtletaId(data, 1, "Goleiro", 1)
	getAtletaId(data, 2, "Lateral", numLateral)
	getAtletaId(data, 3, "Zagueiro", numZagueiro)
	getAtletaId(data, 4, "Meia", numMeia)
	getAtletaId(data, 5, "Atacante", numAtacante)
	getAtletaId(data, 6, "Técnico", 1)
	//msg := <- c
	//fmt.Println(msg)
}
func getAtletaId(data []Atleta, id int, posicao string, cont int){
	times := make(map[int]string)
	times[314] = "Avai"
	times[285] = "Internacional"
	times[262] = "Flamengo"
	times[276] = "Sao Paulo"
	times[263] = "Botafogo"
	times[266] = "Fluminense"
	times[280] = "RB Bragantino"
	times[275] = "Palmeiras"
	times[282] = "Atletico Mineiro"
	times[294] = "Coritiba"
	times[1371] = "Cuiaba"
	times[286] = "Juventude"
	times[290] = "Goias"
	times[327] = "America Mineiro"
	times[373] = "Atletico-Go"
	times[354] = "Ceara"
	times[293] = "Atletico Paranaense"
	times[264] = "Corinthians"
	times[277] = "Santos"
	times[356] = "Fortaleza"

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

		if player.PosicaoID == id && cont != 0{
			fmt.Println(posicao + ":", player.Nome + "(" + times[player.ClubeID] + ")",  "Pontuação:", player.PontosNum, "Media:", player.MediaNum)
			cont -= 1
		}

		// if cont == 0{
		// 	break
		// }
	}
}

var wg sync.WaitGroup
func main() {
	// baseUrl := "https://api.cartola.globo.com/"
	times := make(map[int]string)
	times[314] = "Avai"
	times[285] = "Internacional"
	times[262] = "Flamengo"
	times[276] = "Sao Paulo"
	times[263] = "Botafogo"
	times[266] = "Fluminense"
	times[280] = "RB Bragantino"
	times[275] = "Palmeiras"
	times[282] = "Atletico Mineiro"
	times[294] = "Coritiba"
	times[1371] = "Cuiaba"
	times[286] = "Juventude"
	times[290] = "Goias"
	times[327] = "America Mineiro"
	times[373] = "Atletico-Go"
	times[354] = "Ceara"
	times[293] = "Athletico Paranaense"
	times[264] = "Corinthians"
	times[277] = "Santos"
	times[356] = "Fortaleza"

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
	
	fmt.Println("Time da Rodada")
	start := time.Now()
	wg.Add(1)
	go formacao(atletas, 0, 3, 4, 3)
	wg.Done()

	
	wg.Wait()

	sort.SliceStable(atletas, 
		func (i, j int) bool{
			return atletas[i].MediaNum > atletas[j].MediaNum
	})
	fmt.Println("----------------------------------------------")
	fmt.Println("Time do Campeonato")
	formacao(atletas, 0, 3, 4, 3)

	fmt.Println("----------------------------------------------")
	fmt.Println("Pior time da Rodada")
	sort.SliceStable(atletas, 
		func (i, j int) bool{
			return atletas[i].PontosNum < atletas[j].PontosNum
	})
	formacao(atletas, 0, 3, 4, 3)
	
	fmt.Println("----------------------------------------------")
	fmt.Println("Pior time do Campeonato")

	sort.SliceStable(atletas, 
		func (i, j int) bool{
			return atletas[i].MediaNum < atletas[j].MediaNum
	})

	formacao(atletas, 0, 3, 4, 3)
	sort.SliceStable(atletas, 
		func (i, j int) bool{
			return atletas[i].PontosNum > atletas[j].PontosNum
	})

	fmt.Println("----------------------------------------------")
	fmt.Println("O melhor jodador da rodada:", atletas[0].Nome+ "(" + times[atletas[0].ClubeID] + ")", atletas[0].PontosNum)

	sort.SliceStable(atletas, 
		func (i, j int) bool{
			return atletas[i].MediaNum > atletas[j].MediaNum
	})

	fmt.Println("----------------------------------------------")
	fmt.Println("O melhor jodador do campeonato", atletas[0].Nome + "(" + times[atletas[0].ClubeID] + ")", atletas[0].MediaNum, atletas[0].PosicaoID)

	sort.SliceStable(atletas, 
		func (i, j int) bool{
			return atletas[i].PontosNum < atletas[j].PontosNum
	})

	fmt.Println("----------------------------------------------")
	fmt.Println("O pior jogador da rodada", atletas[0].Nome + "(" + times[atletas[0].ClubeID] + ")", atletas[0].PontosNum)

	sort.SliceStable(atletas, 
		func (i, j int) bool{
			return atletas[i].MediaNum < atletas[j].MediaNum
	})
	fmt.Println("----------------------------------------------")
	fmt.Println("O pior jogador do campeonato"+ "(" + times[atletas[0].ClubeID] + ")",atletas[0].Nome, atletas[0].MediaNum)
	elapsed := time.Since(start)
	fmt.Println("----------------------------------------------")
	fmt.Println("tempo com paralelismo:", elapsed)
	


	

	// file, _ := json.MarshalIndent(atletas, "", " ")
	// _ = ioutil.WriteFile("cartola.json", file, 0644)
}