package repository

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository/methods"
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"

	"database/sql"
)

type Album interface {
	Get() (*[]models.Album, error)
	GetById(id int) (*models.Album, error)
	GetByName(title *string) (*models.Album, error)
	Create(album *models.AlbumInput) (int, error)
	Put(album *models.AlbumInput, iddAlbumOld int) error
	Delete(idAlbum, idArtist int) error
}

type Artist interface {
	Get() (*[]models.Artist, error)
	GetById(id int) (*models.Artist, error)
	GetByName(title *string) (*models.Artist, error)
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
	Create(track *models.TrackInput) (int, error)
	Put(track *models.TrackInput, id int) error
	Delete(id int) error
}

type Repository struct {
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

func New(db *sql.DB) *Repository {
	return &Repository{
		Album:     methods.NewAlbumDB(db),
		Artist:    methods.NewArtistDB(db),
		Customer:  methods.NewCustomerDB(db),
		Employee:  methods.NewEmployeeDB(db),
		Genre:     methods.NewGenreDB(db),
		Invoice:   methods.NewInvoiceDB(db),
		MediaType: methods.NewMediaTypeDB(db),
		Playlist:  methods.NewPlayListDB(db),
		Track:     methods.NewTrackDB(db),
	}
}
