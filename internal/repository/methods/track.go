package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
	"fmt"
	"strings"

	"database/sql"
)

type TrackDB struct {
	db *sql.DB
}

func NewTrackDB(db *sql.DB) *TrackDB {
	return &TrackDB{
		db: db,
	}
}

func (m *TrackDB) Get() (*[]models.Track, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s",
		models.TrackTable,
	)

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Track
	for rows.Next() {
		var track models.Track

		if err := rows.Scan(
			&track.TrackId,
			&track.Name,
			&track.AlbumId,
			&track.MediaTypeId,
			&track.GenreId,
			&track.Composer,
			&track.Milliseconds,
			&track.Bytes,
			&track.UnitPrice,
		); err != nil {
			return nil, err
		}
		results = append(results, track)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func (m *TrackDB) GetById(id int) (*models.Track, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE TrackId = ?",
		models.TrackTable,
	)

	var results models.Track
	if err := m.db.QueryRow(query, id).Scan(
		&results.TrackId,
		&results.Name,
		&results.AlbumId,
		&results.MediaTypeId,
		&results.GenreId,
		&results.Composer,
		&results.Milliseconds,
		&results.Bytes,
		&results.UnitPrice,
	); err != nil {
		return nil, err
	}

	return &results, nil
}

func (m *TrackDB) Create(track *models.TrackInput) (int, error) {
	if m.db == nil {
		return 0, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(`
        INSERT INTO %s 
            (Name, AlbumId, MediaTypeId, GenreId, Composer, Milliseconds, Bytes, UnitPrice)
        VALUES 
            (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		models.PlaylistTable,
	)

	row, err := m.db.Exec(query,
		track.Name,
		toNullInt64(track.AlbumId),
		track.MediaTypeId,
		toNullInt64(track.GenreId),
		toNullString(track.Composer),
		track.Milliseconds,
		toNullInt64(track.Bytes),
		track.UnitPrice,
	)
	if err != nil {
		return 0, err
	}

	resultId, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(resultId), nil
}

func (m *TrackDB) Put(track *models.TrackInput, id int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	setQuery, args := m.buildTrackUpdateSET(track)

	query := fmt.Sprintf(
		"UPDATE %s SET %s"+
			"WHERE TrackId = ?",
		models.TrackTable,
		setQuery,
	)
	args = append(args, id)

	_, err := m.db.Exec(query, args...)
	return err
}

func (m *TrackDB) Delete(id int) error {
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

	// delete invoice line
	if err := m.deleteInvoiceLine(tx, id); err != nil {
		return err
	}

	// delete track
	if err := m.deleteTrack(tx, id); err != nil {
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

func (m *TrackDB) buildTrackUpdateSET(track *models.TrackInput) (string, []interface{}) {
	var setValues []string
	var args []interface{}

	addField := func(field *string, columnName string) {
		if field != nil {
			setValues = append(setValues, columnName+" = ?")
			args = append(args, *field)
		}
	}

	addFieldInt := func(field *int, columnName string) {
		if field != nil {
			setValues = append(setValues, columnName+" = ?")
			args = append(args, *field)
		}
	}

	addField(track.Name, "Name")
	addFieldInt(track.AlbumId, "AlbumId")
	addFieldInt(track.MediaTypeId, "MediaTypeId")
	addFieldInt(track.GenreId, "GenreId")
	addField(track.Composer, "Composer")
	addFieldInt(track.Milliseconds, "Milliseconds")
	addFieldInt(track.Bytes, "Bytes")

	if track.UnitPrice != nil {
		setValues = append(setValues, "UnitPrice = ?")
		args = append(args, *track.UnitPrice)
	}

	return strings.Join(setValues, ", "), args
}

func (m *TrackDB) deletePlaylistTrack(tx *sql.Tx, trackID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE TrackId = ?", models.PlaylistTrackTable)

	_, err := tx.Exec(query, trackID)
	if err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}

func (m *TrackDB) deleteTrack(tx *sql.Tx, trackID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE PlaylistId = ?", models.TrackTable)

	_, err := tx.Exec(query, trackID)
	if err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}

func (m *TrackDB) deleteInvoiceLine(tx *sql.Tx, trackID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE TrackId = ?", models.InvoiceLineTable)

	_, err := tx.Exec(query, trackID)
	if err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}
