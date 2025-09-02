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
type MediaType struct {
	s *service.Service
}

func NewMediaTypeHandler(service *service.Service) *MediaType {
	return &MediaType{s: service}
}

// GET_MediaTypes retrieves all media types
// @Summary Get all media types
// @Description Returns a list of all media types in the system
// @Tags media-types
// @Produce json
// @Success 200 {array} models.MediaType "List of media types"
// @Failure 500 {object} messages.Message "Internal server error"
// @Router /mediaTypes [get]
func (h *MediaType) GET_MediaTypes(c *gin.Context) {

	result, err := h.s.MediaType.Get()
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	} else if result == nil {
		messages.New(c, http.StatusNotFound, "media types not found", messages.Error)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET_MediaTypeById retrieves media type by ID
// @Summary Get media type by ID
// @Description Returns media type by specified identifier
// @Tags media-types
// @Produce json
// @Param id path string true "Media Type ID"
// @Success 200 {object} models.MediaType "Found media type"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Media type not found"
// @Router /mediaTypes/{id} [get]
func (h *MediaType) GET_MediaTypeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	result, err := h.s.MediaType.GetById(id)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	} else if result == nil {
		messages.New(c, http.StatusNotFound, "media type not found", messages.Error)
		return
	}

	c.JSON(http.StatusOK, result)
}

// POST_MediaType creates a new media type
// @Summary Create new media type
// @Description Creates a new media type with provided data
// @Tags media-types
// @Accept json
// @Produce json
// @Param mediaType body models.MediaTypeInput true "Media type data"
// @Success 201 {object} messages.Message "id"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 500 {object} messages.Message "Creation error"
// @Router /mediaTypes [post]
func (h *MediaType) POST_MediaType(c *gin.Context) {
	var mediaType models.MediaTypeInput
	if err := c.BindJSON(&mediaType); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	id, err := h.s.MediaType.Create(&mediaType)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// PUT_MediaType updates an existing media type
// @Summary Update media type
// @Description Updates media type by ID with provided data
// @Tags media-types
// @Accept json
// @Produce json
// @Param id path string true "Media Type ID"
// @Param mediaType body models.MediaTypeInput true "Updated media type data"
// @Success 200 {object} MediaType "Updated media type"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 404 {object} messages.Message "Media type not found"
// @Failure 500 {object} messages.Message "Update error"
// @Router /mediaTypes/{id} [put]
func (h *MediaType) PUT_MediaType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	var mediaType models.MediaTypeInput
	if err := c.BindJSON(&mediaType); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	if err := h.s.MediaType.Put(&mediaType, id); err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}

// DEL_MediaTypeById deletes media type by ID
// @Summary Delete media type
// @Description Deletes media type by specified identifier
// @Tags media-types
// @Produce json
// @Param id path string true "Media Type ID"
// @Success 204 "Media type successfully deleted"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Media type not found"
// @Failure 500 {object} messages.Message "Deletion error"
// @Router /mediaTypes/{id} [delete]
func (h *MediaType) DEL_MediaTypeById(c *gin.Context) {
	mediaType_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	err = h.s.MediaType.Delete(mediaType_id)
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}
