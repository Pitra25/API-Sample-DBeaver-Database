package models

import (
	"errors"
	"time"
)

type Invoice struct {
	InvoiceId         *int       `json:"invoiceId" db:"InvoiceId"`
	CustomerId        *int       `json:"customerId" db:"CustomerId"`
	InvoiceDate       *time.Time `json:"invoiceDate" db:"InvoiceDate"`
	BillingAddress    *string    `json:"billingAddress" db:"BillingAddress"`
	BillingCity       *string    `json:"billingCity" db:"BillingCity"`
	BillingState      *string    `json:"billingState" db:"BillingState"`
	BillingCountry    *string    `json:"billingCountry" db:"BillingCountry"`
	BillingPostalCode *string    `json:"billingPostalCode" db:"BillingPostalCode"`
	Total             *float64   `json:"total" db:"Total"`
}

type InvoiceInput struct {
	CustomerId        *int       `json:"customerId" db:"CustomerId" binding:"required"`
	InvoiceDate       *time.Time `json:"invoiceDate" db:"InvoiceDate" binding:"required"`
	BillingAddress    *string    `json:"billingAddress" db:"BillingAddress"`
	BillingCity       *string    `json:"billingCity" db:"BillingCity"`
	BillingState      *string    `json:"billingState" db:"BillingState"`
	BillingCountry    *string    `json:"billingCountry" db:"BillingCountry"`
	BillingPostalCode *string    `json:"billingPostalCode" db:"BillingPostalCode"`
	Total             *float64   `json:"total" db:"Total" binding:"required"`
}

func (i *InvoiceInput) Validate() error {
	if i.CustomerId == nil || i.InvoiceDate == nil || i.Total == nil {
		return errors.New("input structure has no values")
	}
	return nil
}

type InvoiceLine struct {
	InvoiceLineId *int    `json:"invoiceLineId" db:"InvoiceLineId" binding:"required"`
	InvoiceId     *int    `json:"invoiceId" db:"InvoiceId" binding:"required"`
	TrackId       *int    `json:"trackId" db:"TrackId" binding:"required"`
	UnitPrice     *string `json:"unitPrice" db:"UnitPrice" binding:"required"`
	Quantity      *int    `json:"quantity" db:"Quantity" binding:"required"`
}
