package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
	"fmt"
	"strings"

	"database/sql"

	"github.com/sirupsen/logrus"
)

type AlbumDB struct {
	db *sql.DB
}

func NewAlbumDB(db *sql.DB) *AlbumDB {
	return &AlbumDB{
		db: db,
	}
}

func (m *AlbumDB) Get() (*[]models.Album, error) {
	logrus.Debug("repository")

	if m.db == nil {
		logrus.Debug("db is nil")
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s",
		models.AlbumTable,
	)

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Album
	for rows.Next() {
		var album models.Album

		if err := rows.Scan(
			&album.AlbumId,
			&album.Title,
			&album.ArtistId,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		results = append(results, album)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func (m *AlbumDB) GetById(id int) (*models.Album, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s a WHERE a.AlbumId = %d",
		models.AlbumTable,
		id,
	)

	var results models.Album
	if err := m.db.QueryRow(query).Scan(
		&results.AlbumId,
		&results.Title,
		&results.ArtistId,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &results, nil
}

func (m *AlbumDB) GetByName(title *string) (*models.Album, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * from %s where title = %s",
		models.AlbumTable,
		*title,
	)

	var results models.Album
	if err := m.db.QueryRow(query).Scan(
		&results.AlbumId,
		&results.Title,
		&results.ArtistId,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &results, nil
}

func (m *AlbumDB) Create(album *models.AlbumInput) (int, error) {

	tx, err := m.db.Begin()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (Title, ArtistId) VALUES (?, ?)",
		models.AlbumTable,
	)

	row, err := tx.Exec(query, album.Title, album.ArtistId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	albumId, err := row.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(albumId), tx.Commit()
}

func (m *AlbumDB) Put(album *models.AlbumInput, albumId int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if album.Title != nil {
		setValues = append(setValues, "al.title = ?")
		args = append(args, *album.Title)
	}

	if album.ArtistId != nil {
		setValues = append(setValues, "al.ArtistId = ?")
		args = append(args, *album.ArtistId)
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		"UPDATE %s al "+
			"SET %s "+
			"WHERE al.AlbumId = ?",
		models.AlbumTable,
		setQuery,
	)
	args = append(args, albumId)

	_, err := m.db.Exec(query, args...)

	return err
}

func (m *AlbumDB) Delete(idAlbum, idArtist int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := m.nullifyCustomerEmployeeReferences(tx, idAlbum); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get id customers: %w", err)
	}

	if err := m.deleteAlbum(tx, idAlbum, idArtist); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete album: %w ", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

/*
*	Auxiliary methods:
 */

func (m *AlbumDB) nullifyCustomerEmployeeReferences(tx *sql.Tx, albumID int) error {
	query := fmt.Sprintf(
		"UPDATE %s "+
			"SET AlbumId = NULL WHERE AlbumId = ?",
		models.TrackTable,
	)

	_, err := tx.Exec(query, albumID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (m *AlbumDB) deleteAlbum(tx *sql.Tx, albumID, idArtist int) error {
	query := fmt.Sprintf(
		"DELETE al FROM %s al "+
			"WHERE al.AlbumId = ? AND al.ArtistId = ?",
		models.AlbumTable,
	)

	_, err := tx.Exec(query, albumID, idArtist)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
