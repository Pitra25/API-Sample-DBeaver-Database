package http_m

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
	service "Educational-API-DBeaver-Sample-Database/internal/servise"
	messages "Educational-API-DBeaver-Sample-Database/pkg/messages"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// swagger:ignore
type Genre struct {
	s *service.Service
}

func NewGenreHandler(service *service.Service) *Genre {
	return &Genre{s: service}
}

// GET_Genres retrieves all genres
// @Summary Get all genres
// @Description Returns a list of all genres in the system
// @Tags genres
// @Produce json
// @Success 200 {array} models.Genre "List of genres"
// @Failure 500 {object} messages.Message "Internal server error"
// @Router /genres [get]
func (h *Genre) GET_Genres(c *gin.Context) {

	result, err := h.s.Genre.Get()
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	} else if result == nil {
		messages.New(c, http.StatusNotFound, "genre not found", messages.Error)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET_GenreById retrieves genre by ID
// @Summary Get genre by ID
// @Description Returns genre by specified identifier
// @Tags genres
// @Produce json
// @Param id path string true "Genre ID"
// @Success 200 {object} models.Genre "Found genre"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Genre not found"
// @Router /genres/{id} [get]
func (h *Genre) GET_GenreById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	result, err := h.s.Genre.GetById(id)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	} else if result == nil {
		messages.New(c, http.StatusNotFound, "genre not found", messages.Error)
		return
	}

	c.JSON(http.StatusOK, result)
}

// POST_Genre creates a new genre
// @Summary Create new genre
// @Description Creates a new genre with provided data
// @Tags genres
// @Accept json
// @Produce json
// @Param genre body models.GenreInput true "Genre data"
// @Success 201 {object} messages.Message "id"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 500 {object} messages.Message "Creation error"
// @Router /genres [post]
func (h *Genre) POST_Genre(c *gin.Context) {
	var genre models.GenreInput
	if err := c.BindJSON(&genre); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	id, err := h.s.Genre.Create(&genre)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// PUT_Genre updates an existing genre
// @Summary Update genre
// @Description Updates genre by ID with provided data
// @Tags genres
// @Accept json
// @Produce json
// @Param id path string true "Genre ID"
// @Param genre body models.GenreInput true "Updated genre data"
// @Success 200 {object} Genre "Updated genre"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 404 {object} messages.Message "Genre not found"
// @Failure 500 {object} messages.Message "Update error"
// @Router /genres/{id} [put]
func (h *Genre) PUT_Genre(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	var genre models.GenreInput
	if err := c.BindJSON(&genre); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	if err := h.s.Genre.Put(&genre, id); err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}

// DEL_GenreById deletes genre by ID
// @Summary Delete genre
// @Description Deletes genre by specified identifier
// @Tags genres
// @Produce json
// @Param id path string true "Genre ID"
// @Success 204 "Genre successfully deleted"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Genre not found"
// @Failure 500 {object} messages.Message "Deletion error"
// @Router /genres/{id} [delete]
func (h *Genre) DEL_GenreById(c *gin.Context) {
	genre_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	err = h.s.Genre.Delete(genre_id)
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}
