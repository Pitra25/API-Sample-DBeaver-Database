package models

import "errors"

type Customer struct {
	CustomerId   *int    `json:"customerId" db:"CustomerId"`
	FirstName    *string `json:"firstName" db:"FirstName"`
	LastName     *string `json:"lastName" db:"LastName"`
	Company      *string `json:"company" db:"Company"`
	Address      *string `json:"address" db:"Address"`
	City         *string `json:"city" db:"City"`
	State        *string `json:"state" db:"State"`
	Country      *string `json:"country" db:"Country"`
	PostalCode   *string `json:"postalCode" db:"PostalCode"`
	Phone        *string `json:"phone" db:"Phone"`
	Fax          *string `json:"fax" db:"Fax"`
	Email        *string `json:"email" db:"Email"`
	SupportRepId *int    `json:"supportRepId" db:"SupportRepId"`
}

type CustomerInput struct {
	FirstName    *string `json:"firstName" db:"FirstName" binding:"required"`
	LastName     *string `json:"lastName" db:"LastName" binding:"required"`
	Company      *string `json:"company" db:"Company"`
	Address      *string `json:"address" db:"Address"`
	City         *string `json:"city" db:"City"`
	State        *string `json:"state" db:"State"`
	Country      *string `json:"country" db:"Country"`
	PostalCode   *string `json:"postalCode" db:"PostalCode"`
	Phone        *string `json:"phone" db:"Phone"`
	Fax          *string `json:"fax" db:"Fax"`
	Email        *string `json:"email" db:"Email" binding:"required"`
	SupportRepId *int    `json:"supportRepId" db:"SupportRepId"`
}

func (c *CustomerInput) Validate() error {
	if c.FirstName == nil || c.LastName == nil || c.Email == nil {
		return errors.New("input structure has no values")
	}
	return nil
}
