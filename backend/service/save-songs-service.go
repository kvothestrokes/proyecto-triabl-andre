package service

import (
	"PROYECTO/config"
	"PROYECTO/dto"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type SaveSongsService interface {
	SaveSongs(song []dto.SongSaveInfo) error
	SearchSongs(search dto.SearchRequestDTO) ([]dto.SearchResultDTO, error)
	InitiateDB() error
}

type saveSongsService struct {
	db *sql.DB
}

func NewSaveSongsService() SaveSongsService {
	return &saveSongsService{
		db: config.SetupDB(),
	}
}

func (i *saveSongsService) SaveSongs(song []dto.SongSaveInfo) error {

	db := i.db

	values := []string{}

	for _, song := range song {

		formatedPrice := fmt.Sprintf("%s %.2f", song.Currency, song.Price)

		//Convertir la duracion a formato legible
		miliseconds := int64(song.Duration)
		var t time.Time
		t = t.Add(time.Duration(miliseconds) * time.Millisecond)
		durationMinutes := t.Format("04:05")

		values = append(values, fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %d)", sanitizeString(song.ID), sanitizeString(song.Name), sanitizeString(song.Artist), sanitizeString(song.Album), durationMinutes, song.Artwork, formatedPrice, song.Origin, song.IDOrigin))
	}

	query := fmt.Sprintf("INSERT IGNORE INTO songs_list(id, song_name, artist_name, album_name, duration, artwork_url, price, origin, origin_id) VALUES %s", strings.Join(values, ","))

	// fmt.Println("\n Query: ", query)

	_, err := db.Exec(query)

	if err != nil {
		fmt.Printf("Error al guardar la cancion: %s", err.Error())
		return err
	}

	return nil
}

func sanitizeString(str string) string {
	result := strings.ReplaceAll(str, "'", "")
	result = strings.ReplaceAll(result, "\"", "")
	result = strings.ToLower(result)
	return result
}

func (i *saveSongsService) SearchSongs(search dto.SearchRequestDTO) ([]dto.SearchResultDTO, error) {

	var songs []dto.SearchResultDTO

	db := i.db

	params := ""
	if search.Name != "" {
		params = fmt.Sprintf(` LOWER(song_name) LIKE '%%%s%%' `, sanitizeString(search.Name))
	} else {
		return songs, fmt.Errorf("el nombre de la cancion es requerido %v", search)
	}

	if search.Artist != "" {
		params += fmt.Sprintf(` AND LOWER(artist_name) LIKE '%%%s%%'`, sanitizeString(search.Artist))
	}

	if search.Album != "" {
		params += fmt.Sprintf(` AND LOWER(album_name) LIKE '%%%s%%'`, sanitizeString(search.Album))
	}

	query := fmt.Sprintf("SELECT id, song_name, artist_name, album_name, duration, artwork_url, price, origin FROM songs_list WHERE %s ORDER BY song_name", params)

	// fmt.Printf("\n Query: %s", query)

	rows, err := db.Query(query)

	if err != nil {
		fmt.Printf("Error al buscar la cancion: %s", err.Error())
		return songs, err
	}

	for rows.Next() {
		var song dto.SearchResultDTO
		err := rows.Scan(&song.ID, &song.Name, &song.Artist, &song.Album, &song.Duration, &song.Artwork, &song.Price, &song.Origin)
		if err != nil {
			fmt.Printf("Error al buscar la cancion: %s", err.Error())
			return songs, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (i *saveSongsService) InitiateDB() error {
	db := i.db

	query := `CREATE TABLE IF NOT EXISTS songs_list(
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		song_name VARCHAR(255) NOT NULL,
		artist_name VARCHAR(255) NOT NULL,
		album_name VARCHAR(255) NOT NULL,
		duration VARCHAR(255) NOT NULL,
		artwork_url VARCHAR(255) NOT NULL,
		price VARCHAR(255) NOT NULL,
		origin VARCHAR(255) NOT NULL,
		origin_id INT NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("Error1 al crear la tabla: %s", err.Error())
		return err
	}
	return nil

}
