package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository"
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
)

type InvoiceService struct {
	repo repository.Invoice
}

func NewInvoiceService(repo repository.Invoice) *InvoiceService {
	return &InvoiceService{
		repo: repo,
	}
}

func (m *InvoiceService) Get() (*[]models.Invoice, error) {
	return m.repo.Get()
}

func (m *InvoiceService) GetById(id int) (*models.Invoice, error) {
	return m.repo.GetById(id)
}

func (m *InvoiceService) Create(invoice *models.InvoiceInput) (int, error) {
	if err := invoice.Validate(); err != nil {
		return 0, err
	}

	return m.repo.Create(invoice)
}

func (m *InvoiceService) Put(invoice *models.InvoiceInput, id int) error {
	if err := invoice.Validate(); err != nil {
		return err
	}
	return m.repo.Put(invoice, id)
}

func (m *InvoiceService) Delete(id int) error {
	_, err := m.GetById(id)
	if err != nil {
		return err
	}

	return m.repo.Delete(id)
}
