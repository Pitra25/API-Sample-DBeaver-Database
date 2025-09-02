package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository"
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
)

type GenreService struct {
	repo repository.Genre
}

func NewGenreService(repo repository.Genre) *GenreService {
	return &GenreService{
		repo: repo,
	}
}

func (m *GenreService) Get() (*[]models.Genre, error) {
	return m.repo.Get()
}

func (m *GenreService) GetById(id int) (*models.Genre, error) {
	return m.repo.GetById(id)
}

func (m *GenreService) Create(genre *models.GenreInput) (int, error) {
	if err := genre.Validate(); err != nil {
		return 0, err
	}

	return m.repo.Create(genre)
}

func (m *GenreService) Put(genre *models.GenreInput, id int) error {
	if err := genre.Validate(); err != nil {
		return err
	}
	return m.repo.Put(genre, id)
}

func (m *GenreService) Delete(id int) error {
	_, err := m.GetById(id)
	if err != nil {
		return err
	}

	return m.repo.Delete(id)
}
