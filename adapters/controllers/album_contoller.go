package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/adapters/presenters"
	"github.com/Basabi-lab/lms/usecases"
)

type AlbumControllerExt interface {
	GetAllAlbum(c *gin.Context)
	GetWithId(c *gin.Context)
}

type albumController struct {
	aau   usecases.AllAlbumUsecaseExt
	aap   presenters.AllAlbumPresenterExt
	agbiu usecases.AlbumGetByIDUsecaseExt
	agbip presenters.AlbumGetByIDPresenterExt
}

func NewAlbumController(
	aau usecases.AllAlbumUsecaseExt,
	aap presenters.AllAlbumPresenterExt,
	agbiu usecases.AlbumGetByIDUsecaseExt,
	agbip presenters.AlbumGetByIDPresenterExt,
) AlbumControllerExt {

	return &albumController{
		aau:   aau,
		aap:   aap,
		agbiu: agbiu,
		agbip: agbip,
	}
}

func (h *albumController) GetAllAlbum(c *gin.Context) {
	albums, err := h.aau.All(c)
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
	album, err := h.agbiu.GetByID(c)
	if err != nil {
		ResponseError(c, err)
	}

	json, err := h.agbip.ToByteList(*album)
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
