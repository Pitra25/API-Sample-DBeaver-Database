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
type PlayList struct {
	s *service.Service
}

func NewPlayListHandler(service *service.Service) *PlayList {
	return &PlayList{s: service}
}

// GET_PlayLists retrieves all playlists
// @Summary Get all playlists
// @Description Returns a list of all playlists in the system
// @Tags playlists
// @Produce json
// @Success 200 {array} models.Playlist "List of playlists"
// @Failure 500 {object} messages.Message "Internal server error"
// @Router /playlists [get]
func (h *PlayList) GET_PlayLists(c *gin.Context) {
	result, err := h.s.Playlist.Get()
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET_PlayListById retrieves playlist by ID
// @Summary Get playlist by ID
// @Description Returns playlist by specified identifier
// @Tags playlists
// @Produce json
// @Param id path string true "Playlist ID"
// @Success 200 {object} models.Playlist "Found playlist"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Playlist not found"
// @Router /playlists/{id} [get]
func (h *PlayList) GET_PlayListById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	result, err := h.s.Playlist.GetById(id)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, result)
}

// POST_PlayList creates a new playlist
// @Summary Create new playlist
// @Description Creates a new playlist with provided data
// @Tags playlists
// @Accept json
// @Produce json
// @Param playlist body models.PlaylistInput true "Playlist data"
// @Success 201 {object} messages.Message "id"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 500 {object} messages.Message "Creation error"
// @Router /playlists [post]
func (h *PlayList) POST_PlayList(c *gin.Context) {
	var playlist models.PlaylistInput
	if err := c.BindJSON(&playlist); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	id, err := h.s.Playlist.Create(&playlist)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// PUT_PlayList updates an existing playlist
// @Summary Update playlist
// @Description Updates playlist by ID with provided data
// @Tags playlists
// @Accept json
// @Produce json
// @Param id path string true "Playlist ID"
// @Param playlist body models.PlaylistInput true "Updated playlist data"
// @Success 204 "Updated playlist"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 404 {object} messages.Message "Playlist not found"
// @Failure 500 {object} messages.Message "Update error"
// @Router /playlists/{id} [put]
func (h *PlayList) PUT_PlayList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	var playList models.PlaylistInput
	if err := c.BindJSON(&playList); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	if err := h.s.Playlist.Put(&playList, id); err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.Status(http.StatusOK)
}

// DEL_PlayListById deletes playlist by ID
// @Summary Delete playlist
// @Description Deletes playlist by specified identifier
// @Tags playlists
// @Produce json
// @Param id path string true "Playlist ID"
// @Success 204 "Playlist type successfully deleted"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Playlist not found"
// @Failure 500 {object} messages.Message "Deletion error"
// @Router /playlists/{id} [delete]
func (h *PlayList) DEL_PlayListById(c *gin.Context) {
	playList_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	err = h.s.Playlist.Delete(playList_id)
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	c.Status(http.StatusOK)
}
