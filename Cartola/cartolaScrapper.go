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

func formacao(data []Atleta, numLateral, numZagueiro, numMeia, numAtacante int){
	// retorna os melhores dada a formaçao
	getAtletaId(data, 1, "Goleiro", 1)
	getAtletaId(data, 2, "Lateral", numLateral)
	getAtletaId(data, 3, "Zagueiro", numZagueiro)
	getAtletaId(data, 4, "Meia", numMeia)
	getAtletaId(data, 5, "Atacante", numAtacante)
	getAtletaId(data, 6, "Técnico", 1)
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
		if atleta.PosicaoID == id && cont != 0{
			fmt.Println(posicao + ":", atleta.Nome + "(" + times[atleta.ClubeID] + ")",  "Pontuação:", atleta.PontosNum, "Media:", atleta.MediaNum)
			cont -= 1
		}
	}
}

var wg sync.WaitGroup
func main() {
	// define os times e os ids dele
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

	//consome os dados do cartola
	data := getCartolaData()

	//transforma os dados mais importantes em uma slice que contem structs do atleta
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


	// Estatisticas Do Cartola

	// ordena por posicao e depois por pontuaçao
	sort.SliceStable(atletas, 
		func (i, j int) bool{
			if atletas[i].PosicaoID != atletas[j].PosicaoID{
				return atletas[i].PosicaoID > atletas[j].PosicaoID
			}
			return atletas[i].PontosNum > atletas[j].PontosNum
	})
	file, _ := json.MarshalIndent(atletas, "", " ")
	_ = ioutil.WriteFile("cartola.json", file, 0644)
	
	melhor_jogador_rodada := atletas[0]
	start := time.Now()
	// paralelismo
	wg.Add(1)
	fmt.Println("Time da Rodada")
	formacao(atletas, 0, 3, 4, 3)

	sort.SliceStable(atletas, 
		func (i, j int) bool{
			return atletas[i].MediaNum > atletas[j].MediaNum
	})

	melhor_jogador_campeonato := atletas[0]
	fmt.Println("----------------------------------------------")
	fmt.Println("Time do Campeonato")
	formacao(atletas, 0, 3, 4, 3)

	fmt.Println("----------------------------------------------")
	fmt.Println("Pior time da Rodada")

	sort.SliceStable(atletas, 
		func (i, j int) bool{
			return atletas[i].PontosNum < atletas[j].PontosNum
	})

	pior_jogador_rodada := atletas[0]
	formacao(atletas, 0, 3, 4, 3)
	
	fmt.Println("----------------------------------------------")
	fmt.Println("Pior time do Campeonato")

	sort.SliceStable(atletas, 
		func (i, j int) bool{
			return atletas[i].MediaNum < atletas[j].MediaNum
	})
	pior_jogador_campeonato := atletas[0]
	formacao(atletas, 0, 3, 4, 3)

	// sort.SliceStable(atletas, 
	// 	func (i, j int) bool{
	// 		return atletas[i].PontosNum > atletas[j].PontosNum
	// })

	fmt.Println("----------------------------------------------")
	fmt.Println("O melhor jodador da rodada:", melhor_jogador_rodada.Nome+ "(" + times[melhor_jogador_rodada.ClubeID] + ")", melhor_jogador_rodada.PontosNum)

	sort.SliceStable(atletas, 
		func (i, j int) bool{
			return atletas[i].MediaNum > atletas[j].MediaNum
	})

	fmt.Println("----------------------------------------------")
	fmt.Println("O melhor jodador do campeonato", melhor_jogador_campeonato.Nome + "(" + times[melhor_jogador_campeonato.ClubeID] + ")", melhor_jogador_campeonato.MediaNum, melhor_jogador_campeonato.PosicaoID)

	// sort.SliceStable(atletas, 
	// 	func (i, j int) bool{
	// 		return atletas[i].PontosNum < atletas[j].PontosNum
	// })

	fmt.Println("----------------------------------------------")
	fmt.Println("O pior jogador da rodada", pior_jogador_rodada.Nome + "(" + times[pior_jogador_rodada.ClubeID] + ")", pior_jogador_rodada.PontosNum)

	// sort.SliceStable(atletas, 
	// 	func (i, j int) bool{
	// 		return atletas[i].MediaNum < atletas[j].MediaNum
	// })

	fmt.Println("----------------------------------------------")
	fmt.Println("O pior jogador do campeonato", pior_jogador_campeonato.Nome + "(" + times[pior_jogador_campeonato.ClubeID] + ")", pior_jogador_campeonato.MediaNum)
	elapsed := time.Since(start)
	fmt.Println("----------------------------------------------")
	fmt.Println("tempo com paralelismo:", elapsed)
	
	wg.Done()
	wg.Wait()
}
