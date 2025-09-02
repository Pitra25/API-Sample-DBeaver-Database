package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
	"errors"
	"fmt"
	"strings"

	"database/sql"
)

type CustomerDB struct {
	db *sql.DB
}

func NewCustomerDB(db *sql.DB) *CustomerDB {
	return &CustomerDB{
		db: db,
	}
}

func (m *CustomerDB) Get() (*[]models.Customer, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s",
		models.CustomerTable,
	)

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Customer
	for rows.Next() {
		var customer models.Customer

		if err := rows.Scan(
			&customer.CustomerId,
			&customer.FirstName,
			&customer.LastName,
			&customer.Company,
			&customer.Address,
			&customer.City,
			&customer.State,
			&customer.Country,
			&customer.PostalCode,
			&customer.Phone,
			&customer.Fax,
			&customer.Email,
			&customer.SupportRepId,
		); err != nil {
			return nil, err
		}
		results = append(results, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func (m *CustomerDB) GetById(id int) (*models.Customer, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE CustomerId = ?",
		models.CustomerTable,
	)

	var results models.Customer
	if err := m.db.QueryRow(query, id).Scan(
		&results.CustomerId,
		&results.FirstName,
		&results.LastName,
		&results.Company,
		&results.Address,
		&results.City,
		&results.State,
		&results.Country,
		&results.PostalCode,
		&results.Phone,
		&results.Fax,
		&results.Email,
		&results.SupportRepId,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &results, nil
}

func (m *CustomerDB) Create(customer *models.CustomerInput) (int, error) {
	if m.db == nil {
		return 0, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(`
        INSERT INTO %s 
            (FirstName, LastName, Company, Address, City, State, Country, PostalCode, Phone, Fax, Email, SupportRepId)
        VALUES 
            (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		models.CustomerTable,
	)

	row, err := m.db.Exec(query,
		customer.FirstName,
		customer.LastName,
		toNullString(customer.Company),
		toNullString(customer.Address),
		toNullString(customer.City),
		toNullString(customer.State),
		toNullString(customer.Country),
		toNullString(customer.PostalCode),
		toNullString(customer.Phone),
		toNullString(customer.Fax),
		customer.Email,
		toNullInt64(customer.SupportRepId),
	)
	if err != nil {
		return 0, err
	}

	artistId, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(artistId), nil
}

func (m *CustomerDB) Put(customer *models.CustomerInput, id int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	setQuery, args := m.buildCustomerUpdateSET(customer)

	query := fmt.Sprintf(
		"UPDATE %s SET %s"+
			"WHERE CustomerId = ?",
		models.CustomerTable,
		setQuery,
	)
	args = append(args, id)

	_, err := m.db.Exec(query, args...)
	return err
}

func (m *CustomerDB) Delete(id int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Get InvoiceID
	invoiceID, err := m.getInvoicesCustomer(tx, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get id invoice: %w", err)
	}

	// InvoiceLine
	if len(invoiceID) > 0 {
		if err := m.deleteInvoiceLine(tx, invoiceID); err != nil {
			return err
		}
	}

	// Invoice
	if err := m.deleteInvoiceCustomer(tx, id); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete invoice customer: %w", err)
	}

	// Customer
	if err := m.deleteCustomer(tx, id); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete customer: %w ", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

/*
*	Auxiliary methods:
 */

func (m *CustomerDB) buildCustomerUpdateSET(customer *models.CustomerInput) (string, []interface{}) {
	var setValues []string
	var args []interface{}

	addField := func(field *string, columnName string) {
		if field != nil {
			setValues = append(setValues, columnName+" = ?")
			args = append(args, *field)
		}
	}

	addField(customer.FirstName, "FirstName")
	addField(customer.LastName, "LastName")
	addField(customer.Company, "Company")
	addField(customer.Address, "Address")
	addField(customer.City, "City")
	addField(customer.State, "State")
	addField(customer.Country, "Country")
	addField(customer.PostalCode, "PostalCode")
	addField(customer.Phone, "Phone")
	addField(customer.Fax, "Fax")
	addField(customer.Email, "Email")

	if customer.SupportRepId != nil {
		setValues = append(setValues, "SupportRepId = ?")
		args = append(args, *customer.SupportRepId)
	}

	return strings.Join(setValues, ", "), args
}

func (m *CustomerDB) getInvoicesCustomer(tx *sql.Tx, customerID int) ([]int, error) {
	query := fmt.Sprintf(
		"SELECT AlbumId FROM %s "+
			"WHERE CustomerId = ?",
		models.InvoiceTable,
	)

	rows, err := tx.Query(query, customerID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to get id invoice: %w", err)
	}

	var customersID []int
	for rows.Next() {
		var idCustomer int

		if err := rows.Scan(&idCustomer); err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to scan invoice ID: %w", err)
		}

		customersID = append(customersID, idCustomer)
	}

	if err = rows.Err(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error iterating invoice rows: %w", err)
	}

	return customersID, nil
}

func (m *CustomerDB) deleteInvoiceLine(tx *sql.Tx, customerID []int) error {

	placeholders := strings.Repeat("?, ", len(customerID)-1) + "?"

	query := fmt.Sprintf(
		"DELETE FROM %s "+
			"WHERE InvoiceId IN (%s)",
		models.InvoiceLineTable, placeholders,
	)

	args := make([]interface{}, len(customerID))
	for i, v := range customerID {
		args[i] = v
	}

	_, err := tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete invoice line: %w", err)
	}

	return nil
}

func (m *CustomerDB) deleteInvoiceCustomer(tx *sql.Tx, customerID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE CustomerId = ?", models.InvoiceTable)

	_, err := tx.Exec(query, customerID)
	if err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}

func (m *CustomerDB) deleteCustomer(tx *sql.Tx, customerID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE CustomerId = ?", models.CustomerTable)

	_, err := tx.Exec(query, customerID)
	if err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}
