package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository"
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
)

type CustomerService struct {
	repo repository.Customer
}

func NewCustomerService(repo repository.Customer) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}

func (m *CustomerService) Get() (*[]models.Customer, error) {
	return m.repo.Get()
}

func (m *CustomerService) GetById(id int) (*models.Customer, error) {
	return m.repo.GetById(id)
}

func (m *CustomerService) Create(customer *models.CustomerInput) (int, error) {
	if err := customer.Validate(); err != nil {
		return 0, err
	}

	return m.repo.Create(customer)
}

func (m *CustomerService) Put(customer *models.CustomerInput, id int) error {
	if err := customer.Validate(); err != nil {
		return err
	}
	return m.repo.Put(customer, id)
}

func (m *CustomerService) Delete(id int) error {
	_, err := m.GetById(id)
	if err != nil {
		return err
	}

	return m.repo.Delete(id)
}
