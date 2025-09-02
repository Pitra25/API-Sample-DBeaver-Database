package http_m

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
	service "Educational-API-DBeaver-Sample-Database/internal/servise"
	messages "Educational-API-DBeaver-Sample-Database/pkg/messages"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// swagger:ignore
type Album struct {
	s *service.Service
}

func NewAlbumHandler(service *service.Service) *Album {
	return &Album{s: service}
}

// @Summary Get all albums
// @Description Returns a list of all albums in the system
// @Tags albums
// @Produce json
// @Success 200 {array} models.Album
// @Failure 500 {object} messages.Message "Internal server error"
// @Router /albums [get]
func (h *Album) GET_Albums(c *gin.Context) {
	logrus.Debug("handler")
	result, err := h.s.Album.Get()
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Get album by ID
// @Description Returns album by specified identifier
// @Tags albums
// @Produce json
// @Param id path string true "Album ID"
// @Success 200 {object} models.Album
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Album not found"
// @Router /albums/{id} [get]
func (h *Album) GET_AlbumById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	result, err := h.s.Album.GetById(id)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Create new album
// @Description Creates a new album with provided data
// @Tags albums
// @Accept json
// @Produce json
// @Param album body models.AlbumInput true "Album data"
// @Success 201 {object} messages.Message "id"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 500 {object} messages.Message "Creation error"
// @Router /albums [post]
func (h *Album) POST_Album(c *gin.Context) {
	var album models.AlbumInput
	if err := c.BindJSON(&album); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	id, err := h.s.Album.Create(&album)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Update album
// @Description Updates album by ID with provided data
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Param album body models.AlbumInput true "Updated album data"
// @Success 200 {object} Album "Updated album"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 404 {object} messages.Message "Album not found"
// @Failure 500 {object} messages.Message "Update error"
// @Router /albums/{id} [put]
func (h *Album) PUT_Album(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	var album models.AlbumInput
	if err := c.BindJSON(&album); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	if err := h.s.Album.Put(&album, id); err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}

// @Summary Delete album
// @Description Deletes album by specified identifier
// @Tags albums
// @Produce json
// @Param id path string true "Album ID"
// @Success 204 "Album successfully deleted"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Album not found"
// @Failure 500 {object} messages.Message "Deletion error"
// @Router /albums/{id} [delete]
func (h *Album) DEL_AlbumById(c *gin.Context) {
	idAlbum, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	err = h.s.Album.Delete(idAlbum)
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}

// @Summary Get track album by ID
// @Description Returns track by specified identifier
// @Tags albums
// @Produce json
// @Param id path string true "Track ID"
// @Success 200 {object} models.Track "Found track"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Track not found"
// @Router /albums/{id}/track [get]
func (h *Track) GET_TrackAlbumById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	result, err := h.s.Track.GetAlbumById(id)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Create new track album by id
// @Description Creates a new track with provided data
// @Tags albums
// @Accept json
// @Produce json
// @Param track body models.TrackInput true "Track data"
// @Success 201 {object} Track "Created track"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 500 {object} messages.Message "Creation error"
// @Router /albums/{id}/track [post]
func (h *Track) POST_TrackAlbumById(c *gin.Context) {
	id_album, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	_, err = h.s.Album.GetById(id_album)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	var track models.TrackInput
	if err := c.BindJSON(&track); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	id, err := h.s.Track.CreateAlbumById(&track, id_album)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
