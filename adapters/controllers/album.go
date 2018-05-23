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
	Clear(c *gin.Context)
}

type albumController struct {
	albumAllUse     usecases.AlbumAllUsecaseExt
	albumAllPre     presenters.AlbumAllPresenterExt
	albumGetByIDUse usecases.AlbumGetByIDUsecaseExt
	albumGetByIDPre presenters.AlbumGetByIDPresenterExt
	albumCreateUse  usecases.AlbumCreateUsecaseExt
	albumCreatePre  presenters.AlbumCreatePresenterExt
	albumClearUse   usecases.AlbumClearUsecaseExt
	albumClearPre   presenters.AlbumClearPresenterExt
}

func NewAlbumController(
	albumAllUse usecases.AlbumAllUsecaseExt,
	albumAllPre presenters.AlbumAllPresenterExt,
	albumGetByIDUse usecases.AlbumGetByIDUsecaseExt,
	albumGetByIDPre presenters.AlbumGetByIDPresenterExt,
	albumCreateUse usecases.AlbumCreateUsecaseExt,
	albumCreatePre presenters.AlbumCreatePresenterExt,
	albumClearUse usecases.AlbumClearUsecaseExt,
	albumClearPre presenters.AlbumClearPresenterExt,
) AlbumControllerExt {

	return &albumController{
		albumAllUse:     albumAllUse,
		albumAllPre:     albumAllPre,
		albumGetByIDUse: albumGetByIDUse,
		albumGetByIDPre: albumGetByIDPre,
		albumCreateUse:  albumCreateUse,
		albumCreatePre:  albumCreatePre,
		albumClearUse:   albumClearUse,
		albumClearPre:   albumClearPre,
	}
}

func (h *albumController) GetAll(c *gin.Context) {
	albums, err := h.albumAllUse.All(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	json, err := h.albumAllPre.ToByteList(albums)
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
	album, err := h.albumGetByIDUse.GetByID(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	json, err := h.albumGetByIDPre.ToByteList(album)
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
	cRes, err := h.albumCreateUse.Create(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	pRes, err := h.albumCreatePre.Response(cRes)
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.JSON(200, pRes.Res)
}

func (h *albumController) Clear(c *gin.Context) {
	cRes, err := h.albumClearUse.Clear(c)
	if err != nil {
		ResponseError(c, err)
		return
	}

	pRes, err := h.albumClearPre.Response(cRes)
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.JSON(200, pRes.Res)
}
