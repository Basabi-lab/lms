package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/kokoax/music_lab/adapters/presenters"
	"github.com/kokoax/music_lab/usecases"
)

type CDHandler struct {
	usecases.AllCDsUsecase
	presenters.AllCDsPresenter
}

func (h *CDHandler) getAllCDs(c *gin.Context) {
	cds, _ := h.AllCDs(c)
	json, _ := h.ToJson(*cds)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(json.Json)
}

func (h *CDHandler) getWithId(c *gin.Context) {
	// TODO: contextからIDを取り出してGetByIDに渡す
	// info, _ := h.GetById()
	// h.GetWithID(info)
}

func NewCDHandler(r *gin.Engine, db *gorm.DB) {
	handler := CDHandler{
		usecases.NewAllCDsUsecase(db),
		presenters.NewAllCDsPresenter(db),
	}
	r.GET("/cd", handler.getAllCDs)
	r.GET("/cd/:id", handler.getWithId)
}
