package controller

import (
	"PROYECTO/dto"
	"PROYECTO/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type SearchController interface {
	FindSongs(ctx *gin.Context)
}

type searchController struct {
	itunesService      service.GetItunesSongsService
	musicLyricsService service.GetMLirycsSongsService
	saveSong           service.SaveSongsService
	jwtController      JwtController
}

func NewSearchController() SearchController {
	return &searchController{
		itunesService:      service.NewItunesService(),
		musicLyricsService: service.NewMusicLyricsService(),
		saveSong:           service.NewSaveSongsService(),
		jwtController:      NewJwtController(),
	}
}

func handleErroResponse(ctx *gin.Context, err string, statusCode int) {
	response := dto.SearchResponseDTO{
		Status: false,
		Error:  err,
		Data:   nil,
	}

	ctx.JSON(statusCode, response)
	ctx.Abort()
}

func (services *searchController) FindSongs(ctx *gin.Context) {

	nameGet := ctx.Query("name")
	artistGet := ctx.Query("artist")
	albumGet := ctx.Query("album")
	tokenGet := ctx.Query("token")

	//Verificar que los parametros no esten vacios
	if nameGet == "" && artistGet == "" && albumGet == "" {
		handleErroResponse(ctx, "No se ha ingresado ningun parametro de busqueda", http.StatusBadRequest)
		return
	}

	//Verificar que el token no este vacio
	if tokenGet == "" {
		handleErroResponse(ctx, "No se ha ingresado el token", http.StatusBadRequest)
		return
	}

	//Verificar que el token sea valido
	validate := services.jwtController.ValidateToken(tokenGet)
	if !validate {
		handleErroResponse(ctx, "El token no es valido", http.StatusBadRequest)
		return
	}

	//Guardar la busqueda en el dto
	searchRequest := dto.SearchRequestDTO{
		Name:   nameGet,
		Artist: artistGet,
		Album:  albumGet,
	}

	//Buscar en la base de datos primero
	songsFindedFirst, err := services.saveSong.SearchSongs(searchRequest)
	if err != nil {
		handleErroResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(songsFindedFirst) > 0 {
		response := dto.SearchResponseDTO{
			Status: true,
			Error:  "bd search",
			Data:   songsFindedFirst,
		}
		ctx.JSON(http.StatusOK, response)
		return
	}
	//Si no se encuentra en la base de datos, buscar en las apis

	//buscar en ambas al mismo tiempo
	chanItunes := make(chan dto.ItunesResultDTO)
	chanMusicLyrics := make(chan dto.ChartLirycsResultDTO)

	//Buscar en itunes
	go services.itunesService.GetItunesSongs(searchRequest, chanItunes)

	//Buscar en musiclyrics
	go services.musicLyricsService.GetMLirycsSongs(searchRequest, chanMusicLyrics)

	//Extraer los datos de las respuestas
	songs := extractSongs(<-chanItunes, <-chanMusicLyrics)

	//Guardar la busqueda en la base de datos
	err = services.saveSong.SaveSongs(songs)
	if err != nil {
		handleErroResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	//Buscar en la base de datos
	songsFinded, err := services.saveSong.SearchSongs(searchRequest)
	if err != nil {
		handleErroResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	//Retornar la respuesta
	response := dto.SearchResponseDTO{
		Status: true,
		Error:  "",
		Data:   songsFinded,
	}
	ctx.JSON(http.StatusOK, response)
}

func extractSongs(itunesResponse dto.ItunesResultDTO, musicLyricsResponse dto.ChartLirycsResultDTO) []dto.SongSaveInfo {
	var songs []dto.SongSaveInfo
	for _, itunesSong := range itunesResponse.Results {

		//crear id unico
		id := fmt.Sprintf("I%s%s", strings.ReplaceAll(itunesSong.TrackName, " ", ""), strconv.Itoa(itunesSong.TrackId))

		if id == "I" {
			continue
		}

		song := dto.SongSaveInfo{
			ID:       id,
			Name:     itunesSong.TrackName,
			Artist:   itunesSong.ArtistName,
			Duration: itunesSong.TrackTimeMillis,
			Album:    itunesSong.CollectionName,
			Artwork:  itunesSong.ArtworkURL100,
			Price:    itunesSong.TrackPrice,
			Currency: itunesSong.Currency,
			Origin:   "itunes",
			IDOrigin: itunesSong.TrackId,
		}
		songs = append(songs, song)
	}

	for _, musicLyricsSong := range musicLyricsResponse.SearchLyricResult {

		//crear id unico
		id := fmt.Sprintf("M%s%s", strings.ReplaceAll(musicLyricsSong.Song, " ", ""), strconv.Itoa(musicLyricsSong.LyricId))

		if id == "M" {
			continue
		}

		song := dto.SongSaveInfo{
			ID:       id,
			Name:     musicLyricsSong.Song,
			Artist:   musicLyricsSong.Artist,
			Origin:   "chartlyrics",
			IDOrigin: musicLyricsSong.LyricId,
		}
		songs = append(songs, song)
	}
	return songs
}
