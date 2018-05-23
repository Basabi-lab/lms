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
	sau          usecases.SongAllUsecaseExt
	sap          presenters.SongAllPresenterExt
	sgbiu        usecases.SongGetByIDUsecaseExt
	sgbip        presenters.SongGetByIDPresenterExt
	scu          usecases.SongCreateUsecaseExt
	scp          presenters.SongCreatePresenterExt
	songClearUse usecases.SongClearUsecaseExt
	songClearPre presenters.SongClearPresenterExt
}

func NewSongController(
	sau usecases.SongAllUsecaseExt,
	sap presenters.SongAllPresenterExt,
	sgbiu usecases.SongGetByIDUsecaseExt,
	sgbip presenters.SongGetByIDPresenterExt,
	scu usecases.SongCreateUsecaseExt,
	scp presenters.SongCreatePresenterExt,
	songClearUse usecases.SongClearUsecaseExt,
	songClearPre presenters.SongClearPresenterExt,
) SongControllerExt {

	return &songController{
		sau:          sau,
		sap:          sap,
		sgbiu:        sgbiu,
		sgbip:        sgbip,
		scu:          scu,
		scp:          scp,
		songClearUse: songClearUse,
		songClearPre: songClearPre,
	}
}

func (h *songController) GetAll(c *gin.Context) {
	songs, err := h.sau.All(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	json, err := h.sap.ToByteList(songs)
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
	song, err := h.sgbiu.GetByID(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	json, err := h.sgbip.ToByteList(song)
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
	cRes, err := h.scu.Create(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	pRes, err := h.scp.Response(cRes)
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.JSON(200, pRes.Res)
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
