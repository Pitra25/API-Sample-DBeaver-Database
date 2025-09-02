package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository"
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
)

type MediaTypeService struct {
	repo repository.MediaType
}

func NewMediaTypeService(repo repository.MediaType) *MediaTypeService {
	return &MediaTypeService{
		repo: repo,
	}
}

func (m *MediaTypeService) Get() (*[]models.MediaType, error) {
	return m.repo.Get()
}

func (m *MediaTypeService) GetById(id int) (*models.MediaType, error) {
	return m.repo.GetById(id)
}

func (m *MediaTypeService) Create(mediaType *models.MediaTypeInput) (int, error) {
	if err := mediaType.Validate(); err != nil {
		return 0, err
	}
	return m.repo.Create(mediaType)
}

func (m *MediaTypeService) Put(mediaType *models.MediaTypeInput, id int) error {
	if err := mediaType.Validate(); err != nil {
		return err
	}
	return m.repo.Put(mediaType, id)
}

func (m *MediaTypeService) Delete(id int) error {
	_, err := m.GetById(id)
	if err != nil {
		return err
	}

	return m.repo.Delete(id)
}
