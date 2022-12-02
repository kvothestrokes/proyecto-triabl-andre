package dto

// La estructura de las canciones en la respuesta de la API de busqueda
type SearchResultDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Artist   string `json:"artist"`
	Duration string `json:"duration"`
	Album    string `json:"album"`
	Artwork  string `json:"artwork"`
	Price    string `json:"price"`
	Origin   string `json:"origin"`
}

// La estructura de la peticion de la API de busqueda
type SearchRequestDTO struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Token  string `json:"token"`
}

// Estructura con la que se guarda la cancion en la base de datos
type SongSaveInfo struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Artist   string  `json:"artist"`
	Duration int     `json:"duration"`
	Album    string  `json:"album"`
	Artwork  string  `json:"artwork"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
	Origin   string  `json:"origin"`
	IDOrigin int     `json:"idOrigin"`
}

//Estructura de la respuesta de la API de busqueda
type SearchResponseDTO struct {
	Status bool        `json:"status"`
	Error  string      `json:"Error"`
	Data   interface{} `json:"data"`
}

type TokenHanlder struct {
	Token string `json:"token"`
}
