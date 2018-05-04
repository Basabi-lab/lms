package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/Basabi-lab/lms/adapters/presenters"
	"github.com/Basabi-lab/lms/usecases"
)

type AlbumHandler struct {
	usecases.AllAlbumUsecase
	presenters.AllAlbumPresenter
}

func (h *AlbumHandler) getAllAlbum(c *gin.Context) {
	albums, _ := h.AllAlbum(c)
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
		usecases.NewAllAlbumUsecase(db),
		presenters.NewAllAlbumPresenter(db),
	}
	r.GET("/album", handler.getAllAlbum)
	r.GET("/album/:id", handler.getWithId)
}
