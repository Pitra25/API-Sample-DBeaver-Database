package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository"
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
)

type PlayListService struct {
	repo repository.Playlist
}

func NewPlayListService(repo repository.Playlist) *PlayListService {
	return &PlayListService{
		repo: repo,
	}
}

func (m *PlayListService) Get() (*[]models.Playlist, error) {
	return m.repo.Get()
}

func (m *PlayListService) GetById(id int) (*models.Playlist, error) {
	return m.repo.GetById(id)
}

func (m *PlayListService) Create(playlist *models.PlaylistInput) (int, error) {
	if err := playlist.Validate(); err != nil {
		return 0, err
	}
	return m.repo.Create(playlist)
}

func (m *PlayListService) Put(playlist *models.PlaylistInput, id int) error {
	if err := playlist.Validate(); err != nil {
		return err
	}
	return m.repo.Put(playlist, id)
}

func (m *PlayListService) Delete(id int) error {
	_, err := m.GetById(id)
	if err != nil {
		return err
	}

	return m.repo.Delete(id)
}
