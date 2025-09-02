package http

import (
	service "Educational-API-DBeaver-Sample-Database/internal/servise"

	_ "Educational-API-DBeaver-Sample-Database/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	s *service.Service
	m *HandlerStr
}

func New(service *service.Service, m *HandlerStr) *Handler {
	return &Handler{
		s: service,
		m: m,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		album := api.Group("albums")
		{
			album.GET("/", h.m.GET_Albums)
			album.GET("/:id", h.m.GET_AlbumById)
			album.POST("/", h.m.POST_Album)
			album.PUT("/:id", h.m.PUT_Album)
			album.DELETE("/:id", h.m.DEL_AlbumById)

			track := album.Group(":id/track")
			{
				track.GET("/", h.m.GET_TrackAlbumById)
				track.POST("/", h.m.POST_TrackAlbumById)
			}
		}

		artist := api.Group("artists")
		{
			artist.GET("/", h.m.GET_Artists)
			artist.GET("/:id", h.m.GET_ArtistById)
			artist.POST("/", h.m.POST_Artist)
			artist.PUT("/:id", h.m.PUT_Artist)
			artist.DELETE("/:id", h.m.DEL_ArtistById)
		}

		customer := api.Group("customers")
		{
			customer.GET("/", h.m.GET_Customers)
			customer.GET("/:id", h.m.GET_CustomerById)
			customer.POST("/", h.m.POST_Customer)
			customer.PUT("/", h.m.PUT_Customer)
			customer.DELETE("/:id", h.m.DEL_CustomerById)
		}

		employee := api.Group("employees")
		{
			employee.GET("/", h.m.GET_Employees)
			employee.GET("/:id", h.m.GET_EmployeeById)
			employee.POST("/", h.m.POST_Employee)
			employee.PUT("/:id", h.m.PUT_Employee)
			employee.DELETE("/:id", h.m.DEL_EmployeeById)
		}

		genre := api.Group("genres")
		{
			genre.GET("/", h.m.GET_Genres)
			genre.GET("/:id", h.m.GET_GenreById)
			genre.POST("/", h.m.POST_Genre)
			genre.PUT("/id", h.m.PUT_Genre)
			genre.DELETE("/id", h.m.DEL_GenreById)
		}

		invoice := api.Group("invoices")
		{
			invoice.GET("/", h.m.GET_Invoices)
			invoice.GET("/:id", h.m.GET_InvoiceById)
			invoice.POST("/", h.m.POST_Invoice)
			invoice.PUT("/:id", h.m.PUT_Invoice)
			invoice.DELETE("/:id", h.m.DEL_InvoiceById)
		}

		mediaType := api.Group("mediaType")
		{
			mediaType.GET("/", h.m.GET_MediaTypes)
			mediaType.GET("/:id", h.m.GET_MediaTypeById)
			mediaType.POST("/", h.m.POST_MediaType)
			mediaType.PUT("/id", h.m.PUT_MediaType)
			mediaType.DELETE("/id", h.m.DEL_MediaTypeById)
		}

		playList := api.Group("playlists")
		{
			playList.GET("/", h.m.GET_PlayLists)
			playList.GET("/:id", h.m.GET_PlayListById)
			playList.POST("/", h.m.POST_PlayList)
			playList.PUT("/:id", h.m.PUT_PlayList)
			playList.DELETE("/:id", h.m.DEL_PlayListById)

			playlistTrack := playList.Group(":id/track")
			{
				playlistTrack.GET("/", h.m.GET_PlayListById)
				playlistTrack.POST("/", h.m.POST_PlayList)
			}
		}

		track := api.Group("tracks")
		{
			track.GET("/", h.m.GET_Tracks)
			track.GET("/:id", h.m.GET_TrackById)
			track.POST("/", h.m.POST_Track)
			track.PUT("/id", h.m.PUT_Track)
			track.DELETE("/id", h.m.DEL_TrackById)

			album := track.Group(":id/album")
			{
				album.GET("/:id", h.m.GET_AlbumById)

				artist := album.Group(":id/artist")
				{
					artist.GET("/:id", h.m.GET_ArtistsById)
				}
			}

			playList := album.Group(":id/playlist")
			{
				playList.GET("/:id", h.m.GET_PlayListById)
			}
		}
	}

	return router
}
