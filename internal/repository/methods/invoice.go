package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
	"fmt"
	"strings"

	"database/sql"
)

type InvoiceDB struct {
	db *sql.DB
}

func NewInvoiceDB(db *sql.DB) *InvoiceDB {
	return &InvoiceDB{
		db: db,
	}
}

func (m *InvoiceDB) Get() (*[]models.Invoice, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s",
		models.InvoiceTable,
	)

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Invoice
	for rows.Next() {
		var invoice models.Invoice

		if err := rows.Scan(
			&invoice.InvoiceId,
			&invoice.CustomerId,
			&invoice.InvoiceDate,
			&invoice.BillingAddress,
			&invoice.BillingCity,
			&invoice.BillingState,
			&invoice.BillingCountry,
			&invoice.BillingPostalCode,
			&invoice.Total,
		); err != nil {
			return nil, err
		}
		results = append(results, invoice)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func (m *InvoiceDB) GetById(id int) (*models.Invoice, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE InvoiceId = ?",
		models.InvoiceTable,
	)

	var results models.Invoice
	if err := m.db.QueryRow(query, id).Scan(
		&results.InvoiceId,
		&results.CustomerId,
		&results.InvoiceDate,
		&results.BillingAddress,
		&results.BillingCity,
		&results.BillingState,
		&results.BillingCountry,
		&results.BillingPostalCode,
		&results.Total,
	); err != nil {
		return nil, err
	}

	return &results, nil
}

func (m *InvoiceDB) Create(invoice *models.InvoiceInput) (int, error) {
	if m.db == nil {
		return 0, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(`
        INSERT INTO %s 
            (CustomerId, InvoiceDate, BillingAddress, BillingCity, BillingState, BillingCountry, BillingPostalCode, Total)
        VALUES 
            (?, ?, ?, ?, ?, ?, ?, ?)`,
		models.InvoiceTable,
	)

	row, err := m.db.Exec(query,
		invoice.CustomerId,
		invoice.InvoiceDate,
		toNullString(invoice.BillingAddress),
		toNullString(invoice.BillingCity),
		toNullString(invoice.BillingState),
		toNullString(invoice.BillingCountry),
		toNullString(invoice.BillingPostalCode),
		toNullFloat64(invoice.Total),
	)
	if err != nil {
		return 0, err
	}

	invoiceId, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(invoiceId), nil
}

func (m *InvoiceDB) Put(invoice *models.InvoiceInput, id int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	setQuery, args := m.buildInvoiceUpdateSET(invoice)

	query := fmt.Sprintf(
		"UPDATE %s SET %s"+
			"WHERE InvoiceId = ?",
		models.InvoiceTable,
		setQuery,
	)
	args = append(args, id)

	_, err := m.db.Exec(query, args...)
	return err
}

func (m *InvoiceDB) Delete(id int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// InvoiceLine id
	if err := m.deleteInvoiceLine(tx, id); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get id InvoiceLine: %w", err)
	}

	// Invoice
	if err := m.deleteInvoice(tx, id); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete Invoice: %w ", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

/*
*	Auxiliary methods:
 */

func (m *InvoiceDB) buildInvoiceUpdateSET(invoice *models.InvoiceInput) (string, []interface{}) {
	var setValues []string
	var args []interface{}

	addField := func(field *string, columnName string) {
		if field != nil {
			setValues = append(setValues, columnName+" = ?")
			args = append(args, *field)
		}
	}

	addField(invoice.BillingAddress, "BillingAddress")
	addField(invoice.BillingCity, "BillingCity")
	addField(invoice.BillingState, "BillingState")
	addField(invoice.BillingCountry, "BillingCountry")
	addField(invoice.BillingPostalCode, "BillingPostalCode")

	if invoice.Total != nil {
		setValues = append(setValues, "CustomerId = ?")
		args = append(args, *invoice.CustomerId)
	}

	if invoice.CustomerId != nil {
		setValues = append(setValues, "CustomerId = ?")
		args = append(args, *invoice.CustomerId)
	}

	if invoice.InvoiceDate != nil {
		setValues = append(setValues, "InvoiceDate = ?")
		args = append(args, *invoice.InvoiceDate)
	}

	return strings.Join(setValues, ", "), args
}

func (m *InvoiceDB) deleteInvoiceLine(tx *sql.Tx, genreID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE InvoiceId = ?", models.InvoiceLineTable)

	_, err := tx.Exec(query, genreID)
	if err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}

func (m *InvoiceDB) deleteInvoice(tx *sql.Tx, genreID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE InvoiceId = ?", models.InvoiceTable)

	_, err := tx.Exec(query, genreID)
	if err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}
