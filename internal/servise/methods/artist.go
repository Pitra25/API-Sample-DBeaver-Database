package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository"
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
)

type ArtistService struct {
	repo repository.Artist
}

func NewArtistService(repo repository.Artist) *ArtistService {
	return &ArtistService{
		repo: repo,
	}
}

func (m *ArtistService) Get() (*[]models.Artist, error) {
	return m.repo.Get()
}

func (m *ArtistService) GetById(id int) (*models.Artist, error) {
	return m.repo.GetById(id)
}

func (m *ArtistService) GetByName(name *string) (*models.Artist, error) {
	return m.repo.GetByName(name)
}

func (m *ArtistService) Create(artist *models.ArtistInput) (int, error) {
	if err := artist.Validate(); err != nil {
		return 0, err
	}

	_, err := m.GetByName(artist.Name)
	if err != nil {
		return 0, err
	}

	return m.repo.Create(artist)
}

func (m *ArtistService) Put(artist *models.ArtistInput, id int) error {
	if err := artist.Validate(); err != nil {
		return err
	}
	return m.repo.Put(artist, id)
}

func (m *ArtistService) Delete(id int) error {
	_, err := m.GetById(id)
	if err != nil {
		return err
	}

	return m.repo.Delete(id)
}
