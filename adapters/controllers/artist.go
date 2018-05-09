package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/adapters/presenters"
	"github.com/Basabi-lab/lms/usecases"
)

type ArtistControllerExt interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Post(c *gin.Context)
}

type artistController struct {
	aau   usecases.ArtistAllUsecaseExt
	aap   presenters.ArtistAllPresenterExt
	agbiu usecases.ArtistGetByIDUsecaseExt
	agbip presenters.ArtistGetByIDPresenterExt
	acu   usecases.ArtistCreateUsecaseExt
	acp   presenters.ArtistCreatePresenterExt
}

func NewArtistController(
	aau usecases.ArtistAllUsecaseExt,
	aap presenters.ArtistAllPresenterExt,
	agbiu usecases.ArtistGetByIDUsecaseExt,
	agbip presenters.ArtistGetByIDPresenterExt,
	acu usecases.ArtistCreateUsecaseExt,
	acp presenters.ArtistCreatePresenterExt,
) ArtistControllerExt {

	return &artistController{
		aau:   aau,
		aap:   aap,
		agbiu: agbiu,
		agbip: agbip,
		acu:   acu,
		acp:   acp,
	}
}

func (h *artistController) GetAll(c *gin.Context) {
	artists, err := h.aau.All(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	json, err := h.aap.ToByteList(artists)
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

func (h *artistController) GetById(c *gin.Context) {
	artist, err := h.agbiu.GetByID(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	json, err := h.agbip.ToByteList(artist)
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

func (h *artistController) Post(c *gin.Context) {
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
