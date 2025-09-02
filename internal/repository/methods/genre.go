package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
	"fmt"
	"strings"

	"database/sql"
)

type GenreDB struct {
	db *sql.DB
}

func NewGenreDB(db *sql.DB) *GenreDB {
	return &GenreDB{
		db: db,
	}
}

func (m *GenreDB) Get() (*[]models.Genre, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s",
		models.GenreTable,
	)

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Genre
	for rows.Next() {
		var genre models.Genre

		if err := rows.Scan(
			&genre.GenreId,
			&genre.Name,
		); err != nil {
			return nil, err
		}
		results = append(results, genre)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func (m *GenreDB) GetById(id int) (*models.Genre, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE GenreId = ?",
		models.GenreTable,
	)

	var results models.Genre
	if err := m.db.QueryRow(query, id).Scan(
		&results.GenreId,
		&results.Name,
	); err != nil {
		return nil, err
	}

	return &results, nil
}

func (m *GenreDB) Create(genre *models.GenreInput) (int, error) {
	if m.db == nil {
		return 0, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(`
        INSERT INTO %s 
            (Name)
        VALUES 
            (?)`,
		models.GenreTable,
	)

	row, err := m.db.Exec(query, genre.Name)
	if err != nil {
		return 0, err
	}

	resultId, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(resultId), nil
}

func (m *GenreDB) Put(genre *models.GenreInput, id int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"UPDATE %s SET Name = ?"+
			"WHERE GenreId = ?",
		models.EmployeeTable,
	)

	_, err := m.db.Exec(query, genre.Name, id)
	return err
}

func (m *GenreDB) Delete(id int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Get Track
	trackID, err := m.getTrackGenre(tx, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get id customers: %w", err)
	}

	// Set null genre track
	if len(trackID) > 0 {
		if err := m.nullifyTrackGenreReferences(tx, trackID); err != nil {
			return err
		}
	}

	// Employee
	if err := m.deleteGenre(tx, id); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete employee: %w ", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

/*
*	Auxiliary methods:
 */

func (m *GenreDB) getTrackGenre(tx *sql.Tx, genreID int) ([]int, error) {
	query := fmt.Sprintf(
		"SELECT TrackId FROM %s "+
			"WHERE GenreId = ?",
		models.TrackTable,
	)

	rows, err := tx.Query(query, genreID)
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

func (m *GenreDB) nullifyTrackGenreReferences(tx *sql.Tx, tracksID []int) error {

	placeholders := strings.Repeat("?, ", len(tracksID)-1) + "?"

	query := fmt.Sprintf(
		"UPDATE %s "+
			"SET GenreId = NULL WHERE GenreId IN (%s)",
		models.TrackTable, placeholders,
	)

	args := make([]interface{}, len(tracksID))
	for i, v := range tracksID {
		args[i] = v
	}

	_, err := tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update SupportRepId customer: %w", err)
	}

	return nil
}

func (m *GenreDB) deleteGenre(tx *sql.Tx, genreID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE CustomerId = ?", models.GenreTable)

	_, err := tx.Exec(query, genreID)
	if err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}
