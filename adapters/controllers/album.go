package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/adapters/presenters"
	"github.com/Basabi-lab/lms/usecases"
)

type AlbumControllerExt interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Post(c *gin.Context)
}

type albumController struct {
	aau   usecases.AlbumAllUsecaseExt
	aap   presenters.AlbumAllPresenterExt
	agbiu usecases.AlbumGetByIDUsecaseExt
	agbip presenters.AlbumGetByIDPresenterExt
	acu   usecases.AlbumCreateUsecaseExt
	acp   presenters.AlbumCreatePresenterExt
}

func NewAlbumController(
	aau usecases.AlbumAllUsecaseExt,
	aap presenters.AlbumAllPresenterExt,
	agbiu usecases.AlbumGetByIDUsecaseExt,
	agbip presenters.AlbumGetByIDPresenterExt,
	acu usecases.AlbumCreateUsecaseExt,
	acp presenters.AlbumCreatePresenterExt,
) AlbumControllerExt {

	return &albumController{
		aau:   aau,
		aap:   aap,
		agbiu: agbiu,
		agbip: agbip,
		acu:   acu,
		acp:   acp,
	}
}

func (h *albumController) GetAll(c *gin.Context) {
	albums, err := h.aau.All(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	json, err := h.aap.ToByteList(albums)
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	_, err = c.Writer.Write(json.JsonByteList)
	if err != nil {
		ResponseError(c, err)
		return
	}
}

func (h *albumController) GetById(c *gin.Context) {
	album, err := h.agbiu.GetByID(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	json, err := h.agbip.ToByteList(album)
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	_, err = c.Writer.Write(json.JsonByteList)
	if err != nil {
		ResponseError(c, err)
		return
	}
}

func (h *albumController) Post(c *gin.Context) {
	cRes, err := h.acu.Create(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	pRes, err := h.acp.Response(cRes)
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.JSON(200, pRes.Res)
}
