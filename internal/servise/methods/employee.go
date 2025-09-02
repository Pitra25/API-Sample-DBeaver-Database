package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository"
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
)

type EmployeeService struct {
	repo repository.Employee
}

func NewEmployeeService(repo repository.Employee) *EmployeeService {
	return &EmployeeService{
		repo: repo,
	}
}

func (m *EmployeeService) Get() (*[]models.Employee, error) {
	return m.repo.Get()
}

func (m *EmployeeService) GetById(id int) (*models.Employee, error) {
	return m.repo.GetById(id)
}

func (m *EmployeeService) Create(employee *models.EmployeeInput) (int, error) {
	if err := employee.Validate(); err != nil {
		return 0, err
	}

	return m.repo.Create(employee)
}

func (m *EmployeeService) Put(employee *models.EmployeeInput, id int) error {
	if err := employee.Validate(); err != nil {
		return err
	}
	return m.repo.Put(employee, id)
}

func (m *EmployeeService) Delete(id int) error {
	_, err := m.GetById(id)
	if err != nil {
		return err
	}

	return m.repo.Delete(id)
}
