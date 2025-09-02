package models

import "errors"

type Track struct {
	TrackId      *int     `json:"trackId" db:"TrackId"`
	Name         *string  `json:"name" db:"Name"`
	AlbumId      *int     `json:"albumId" db:"AlbumId"`
	MediaTypeId  *int     `json:"mediaTypeId" db:"MediaTypeId"`
	GenreId      *int     `json:"genreId" db:"GenreId"`
	Composer     *string  `json:"composer" db:"Composer"`
	Milliseconds *int     `json:"milliseconds" db:"Milliseconds"`
	Bytes        *int     `json:"bytes" db:"Bytes"`
	UnitPrice    *float64 `json:"unitPrice" db:"UnitPrice"`
}

type TrackInput struct {
	Name         *string  `json:"name" db:"Name" binding:"required"`
	AlbumId      *int     `json:"albumId" db:"AlbumId"`
	MediaTypeId  *int     `json:"mediaTypeId" db:"MediaTypeId" binding:"required"`
	GenreId      *int     `json:"genreId" db:"GenreId"`
	Composer     *string  `json:"composer" db:"Composer"`
	Milliseconds *int     `json:"milliseconds" db:"Milliseconds" binding:"required"`
	Bytes        *int     `json:"bytes" db:"Bytes"`
	UnitPrice    *float64 `json:"unitPrice" db:"UnitPrice" binding:"required"`
}

func (t *TrackInput) Validate() error {
	if t.Name == nil || t.MediaTypeId == nil || t.Milliseconds == nil {
		return errors.New("input structure has no values")
	}
	return nil
}

// =============================================

type Playlist struct {
	PlaylistId *int    `json:"playlistId" db:"PlaylistId"`
	Name       *string `json:"name" db:"Name"`
}

type PlaylistInput struct {
	Name *string `json:"name" db:"Name" binding:"required"`
}

func (t *PlaylistInput) Validate() error {
	if t.Name == nil {
		return errors.New("input structure has no values")
	}
	return nil
}

// =============================================

type PlaylistTrack struct {
	PlaylistId *int `json:"playlistId" db:"PlaylistId" binding:"required"`
	TrackId    *int `json:"trackId" db:"TrackId" binding:"required"`
}

// =============================================

type MediaType struct {
	MediaTypeId *int    `json:"mediaTypeId" db:"MediaTypeId"`
	Name        *string `json:"name" db:"Name"`
}

type MediaTypeInput struct {
	Name *string `json:"name" db:"Name"`
}

func (t *MediaTypeInput) Validate() error {
	if t.Name == nil {
		return errors.New("input structure has no values")
	}
	return nil
}

// =============================================

type Genre struct {
	GenreId *int    `json:"genreId" db:"GenreId"`
	Name    *string `json:"name" db:"Name"`
}

type GenreInput struct {
	Name *string `json:"name" db:"Name" binding:"required"`
}

func (t *GenreInput) Validate() error {
	if t.Name == nil {
		return errors.New("input structure has no values")
	}
	return nil
}

// =============================================

type Album struct {
	AlbumId  *int    `json:"albumId" db:"AlbumId"`
	Title    *string `json:"title" db:"Title"`
	ArtistId *int    `json:"artistId" db:"ArtistId"`
}

type AlbumInput struct {
	Title    *string `json:"title" db:"Title" binding:"required"`
	ArtistId *int    `json:"artistId" db:"ArtistId" binding:"required"`
}

func (t *AlbumInput) Validate() error {
	if t.Title == nil {
		return errors.New("input structure has no values")
	}
	return nil
}

// =============================================

type Artist struct {
	ArtistId *int    `json:"artistId" db:"ArtistId"`
	Name     *string `json:"name" db:"Name"`
}

type ArtistInput struct {
	Name *string `json:"name" db:"Name"`
}

func (t *ArtistInput) Validate() error {
	if t.Name == nil {
		return errors.New("input structure has no values")
	}
	return nil
}
