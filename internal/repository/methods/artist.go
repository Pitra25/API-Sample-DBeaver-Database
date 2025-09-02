package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
	"errors"
	"fmt"
	"strings"

	"database/sql"
)

type ArtistDB struct {
	db *sql.DB
}

func NewArtistDB(db *sql.DB) *ArtistDB {
	return &ArtistDB{
		db: db,
	}
}

func (m *ArtistDB) Get() (*[]models.Artist, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s",
		models.ArtistTable,
	)

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Artist
	for rows.Next() {
		var artist models.Artist

		if err := rows.Scan(
			&artist.ArtistId,
			&artist.Name,
		); err != nil {
			return nil, err
		}
		results = append(results, artist)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func (m *ArtistDB) GetById(id int) (*models.Artist, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE ArtistId = ?",
		models.ArtistTable,
	)

	var results models.Artist
	if err := m.db.QueryRow(query, id).Scan(
		&results.ArtistId,
		&results.Name,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &results, nil
}

func (m *ArtistDB) GetByName(title *string) (*models.Artist, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s",
		models.ArtistTable,
	)

	var results models.Artist
	if err := m.db.QueryRow(query).Scan(
		&results.ArtistId,
		&results.Name,
	); err != nil {
		return nil, err
	}

	return &results, nil
}

func (m *ArtistDB) Create(artist *models.ArtistInput) (int, error) {
	if m.db == nil {
		return 0, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (Name) VALUES (?)",
		models.ArtistTable,
	)

	row, err := m.db.Exec(query, artist.Name)
	if err != nil {
		return 0, err
	}

	artistId, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(artistId), nil
}

func (m *ArtistDB) Put(artist *models.ArtistInput, id int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"UPDATE %s "+
			"SET Name = ? "+
			"WHERE ArtistId = ?",
		models.ArtistTable,
	)

	_, err := m.db.Exec(query, artist.Name, id)

	return err
}

func (m *ArtistDB) Delete(idArtist int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// AlumID
	albumIDs, err := m.getArtistAlbums(tx, idArtist)
	if err != nil {
		return err
	}

	// Track
	if len(albumIDs) > 0 {
		if err := m.nullifyTrackAlbumReferences(tx, albumIDs); err != nil {
			return err
		}
	}

	// Album
	if err := m.deleteArtistAlbums(tx, idArtist); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete album: %w", err)
	}

	// Artist
	if err := m.deleteArtist(tx, idArtist); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete artist: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

/*
*	Auxiliary methods:
 */

func (m *ArtistDB) getArtistAlbums(tx *sql.Tx, artistID int) ([]int, error) {
	query := fmt.Sprintf(
		"SELECT AlbumId FROM %s "+
			"WHERE ArtistId = ?",
		models.AlbumTable,
	)

	rows, err := tx.Query(query, artistID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to get id alums: %w", err)
	}

	var idAlbums []int
	for rows.Next() {
		var idAlbum int

		if err := rows.Scan(&idAlbum); err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to scan album ID: %w", err)
		}

		idAlbums = append(idAlbums, idAlbum)
	}

	if err = rows.Err(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error iterating album rows: %w", err)
	}

	return idAlbums, nil
}

func (m *ArtistDB) nullifyTrackAlbumReferences(tx *sql.Tx, albumsID []int) error {

	placeholders := strings.Repeat("?, ", len(albumsID)-1) + "?"

	query := fmt.Sprintf(
		"UPDATE %s "+
			"SET AlbumId = NULL WHERE AlbumId IN (%s)",
		models.TrackTable, placeholders,
	)

	args := make([]interface{}, len(albumsID))
	for i, v := range albumsID {
		args[i] = v
	}

	_, err := tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete track: %w", err)
	}

	return nil
}

func (m *ArtistDB) deleteArtistAlbums(tx *sql.Tx, artistID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE ArtistId = ?", models.AlbumTable)

	_, err := tx.Exec(query, artistID)
	if err != nil {
		return fmt.Errorf("failed to delete albums: %w", err)
	}

	return nil
}

func (m *ArtistDB) deleteArtist(tx *sql.Tx, id int) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE ArtistId = ?",
		models.ArtistTable,
	)

	_, err := tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete artist: %w", err)
	}

	return nil
}
