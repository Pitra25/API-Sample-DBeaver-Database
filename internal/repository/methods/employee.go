package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
	"errors"
	"fmt"
	"strings"
	"time"

	"database/sql"
)

type EmployeeDB struct {
	db *sql.DB
}

func NewEmployeeDB(db *sql.DB) *EmployeeDB {
	return &EmployeeDB{
		db: db,
	}
}

func (m *EmployeeDB) Get() (*[]models.Employee, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s",
		models.EmployeeTable,
	)

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Employee
	for rows.Next() {
		var employee models.Employee

		if err := rows.Scan(
			&employee.EmployeeId,
			&employee.LastName,
			&employee.FirstName,
			&employee.Title,
			&employee.ReportsTo,
			&employee.BirthDate,
			&employee.HireDate,
			&employee.Address,
			&employee.City,
			&employee.State,
			&employee.Country,
			&employee.PostalCode,
			&employee.Phone,
			&employee.Fax,
			&employee.Email,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		results = append(results, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func (m *EmployeeDB) GetById(id int) (*models.Employee, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE EmployeeId = ?",
		models.EmployeeTable,
	)

	var results models.Employee
	if err := m.db.QueryRow(query, id).Scan(
		&results.EmployeeId,
		&results.LastName,
		&results.FirstName,
		&results.Title,
		&results.ReportsTo,
		&results.BirthDate,
		&results.HireDate,
		&results.Address,
		&results.City,
		&results.State,
		&results.Country,
		&results.PostalCode,
		&results.Phone,
		&results.Fax,
		&results.Email,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &results, nil
}

func (m *EmployeeDB) Create(employee *models.EmployeeInput) (int, error) {
	if m.db == nil {
		return 0, fmt.Errorf("db is nil")
	}

	query := fmt.Sprintf(`
        INSERT INTO %s 
            (LastName, FirstName, Title, BirthDate, HireDate, Address, City, State, Country, PostalCode, Phone, Fax, Email)
        VALUES 
            (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		models.EmployeeTable,
	)

	row, err := m.db.Exec(query,
		employee.LastName,
		employee.FirstName,
		toNullString(employee.Title),
		toNullInt64(employee.ReportsTo),
		toNullData(employee.BirthDate),
		toNullData(employee.HireDate),
		toNullString(employee.Address),
		toNullString(employee.City),
		toNullString(employee.State),
		toNullString(employee.Country),
		toNullString(employee.PostalCode),
		toNullString(employee.Phone),
		toNullString(employee.Fax),
		employee.Email,
	)
	if err != nil {
		return 0, err
	}

	employeeId, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(employeeId), nil
}

func (m *EmployeeDB) Put(employee *models.EmployeeInput, id int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	setQuery, args := m.buildEmployeeUpdateSET(employee)

	query := fmt.Sprintf(
		"UPDATE %s SET %s"+
			"WHERE EmployeeId = ?",
		models.EmployeeTable,
		setQuery,
	)
	args = append(args, id)

	_, err := m.db.Exec(query, args...)
	return err
}

func (m *EmployeeDB) Delete(id int) error {
	if m.db == nil {
		return fmt.Errorf("db is nil")
	}

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Customer
	customersID, err := m.getCustomerEmployee(tx, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get id customers: %w", err)
	}

	// InvoiceLine
	if len(customersID) > 0 {
		if err := m.nullifyCustomerEmployeeReferences(tx, customersID); err != nil {
			return err
		}
	}

	// Employee
	if err := m.deleteEmployee(tx, id); err != nil {
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

func (m *EmployeeDB) buildEmployeeUpdateSET(customer *models.EmployeeInput) (string, []interface{}) {
	var setValues []string
	var args []interface{}

	addField := func(field *string, columnName string) {
		if field != nil {
			setValues = append(setValues, columnName+" = ?")
			args = append(args, *field)
		}
	}

	addFieldTime := func(field *time.Time, columnName string) {
		if field != nil {
			setValues = append(setValues, columnName+" = ?")
			args = append(args, *field)
		}
	}

	addField(customer.LastName, "LastName")
	addField(customer.FirstName, "FirstName")
	addField(customer.Title, "Title")
	addFieldTime(customer.BirthDate, "BirthDate")
	addFieldTime(customer.HireDate, "HireDate")
	addField(customer.Address, "Address")
	addField(customer.City, "City")
	addField(customer.State, "State")
	addField(customer.Country, "Country")
	addField(customer.PostalCode, "PostalCode")
	addField(customer.Phone, "Phone")
	addField(customer.Fax, "Fax")
	addField(customer.Email, "Email")

	if customer.ReportsTo != nil {
		setValues = append(setValues, "ReportsTo = ?")
		args = append(args, *customer.ReportsTo)
	}

	return strings.Join(setValues, ", "), args
}

func (m *EmployeeDB) getCustomerEmployee(tx *sql.Tx, employeeID int) ([]int, error) {
	query := fmt.Sprintf(
		"SELECT CustomerId FROM %s "+
			"WHERE SupportRepId = ?",
		models.CustomerTable,
	)

	rows, err := tx.Query(query, employeeID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to get id invoice: %w", err)
	}

	var resultsID []int
	for rows.Next() {
		var idEmployee int

		if err := rows.Scan(&idEmployee); err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to scan invoice ID: %w", err)
		}

		resultsID = append(resultsID, idEmployee)
	}

	if err = rows.Err(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error iterating invoice rows: %w", err)
	}

	return resultsID, nil
}

func (m *EmployeeDB) nullifyCustomerEmployeeReferences(tx *sql.Tx, customersID []int) error {

	placeholders := strings.Repeat("?, ", len(customersID)-1) + "?"

	query := fmt.Sprintf(
		"UPDATE %s "+
			"SET SupportRepId = NULL WHERE SupportRepId IN (%s)",
		models.CustomerTable, placeholders,
	)

	args := make([]interface{}, len(customersID))
	for i, v := range customersID {
		args[i] = v
	}

	_, err := tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update SupportRepId customer: %w", err)
	}

	return nil
}

func (m *EmployeeDB) deleteEmployee(tx *sql.Tx, employeeID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE CustomerId = ?", models.EmployeeTable)

	_, err := tx.Exec(query, employeeID)
	if err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}
