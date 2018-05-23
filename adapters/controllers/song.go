package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/Basabi-lab/lms/adapters/presenters"
	"github.com/Basabi-lab/lms/usecases"
)

type SongControllerExt interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Post(c *gin.Context)
	Clear(c *gin.Context)
}

type songController struct {
	songAllUse     usecases.SongAllUsecaseExt
	songAllPre     presenters.SongAllPresenterExt
	songGetByIDUse usecases.SongGetByIDUsecaseExt
	songGetByIDPre presenters.SongGetByIDPresenterExt
	songCreateUse  usecases.SongCreateUsecaseExt
	songCreatePre  presenters.SongCreatePresenterExt
	songClearUse   usecases.SongClearUsecaseExt
	songClearPre   presenters.SongClearPresenterExt
}

func NewSongController(
	songAllUse usecases.SongAllUsecaseExt,
	songAllPre presenters.SongAllPresenterExt,
	songGetByIDUse usecases.SongGetByIDUsecaseExt,
	songGetByIDPre presenters.SongGetByIDPresenterExt,
	songCreateUse usecases.SongCreateUsecaseExt,
	songCreatePre presenters.SongCreatePresenterExt,
	songClearUse usecases.SongClearUsecaseExt,
	songClearPre presenters.SongClearPresenterExt,
) SongControllerExt {

	return &songController{
		songAllUse:     songAllUse,
		songAllPre:     songAllPre,
		songGetByIDUse: songGetByIDUse,
		songGetByIDPre: songGetByIDPre,
		songCreateUse:  songCreateUse,
		songCreatePre:  songCreatePre,
		songClearUse:   songClearUse,
		songClearPre:   songClearPre,
	}
}

func (h *songController) GetAll(c *gin.Context) {
	songs, err := h.songAllUse.All(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	json, err := h.songAllPre.ToByteList(songs)
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

func (h *songController) GetById(c *gin.Context) {
	song, err := h.songGetByIDUse.GetByID(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	json, err := h.songGetByIDPre.ToByteList(song)
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

func (h *songController) Post(c *gin.Context) {
	cRes, err := h.songCreateUse.Create(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	pRes, err := h.songCreatePre.ToByteList(cRes)
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

func (h *songController) Clear(c *gin.Context) {
	cRes, err := h.songClearUse.Clear(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	pRes, err := h.songClearPre.Response(cRes)
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.JSON(200, pRes.Res)
}
