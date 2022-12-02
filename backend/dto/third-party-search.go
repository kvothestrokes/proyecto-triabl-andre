package dto

import (
	"encoding/xml"
)

// La estructura de la respuesta de la API de ChartLirycs
type ChartLyricSong struct {
	Artist  string `xml:"Artist"`  // Nombre del artista
	Song    string `xml:"Song"`    // Nombre de la cancion
	LyricId int    `xml:"LyricId"` // ID de la cancion
}

type ChartLirycsResultDTO struct {
	XMLName           xml.Name         `xml:"ArrayOfSearchLyricResult"`
	Text              string           `xml:",chardata"`
	Xsd               string           `xml:"xsd,attr"`
	Xsi               string           `xml:"xsi,attr"`
	Xmlns             string           `xml:"xmlns,attr"`
	SearchLyricResult []ChartLyricSong `xml:"SearchLyricResult"`
}

// La estructura de la respuesta de la API de ITunes
type ItunesSong struct {
	ArtistName      string  `json:"artistName,omitempty"`      // Nombre del artista
	CollectionName  string  `json:"collectionName,omitempty"`  // Nombre del album
	TrackName       string  `json:"trackName,omitempty"`       // Nombre de la cancion
	ArtworkURL100   string  `json:"artworkUrl100,omitempty"`   // URL de la imagen del album
	TrackTimeMillis int     `json:"trackTimeMillis,omitempty"` // Duracion de la cancion
	Currency        string  `json:"currency,omitempty"`        // Moneda
	TrackPrice      float64 `json:"trackPrice,omitempty"`      // Precio de la cancion
	TrackId         int     `json:"trackId,omitempty"`         // ID de la cancion
}

type ItunesResultDTO struct {
	ResultCount int          `json:"resultCount,omitempty"`
	Results     []ItunesSong `json:"results,omitempty"`
}
