package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CDHandler struct {
	cu usecases.CDUsecase
}

func (h *CDHandler) getWithId(c *gin.Context) {
	// TODO: contextからIDを取り出してGetByIDに渡す
	info, _ := h.cu.GetById()
	presenters.GetWithID(info)
}

func (h *CDHandler) post(c *gin.Context) {
	h.cu.Create(c)
	presenters.Create()
}

func NewCDHandler(r *gin.Engine, db *gorm.DB) {
	handler := CDHandler{
		cu: usecases.NewCDUsecase(dbs.NewCDMysql(db)),
	}
	r.GET("/cd/:id", handler.getWithId())
	r.POST("/cd", handler.post())
}
