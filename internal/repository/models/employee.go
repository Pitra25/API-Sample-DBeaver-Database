package models

import (
	"errors"
	"time"
)

type Employee struct {
	EmployeeId *int       `json:"employeeId" db:"EmployeeId"`
	LastName   *string    `json:"lastName" db:"LastName"`
	FirstName  *string    `json:"firstName" db:"FirstName"`
	Title      *string    `json:"title" db:"Title"`
	ReportsTo  *int       `json:"reportsTo" db:"ReportsTo"`
	BirthDate  *time.Time `json:"birthDate" db:"BirthDate"`
	HireDate   *time.Time `json:"hireDate" db:"HireDate"`
	Address    *string    `json:"address" db:"Address"`
	City       *string    `json:"city" db:"City"`
	State      *string    `json:"state" db:"State"`
	Country    *string    `json:"country" db:"Country"`
	PostalCode *string    `json:"postalCode" db:"PostalCode"`
	Phone      *string    `json:"phone" db:"Phone"`
	Fax        *string    `json:"fax" db:"Fax"`
	Email      *string    `json:"email" db:"Email"`
}

type EmployeeInput struct {
	LastName   *string    `json:"lastName" db:"LastName" binding:"required"`
	FirstName  *string    `json:"firstName" db:"FirstName" binding:"required"`
	Title      *string    `json:"title" db:"Title"`
	ReportsTo  *int       `json:"reportsTo" db:"ReportsTo"`
	BirthDate  *time.Time `json:"birthDate" db:"BirthDate"`
	HireDate   *time.Time `json:"hireDate" db:"HireDate"`
	Address    *string    `json:"address" db:"Address"`
	City       *string    `json:"city" db:"City"`
	State      *string    `json:"state" db:"State"`
	Country    *string    `json:"country" db:"Country"`
	PostalCode *string    `json:"postalCode" db:"PostalCode"`
	Phone      *string    `json:"phone" db:"Phone"`
	Fax        *string    `json:"fax" db:"Fax"`
	Email      *string    `json:"email" db:"Email"`
}

func (e *EmployeeInput) Validate() error {
	if e.LastName == nil || e.FirstName == nil {
		return errors.New("input structure has no values")
	}
	return nil
}
