package http

import (
	http_m "Educational-API-DBeaver-Sample-Database/internal/handler/http/methods"
	service "Educational-API-DBeaver-Sample-Database/internal/servise"

	"github.com/gin-gonic/gin"
)

type Album interface {
	GET_Albums(c *gin.Context)
	GET_AlbumById(c *gin.Context)
	POST_Album(c *gin.Context)
	PUT_Album(c *gin.Context)
	DEL_AlbumById(c *gin.Context)
}

type Artist interface {
	GET_Artists(c *gin.Context)
	GET_ArtistsById(c *gin.Context)
	GET_ArtistById(c *gin.Context)
	POST_Artist(c *gin.Context)
	PUT_Artist(c *gin.Context)
	DEL_ArtistById(c *gin.Context)
}

type Customer interface {
	GET_Customers(c *gin.Context)
	GET_CustomerById(c *gin.Context)
	POST_Customer(c *gin.Context)
	PUT_Customer(c *gin.Context)
	DEL_CustomerById(c *gin.Context)
}

type Employee interface {
	GET_Employees(c *gin.Context)
	GET_EmployeeById(c *gin.Context)
	POST_Employee(c *gin.Context)
	PUT_Employee(c *gin.Context)
	DEL_EmployeeById(c *gin.Context)
}

type Genre interface {
	GET_Genres(c *gin.Context)
	GET_GenreById(c *gin.Context)
	POST_Genre(c *gin.Context)
	PUT_Genre(c *gin.Context)
	DEL_GenreById(c *gin.Context)
}

type Invoice interface {
	GET_Invoices(c *gin.Context)
	GET_InvoiceById(c *gin.Context)
	POST_Invoice(c *gin.Context)
	PUT_Invoice(c *gin.Context)
	DEL_InvoiceById(c *gin.Context)
}

type MediaType interface {
	GET_MediaTypes(c *gin.Context)
	GET_MediaTypeById(c *gin.Context)
	POST_MediaType(c *gin.Context)
	PUT_MediaType(c *gin.Context)
	DEL_MediaTypeById(c *gin.Context)
}

type Playlist interface {
	GET_PlayLists(c *gin.Context)
	GET_PlayListById(c *gin.Context)
	POST_PlayList(c *gin.Context)
	PUT_PlayList(c *gin.Context)
	DEL_PlayListById(c *gin.Context)
}

type Track interface {
	GET_Tracks(c *gin.Context)
	GET_TrackById(c *gin.Context)
	POST_Track(c *gin.Context)
	PUT_Track(c *gin.Context)
	DEL_TrackById(c *gin.Context)

	GET_TrackAlbumById(c *gin.Context)
	POST_TrackAlbumById(c *gin.Context)
}

type HandlerStr struct {
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

func NewHandler(s *service.Service) *HandlerStr {
	return &HandlerStr{
		Album:     http_m.NewAlbumHandler(s),
		Artist:    http_m.NewArtistHandler(s),
		Customer:  http_m.NewCustomerHandler(s),
		Employee:  http_m.NewEmployeeHandler(s),
		Genre:     http_m.NewGenreHandler(s),
		Invoice:   http_m.NewInvoiceHandler(s),
		MediaType: http_m.NewMediaTypeHandler(s),
		Playlist:  http_m.NewPlayListHandler(s),
		Track:     http_m.NewTrackHandler(s),
	}
}
