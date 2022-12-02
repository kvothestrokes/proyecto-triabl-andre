package service

import (
	"PROYECTO/dto"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type GetMLirycsSongsService interface {
	GetMLirycsSongs(song dto.SearchRequestDTO, ch chan<- dto.ChartLirycsResultDTO)
}

type getMLirycsSongsService struct{}

func NewMusicLyricsService() GetMLirycsSongsService {
	return &getMLirycsSongsService{}
}

const MLIRYCS_ENDPOINT = "http://api.chartlyrics.com/apiv1.asmx/SearchLyric?"

func (i *getMLirycsSongsService) GetMLirycsSongs(song dto.SearchRequestDTO, ch chan<- dto.ChartLirycsResultDTO) {

	var mlirycsResult dto.ChartLirycsResultDTO

	fmt.Println("\nBuscando en MusicLyrics...")

	params := ""
	if song.Name != "" {
		// sustituir espacios por %20 para lectura de url
		params += "song=" + strings.ReplaceAll(song.Name, " ", "%20")
	}
	if song.Artist != "" {
		params += "&artist=" + strings.ReplaceAll(song.Artist, " ", "%20")
	}

	//Realizar la peticion a ChartLyrics
	responseRaw, err := http.Get(MLIRYCS_ENDPOINT + params)

	if err != nil {
		return
	}
	defer responseRaw.Body.Close()

	response, err := ioutil.ReadAll(responseRaw.Body)
	if err != nil {
		return
	}

	// fmt.Printf("Respuesta: %s", string(response))

	xml.Unmarshal(response, &mlirycsResult)

	ch <- mlirycsResult
}
