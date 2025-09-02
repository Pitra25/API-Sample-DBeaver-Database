package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
	"fmt"

	"database/sql"
)

type PlayListDB struct {
	db *sql.DB
}

func NewPlayListDB(db *sql.DB) *PlayListDB {
	return &PlayListDB{
		db: db,
	}
}

func (m *PlayListDB) Get() (*[]models.Playlist, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s",
		models.PlaylistTable,
	)

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Playlist
	for rows.Next() {
		var playlist models.Playlist

		if err := rows.Scan(
			&playlist.PlaylistId,
			&playlist.Name,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		results = append(results, playlist)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func (m *PlayListDB) GetById(id int) (*models.Playlist, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE PlaylistId = ?",
		models.PlaylistTable,
	)

	var results models.Playlist
	if err := m.db.QueryRow(query, id).Scan(
		&results.PlaylistId,
		&results.Name,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &results, nil
}

func (m *PlayListDB) Create(playlist *models.PlaylistInput) (int, error) {
	if m.db == nil {
		return 0, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(`
        INSERT INTO %s 
            (Name)
        VALUES 
            (?)`,
		models.PlaylistTable,
	)

	row, err := m.db.Exec(query, playlist.Name)
	if err != nil {
		return 0, err
	}

	resultId, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(resultId), nil
}

func (m *PlayListDB) Put(playlist *models.PlaylistInput, id int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"UPDATE %s SET Name = ?"+
			"WHERE PlaylistId = ?",
		models.PlaylistTable,
	)

	_, err := m.db.Exec(query, playlist.Name, id)
	return err
}

func (m *PlayListDB) Delete(id int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// delete play list track
	if err := m.deletePlaylistTrack(tx, id); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get id play list track: %w", err)
	}

	// delete play list
	if err := m.deletePlaylist(tx, id); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete play list: %w ", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

/*
*	Auxiliary methods:
 */

func (m *PlayListDB) deletePlaylistTrack(tx *sql.Tx, genreID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE PlaylistId = ?", models.PlaylistTrackTable)

	_, err := tx.Exec(query, genreID)
	if err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}

func (m *PlayListDB) deletePlaylist(tx *sql.Tx, genreID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE PlaylistId = ?", models.PlaylistTable)

	_, err := tx.Exec(query, genreID)
	if err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}
