package api

import (
	"go/scr/hhruxongs/storage"

	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
)

func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a *API) configureRouterField() {
	a.router.HandleFunc(prefix+"/song", a.PostSong).Methods("POST")
	a.router.HandleFunc(prefix+"/library/{page}/{pageSize}", a.GetSongs).Methods("GET")
	a.router.HandleFunc(prefix+"/song/{id}/{page}/{pageSize}", a.GetSongsIdPart).Methods("GET")
	a.router.HandleFunc(prefix+"/song/{id}", a.PutSongs).Methods("PUT")
	a.router.HandleFunc(prefix+"/song/{id}", a.DeleteSongsById).Methods("DELET")
}

// swagger:route POST /api/v1/song song PostSong
// Post song to library
//
// @Summary Create new song in library
// @Description Create new song with given data
// @Tags songs
// @Accept  json
// @Produce  json
// @Param body body SongRequest true "Song details"
// @Success 201 {object} SongResponse
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /song [post]

// swagger:route GET /api/v1/library/{page}/{pageSize} library GetSongs
// Get songs from library
//
// @Summary Get songs from library
// @Description Get songs from library with pagination
// @Tags library
// @Produce  json
// @Param page path int false "Page number"
// @Param pageSize path int false "Page size"
// @Success 200 {object} SongsResponse
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /library/{page}/{pageSize} [get]

// swagger:route GET /api/v1/song/{id}/{page}/{pageSize} library GetSongsIdPart
// Get songs from library by id with pagination
//
// @Summary Get songs from library by id with pagination
// @Description Get songs from library by id with pagination
// @Tags library
// @Produce  json
// @Param id path int true "Song ID"
// @Param page path int false "Page number"
// @Param pageSize path int false "Page size"
// @Success 200 {object} SongsResponse
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Router /song/{id}/{page}/{pageSize} [get]

// swagger:route PUT /api/v1/song/{id} song PutSongs
// Update song by id
//
// @Summary Update song by id
// @Description Update song details by given id
// @Tags songs
// @Accept  json
// @Produce  json
// @Param id path int true "Song ID"
// @Param body body SongRequest true "Updated song details"
// @Success 200 {object} SongResponse
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Router /song/{id} [put]

// swagger:route DELETE /api/v1/song/{id} song DeleteSongsById
// Delete song by id
//
// @Summary Delete song by id
// @Description Delete song with given id from library
// @Tags songs
// @Param id path int true "Song ID"
// @Success 204 {object} DeleteResponse
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Router /song/{id} [delete]

func (a *API) configuresStorageFild() error {
	storage := storage.New(a.config.Storage)
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
