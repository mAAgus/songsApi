package storage

import (
	"fmt"
	"go/scr/hhruxongs/models"
	"log"
)

type SongsRepository struct {
	storage *Storage
}

var (
	tableSongs string = "songs"
)

func (song *SongsRepository) Create(sg *models.Songs) (*models.Songs, error) {
	query := fmt.Sprintf("INSERT INTO %s (group, song) VALUES ($1, $2) RETURNING id", tableSongs)

	if err := song.storage.db.QueryRow(query, sg.Group, sg.Song).Scan(&sg.ID); err != nil {
		return nil, err
	}
	return sg, nil

}
func (song *SongsRepository) GetSongsParts(id int, page int, pageSize int) (*models.Songs, bool, error) {
	son, ok, err := song.FindById(id)
	if err != nil {
		return nil, false, err
	}

	if !ok {
		return nil, false, nil
	}

	var countQuery = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE song_id = %d", tableSongs, id)
	var count int64

	err = song.storage.db.QueryRow(countQuery, id).Scan(&count)
	if err != nil {
		return nil, false, err
	}

	offset := (page - 1) * pageSize

	query := fmt.Sprintf("SELECT text_parts FROM %s WHERE song_id = %d LIMIT %d OFFSET %d", tableSongs, id, pageSize, offset)
	rows, err := song.storage.db.Query(query, id, pageSize, offset)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()

	var textParts string
	for _, part := range son.TextParts {
		textParts += string(part) + " "
	}

	son.TextParts = textParts

	return son, true, nil
}

func (song *SongsRepository) DeleteSongById(id int) (*models.Songs, error) {
	son, ok, err := song.FindById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableSongs)
		_, err := song.storage.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}
	return son, nil
}
func (song *SongsRepository) UpdteSongs(id int, upSong *models.Songs) (*models.Songs, bool, error) {
	var founded bool
	_, ok, err := song.FindById(id)
	if err != nil {
		return nil, founded, err
	}
	if ok {
		query := fmt.Sprintf("UPDATE %s SET release_date = $2, group = $3, song = $4, textSongs = $5, link = $6 WHERE id = $1", tableSongs)
		_, err := song.storage.db.Exec(query, id, upSong.Release_data, upSong.Group, upSong.Song, upSong.TextParts, upSong.Link)
		if err != nil {
			return nil, founded, err
		}
		founded = true
	}
	return upSong, founded, nil
}
func (song *SongsRepository) FindById(id int) (*models.Songs, bool, error) {
	son, err := song.SelectAll()
	var found bool
	if err != nil {
		return nil, found, err
	}
	var songFind *models.Songs
	for _, a := range son {
		if a.ID == id {
			songFind = a
			found = true
			break
		}
	}
	return songFind, found, nil
}
func (song *SongsRepository) SelectAll() ([]*models.Songs, error) {
	query := fmt.Sprintf("SELECT * FROM %s ", tableSongs)
	rows, err := song.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	songs := make([]*models.Songs, 0)
	for rows.Next() {
		a := models.Songs{}
		err := rows.Scan(&a.ID, a.Release_data, a.Group, a.Song, a.TextParts, a.Link)
		if err != nil {
			log.Println(err)
			continue
		}
		songs = append(songs, &a)
	}
	return songs, nil
}
func (song *SongsRepository) SelectAllFiltrPangination(filter string, page int, pageSize int) ([]*models.Songs, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE 1=1", tableSongs)

	if filter != "" {
		query += fmt.Sprintf(" AND (%s)", filter)
	}

	query += fmt.Sprintf(" ORDER BY ID LIMIT %d OFFSET %d", pageSize, page*pageSize)

	rows, err := song.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	songs := make([]*models.Songs, 0)
	for rows.Next() {
		song := models.Songs{}
		err := rows.Scan(&song.ID, &song.Release_data, &song.Group, &song.Song, &song.TextParts, &song.Link)
		if err != nil {
			log.Println(err)
			continue
		}
		songs = append(songs, &song)
	}

	return songs, nil
}
