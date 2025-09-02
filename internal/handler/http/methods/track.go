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
type Track struct {
	s *service.Service
}

func NewTrackHandler(service *service.Service) *Track {
	return &Track{s: service}
}

// GET_Tracks retrieves all tracks
// @Summary Get all tracks
// @Description Returns a list of all tracks in the system
// @Tags tracks
// @Produce json
// @Success 200 {array} models.Track "List of tracks"
// @Failure 500 {object} messages.Message "Internal server error"
// @Router /tracks [get]
func (h *Track) GET_Tracks(c *gin.Context) {

	result, err := h.s.Track.Get()
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	} else if result == nil {
		messages.New(c, http.StatusNotFound, "tracks not found", messages.Error)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Get track by ID
// @Description Returns track by specified identifier
// @Tags tracks
// @Produce json
// @Param id path string true "Track ID"
// @Success 200 {object} models.Track "Found track"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Track not found"
// @Router /tracks/{id} [get]
func (h *Track) GET_TrackById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	result, err := h.s.Track.GetById(id)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	} else if result == nil {
		messages.New(c, http.StatusNotFound, "track not found", messages.Error)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Create new track
// @Description Creates a new track with provided data
// @Tags tracks
// @Accept json
// @Produce json
// @Param track body models.TrackInput true "Track data"
// @Success 201 {object} messages.Message "id"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 500 {object} messages.Message "Creation error"
// @Router /tracks [post]
func (h *Track) POST_Track(c *gin.Context) {
	var track models.TrackInput
	if err := c.BindJSON(&track); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	id, err := h.s.Track.Create(&track)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Update track
// @Description Updates track by ID with provided data
// @Tags tracks
// @Accept json
// @Produce json
// @Param id path string true "Track ID"
// @Param track body models.TrackInput true "Updated track data"
// @Success 200 {object} Track "Updated track"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 404 {object} messages.Message "Track not found"
// @Failure 500 {object} messages.Message "Update error"
// @Router /tracks/{id} [put]
func (h *Track) PUT_Track(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	var track models.TrackInput
	if err := c.BindJSON(&track); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	if err := h.s.Track.Put(&track, id); err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}

// @Summary Delete track
// @Description Deletes track by specified identifier
// @Tags tracks
// @Produce json
// @Param id path string true "Track ID"
// @Success 204 "Track successfully deleted"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Track not found"
// @Failure 500 {object} messages.Message "Deletion error"
// @Router /tracks/{id} [delete]
func (h *Track) DEL_TrackById(c *gin.Context) {
	track_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	err = h.s.Track.Delete(track_id)
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}

// Methods by Album ID
