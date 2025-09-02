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
type Artist struct {
	s *service.Service
}

func NewArtistHandler(service *service.Service) *Artist {
	return &Artist{s: service}
}

// GET_Artists retrieves all artists
// @Summary Get all artists
// @Description Returns a list of all artists in the system
// @Tags artists
// @Produce json
// @Success 200 {array} models.Artist
// @Failure 500 {object} messages.Message "Internal server error"
// @Router /artists [get]
func (h *Artist) GET_Artists(c *gin.Context) {
	result, err := h.s.Artist.Get()
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET_ArtistsById retrieves artists by ID (plural version)
// @Summary Get artists by ID
// @Description Returns artists by specified ID
// @Tags artists
// @Produce json
// @Param id path string true "Artist ID"
// @Success 200 {object} models.Artist
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Artist not found"
// @Router /artists/{id} [get]
func (h *Artist) GET_ArtistsById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
	}

	result, err := h.s.Artist.GetById(id)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
	}

	if result != nil {
		logrus.Debug("Artist found")
		c.JSON(http.StatusOK, result)
	} else if result == nil && err == nil {
		logrus.Debug("Artist not found")
		c.JSON(http.StatusNotFound, "Artist not found")
	} else {
		logrus.Debug("Internal server error")
		c.JSON(http.StatusInternalServerError, "Internal server error")
	}
}

// GET_ArtistById retrieves artist by ID (singular version)
// @Summary Get artist by ID
// @Description Returns artist by specified identifier
// @Tags artists
// @Produce json
// @Param id path string true "Artist ID"
// @Success 200 {object} models.Artist
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Artist not found"
// @Router /artist/{id} [get]
func (h *Artist) GET_ArtistById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	result, err := h.s.Artist.GetById(id)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, result)
}

// POST_Artist creates a new artist
// @Summary Create new artist
// @Description Creates a new artist with provided data
// @Tags artists
// @Accept json
// @Produce json
// @Param artist body Artist true "Artist data"
// @Success 201 {object} messages.Message "id"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 500 {object} messages.Message "Creation error"
// @Router /artists [post]
func (h *Artist) POST_Artist(c *gin.Context) {
	var artist models.ArtistInput
	if err := c.BindJSON(&artist); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	id, err := h.s.Artist.Create(&artist)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// PUT_Artist updates an existing artist
// @Summary Update artist
// @Description Updates artist by ID with provided data
// @Tags artists
// @Accept json
// @Produce json
// @Param id path string true "Artist ID"
// @Param artist body models.ArtistInput true "Updated artist data"
// @Success 200 {object} Artist "Updated artist"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 404 {object} messages.Message "Artist not found"
// @Failure 500 {object} messages.Message "Update error"
// @Router /artists/{id} [put]
func (h *Artist) PUT_Artist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	var artist models.ArtistInput
	if err := c.BindJSON(&artist); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	if err := h.s.Artist.Put(&artist, id); err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}

// DEL_ArtistById deletes artist by ID
// @Summary Delete artist
// @Description Deletes artist by specified identifier
// @Tags artists
// @Produce json
// @Param id path string true "Artist ID"
// @Success 204 "Artist successfully deleted"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Artist not found"
// @Failure 500 {object} messages.Message "Deletion error"
// @Router /artists/{id} [delete]
func (h *Artist) DEL_ArtistById(c *gin.Context) {
	artist_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	err = h.s.Artist.Delete(artist_id)
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}
