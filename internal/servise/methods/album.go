package methods

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository"
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"

	"github.com/sirupsen/logrus"
)

type AlbumService struct {
	repo_Al repository.Album
	repo_Ar repository.Artist
	repo_Tr repository.Track
}

func NewAlbumService(
	repo repository.Album,
	repo_Ar repository.Artist,
	repo_Tr repository.Track,
) *AlbumService {
	return &AlbumService{
		repo_Al: repo,
		repo_Ar: repo_Ar,
		repo_Tr: repo_Tr,
	}
}

func (m *AlbumService) Get() (*[]models.Album, error) {
	logrus.Debug("service")
	return m.repo_Al.Get()
}

func (m *AlbumService) GetById(id int) (*models.Album, error) {
	return m.repo_Al.GetById(id)
}

func (m *AlbumService) GetByName(title *string) (*models.Album, error) {
	return m.repo_Al.GetByName(title)
}

func (m *AlbumService) Create(album *models.AlbumInput) (int, error) {
	if err := album.Validate(); err != nil {
		return 0, err
	}

	_, err := m.GetByName(album.Title)
	if err != nil {
		return 0, err
	}

	return m.repo_Al.Create(album)
}

func (m *AlbumService) Put(album *models.AlbumInput, idAlbum int) error {
	if err := album.Validate(); err != nil {
		return err
	}

	_, err := m.repo_Al.GetById(idAlbum)
	if err != nil {
		return err
	}

	_, err = m.repo_Ar.GetById(*album.ArtistId)
	if err != nil {
		return err
	}

	return m.repo_Al.Put(album, idAlbum)
}

func (m *AlbumService) Delete(idAlbum int) error {
	_, err := m.GetById(idAlbum)
	if err != nil {
		return err
	}

	album, err := m.repo_Al.GetById(idAlbum)
	if err != nil {
		return err
	}

	// _, err := m.repo_Tr.GetById(album.)

	return m.repo_Al.Delete(idAlbum, *album.ArtistId)
}
