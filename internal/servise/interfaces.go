package service

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository"
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
	"Educational-API-DBeaver-Sample-Database/internal/servise/methods"
)

type Album interface {
	Get() (*[]models.Album, error)
	GetById(id int) (*models.Album, error)
	Create(album *models.AlbumInput) (int, error)
	Put(album *models.AlbumInput, idAlbum int) error
	Delete(idAlbum int) error
}

type Artist interface {
	Get() (*[]models.Artist, error)
	GetById(id int) (*models.Artist, error)
	Create(artist *models.ArtistInput) (int, error)
	Put(artist *models.ArtistInput, id int) error
	Delete(id int) error
}

type Customer interface {
	Get() (*[]models.Customer, error)
	GetById(id int) (*models.Customer, error)
	Create(customer *models.CustomerInput) (int, error)
	Put(customer *models.CustomerInput, id int) error
	Delete(id int) error
}

type Employee interface {
	Get() (*[]models.Employee, error)
	GetById(id int) (*models.Employee, error)
	Create(employee *models.EmployeeInput) (int, error)
	Put(employee *models.EmployeeInput, id int) error
	Delete(id int) error
}

type Genre interface {
	Get() (*[]models.Genre, error)
	GetById(id int) (*models.Genre, error)
	Create(genre *models.GenreInput) (int, error)
	Put(artist *models.GenreInput, id int) error
	Delete(id int) error
}

type Invoice interface {
	Get() (*[]models.Invoice, error)
	GetById(id int) (*models.Invoice, error)
	Create(invoice *models.InvoiceInput) (int, error)
	Put(invoice *models.InvoiceInput, id int) error
	Delete(id int) error
}

type MediaType interface {
	Get() (*[]models.MediaType, error)
	GetById(id int) (*models.MediaType, error)
	Create(mediaType *models.MediaTypeInput) (int, error)
	Put(mediaType *models.MediaTypeInput, id int) error
	Delete(id int) error
}

type Playlist interface {
	Get() (*[]models.Playlist, error)
	GetById(id int) (*models.Playlist, error)
	Create(playlist *models.PlaylistInput) (int, error)
	Put(playlist *models.PlaylistInput, id int) error
	Delete(id int) error
}

type Track interface {
	Get() (*[]models.Track, error)
	GetById(id int) (*models.Track, error)
	GetAlbumById(id int) (*models.Track, error)
	Create(track *models.TrackInput) (int, error)
	CreateAlbumById(track *models.TrackInput, id int) (int, error)
	Put(track *models.TrackInput, id int) error
	Delete(id int) error
}

type Service struct {
	Album
	Artist
	Customer
	Employee
	Genre
	Invoice
	MediaType
	Playlist
	Track
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Album:     methods.NewAlbumService(repo.Album, repo.Artist, repo.Track),
		Artist:    methods.NewArtistService(repo.Artist),
		Customer:  methods.NewCustomerService(repo.Customer),
		Employee:  methods.NewEmployeeService(repo.Employee),
		Genre:     methods.NewGenreService(repo.Genre),
		Invoice:   methods.NewInvoiceService(repo.Invoice),
		MediaType: methods.NewMediaTypeService(repo.MediaType),
		Playlist:  methods.NewPlayListService(repo.Playlist),
		Track:     methods.NewTrackService(repo.Track),
	}
}
