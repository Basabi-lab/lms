package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/adapters/presenters"
	"github.com/Basabi-lab/lms/usecases"
)

type AlbumControllerExt interface {
	GetAllAlbum(c *gin.Context)
	GetWithId(c *gin.Context)
}

type albumController struct {
	aau usecases.AllAlbumUsecaseExt
	aap presenters.AllAlbumPresenterExt
}

func NewAlbumController(aau usecases.AllAlbumUsecaseExt, aap presenters.AllAlbumPresenterExt) AlbumControllerExt {
	return &albumController{
		aau: aau,
		aap: aap,
	}
}

func ResponseError(c *gin.Context, err error) {
	c.Error(err)
	c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
}

func (h *albumController) GetAllAlbum(c *gin.Context) {
	albums, err := h.aau.AllAlbum(c)
	if err != nil {
		ResponseError(c, err)
	}

	json, err := h.aap.ToByteList(*albums)
	if err != nil {
		ResponseError(c, err)
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	_, err = c.Writer.Write(json.JsonByteList)
	if err != nil {
		ResponseError(c, err)
	}
}

func (h *albumController) GetWithId(c *gin.Context) {
	// TODO: contextからIDを取り出してGetByIDに渡す
	// info, _ := h.GetById()
	// h.GetWithID(info)
	ResponseError(c, errors.New("not implement"))
}
