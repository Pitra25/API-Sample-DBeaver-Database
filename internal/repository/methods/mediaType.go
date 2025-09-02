package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
	"fmt"
	"strings"

	"database/sql"
)

type MediaTypeDB struct {
	db *sql.DB
}

func NewMediaTypeDB(db *sql.DB) *MediaTypeDB {
	return &MediaTypeDB{
		db: db,
	}
}

func (m *MediaTypeDB) Get() (*[]models.MediaType, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s",
		models.MediaTypeTable,
	)

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.MediaType
	for rows.Next() {
		var mediaType models.MediaType

		if err := rows.Scan(
			&mediaType.MediaTypeId,
			&mediaType.Name,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		results = append(results, mediaType)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func (m *MediaTypeDB) GetById(id int) (*models.MediaType, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE MediaTypeId = ?",
		models.MediaTypeTable,
	)

	var results models.MediaType
	if err := m.db.QueryRow(query, id).Scan(
		&results.MediaTypeId,
		&results.Name,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &results, nil
}

func (m *MediaTypeDB) Create(mediaType *models.MediaTypeInput) (int, error) {
	if m.db == nil {
		return 0, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(`
        INSERT INTO %s 
            (Name)
        VALUES 
            (?)`,
		models.MediaTypeTable,
	)

	row, err := m.db.Exec(query, mediaType.Name)
	if err != nil {
		return 0, err
	}

	resultId, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(resultId), nil
}

func (m *MediaTypeDB) Put(mediaType *models.MediaTypeInput, id int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"UPDATE %s SET Name = ?"+
			"WHERE MediaTypeId = ?",
		models.MediaTypeTable,
	)

	_, err := m.db.Exec(query, mediaType.Name, id)
	return err
}

func (m *MediaTypeDB) Delete(id int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Get MediaType
	mediaTypeID, err := m.getTrackMediaType(tx, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get id media type: %w", err)
	}

	// Set null media type track
	if len(mediaTypeID) > 0 {
		if err := m.nullifyTrackMediaTypeReferences(tx, mediaTypeID); err != nil {
			return err
		}
	}

	// media type
	if err := m.deleteMediaType(tx, id); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete media type: %w ", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

/*
*	Auxiliary methods:
 */

func (m *MediaTypeDB) getTrackMediaType(tx *sql.Tx, mediaTypeID int) ([]int, error) {
	query := fmt.Sprintf(
		"SELECT TrackId FROM %s "+
			"WHERE MediaTypeId = ?",
		models.TrackTable,
	)

	rows, err := tx.Query(query, mediaTypeID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to get id invoice: %w", err)
	}

	var resultsID []int
	for rows.Next() {
		var trackID int

		if err := rows.Scan(&trackID); err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to scan invoice ID: %w", err)
		}

		resultsID = append(resultsID, trackID)
	}

	if err = rows.Err(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error iterating invoice rows: %w", err)
	}

	return resultsID, nil
}

func (m *MediaTypeDB) nullifyTrackMediaTypeReferences(tx *sql.Tx, mediaTypesID []int) error {

	placeholders := strings.Repeat("?, ", len(mediaTypesID)-1) + "?"

	query := fmt.Sprintf(
		"UPDATE %s "+
			"SET MediaTypeId = NULL WHERE MediaTypeId IN (%s)",
		models.TrackTable, placeholders,
	)

	args := make([]interface{}, len(mediaTypesID))
	for i, v := range mediaTypesID {
		args[i] = v
	}

	_, err := tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update MediaTypeId track: %w", err)
	}

	return nil
}

func (m *MediaTypeDB) deleteMediaType(tx *sql.Tx, mediaTypeID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE MediaTypeId = ?", models.MediaTypeTable)

	_, err := tx.Exec(query, mediaTypeID)
	if err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}
