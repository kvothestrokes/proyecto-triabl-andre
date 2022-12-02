package service

import (
	"PROYECTO/dto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type GetItunesSongsService interface {
	GetItunesSongs(song dto.SearchRequestDTO, ch chan<- dto.ItunesResultDTO)
}

type getItunesSongsService struct{}

func NewItunesService() GetItunesSongsService {
	return &getItunesSongsService{}
}

const ITUNES_ENDPOINT = "https://itunes.apple.com/search?term="

func (i *getItunesSongsService) GetItunesSongs(song dto.SearchRequestDTO, ch chan<- dto.ItunesResultDTO) {

	var itunesResult dto.ItunesResultDTO

	fmt.Println("\nBuscando en Itunes...")

	//Verificar que los parametros no esten vacios
	params := ""
	if song.Name != "" {
		params += strings.ReplaceAll(song.Name, " ", "+")
	}
	if song.Artist != "" {
		params += "+" + strings.ReplaceAll(song.Artist, " ", "+")
	}
	if song.Album != "" {
		params += "+" + strings.ReplaceAll(song.Album, " ", "+")
	}

	//Realizar la peticion a Itunes
	responseRaw, err := http.Get(ITUNES_ENDPOINT + params)

	if err != nil {
		return
	}
	defer responseRaw.Body.Close()

	response, err := ioutil.ReadAll(responseRaw.Body)
	if err != nil {
		return
	}

	json.Unmarshal(response, &itunesResult)

	ch <- itunesResult

}
