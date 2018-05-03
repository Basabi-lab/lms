package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/kokoax/music_lab/adapters/presenters"
	"github.com/kokoax/music_lab/usecases"
)

type AlbumHandler struct {
	usecases.AllAlbumsUsecase
	presenters.AllAlbumsPresenter
}

func (h *AlbumHandler) getAllAlbums(c *gin.Context) {
	albums, _ := h.AllAlbums(c)
	json, _ := h.ToJson(*albums)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(json.Json)
}

func (h *AlbumHandler) getWithId(c *gin.Context) {
	// TODO: contextからIDを取り出してGetByIDに渡す
	// info, _ := h.GetById()
	// h.GetWithID(info)
}

func NewAlbumHandler(r *gin.Engine, db *gorm.DB) {
	handler := AlbumHandler{
		usecases.NewAllAlbumsUsecase(db),
		presenters.NewAllAlbumsPresenter(db),
	}
	r.GET("/album", handler.getAllAlbums)
	r.GET("/album/:id", handler.getWithId)
}
