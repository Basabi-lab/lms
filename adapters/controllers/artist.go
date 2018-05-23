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
	Clear(c *gin.Context)
}

type artistController struct {
	artistAllUse     usecases.ArtistAllUsecaseExt
	artistAllPre     presenters.ArtistAllPresenterExt
	artistGetByIDUse usecases.ArtistGetByIDUsecaseExt
	artistGetByIDPre presenters.ArtistGetByIDPresenterExt
	artistCreateUse  usecases.ArtistCreateUsecaseExt
	artistCreatePre  presenters.ArtistCreatePresenterExt
	artistClearUse   usecases.ArtistClearUsecaseExt
	artistClearPre   presenters.ArtistClearPresenterExt
}

func NewArtistController(
	artistAllUse usecases.ArtistAllUsecaseExt,
	artistAllPre presenters.ArtistAllPresenterExt,
	artistGetByIDUse usecases.ArtistGetByIDUsecaseExt,
	artistGetByIDPre presenters.ArtistGetByIDPresenterExt,
	artistCreateUse usecases.ArtistCreateUsecaseExt,
	artistCreatePre presenters.ArtistCreatePresenterExt,
	artistClearUse usecases.ArtistClearUsecaseExt,
	artistClearPre presenters.ArtistClearPresenterExt,
) ArtistControllerExt {

	return &artistController{
		artistAllUse:     artistAllUse,
		artistAllPre:     artistAllPre,
		artistGetByIDUse: artistGetByIDUse,
		artistGetByIDPre: artistGetByIDPre,
		artistCreateUse:  artistCreateUse,
		artistCreatePre:  artistCreatePre,
		artistClearUse:   artistClearUse,
		artistClearPre:   artistClearPre,
	}
}

func (h *artistController) GetAll(c *gin.Context) {
	artists, err := h.artistAllUse.All(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	json, err := h.artistAllPre.ToByteList(artists)
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
	artist, err := h.artistGetByIDUse.GetByID(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	json, err := h.artistGetByIDPre.ToByteList(artist)
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
	cRes, err := h.artistCreateUse.Create(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	pRes, err := h.artistCreatePre.ToByteList(cRes)
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	_, err = c.Writer.Write(pRes.JsonByteList)
	if err != nil {
		ResponseError(c, err)
		return
	}
}

func (h *artistController) Clear(c *gin.Context) {
	cRes, err := h.artistClearUse.Clear(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	pRes, err := h.artistClearPre.Response(cRes)
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.JSON(200, pRes.Res)
}
