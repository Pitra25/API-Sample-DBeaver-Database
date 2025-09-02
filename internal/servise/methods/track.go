package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository"
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
)

type TrackService struct {
	repo repository.Track
}

func NewTrackService(repo repository.Track) *TrackService {
	return &TrackService{
		repo: repo,
	}
}

func (m *TrackService) Get() (*[]models.Track, error) {
	return m.repo.Get()
}

func (m *TrackService) GetById(id int) (*models.Track, error) {
	return m.repo.GetById(id)
}

func (m *TrackService) Create(track *models.TrackInput) (int, error) {
	if err := track.Validate(); err != nil {
		return 0, err
	}
	return m.repo.Create(track)
}

func (m *TrackService) Put(track *models.TrackInput, id int) error {
	if err := track.Validate(); err != nil {
		return err
	}
	return m.repo.Put(track, id)
}

func (m *TrackService) Delete(id int) error {
	_, err := m.GetById(id)
	if err != nil {
		return err
	}

	return m.repo.Delete(id)
}

// === 

func (m *TrackService) GetAlbumById(id int) (*models.Track, error) {
	return nil, nil
}

func (m *TrackService) CreateAlbumById(track *models.TrackInput, id int) (int, error) {
	return 0, nil
}
