package presenters

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestAlbumCreateResponse(t *testing.T) {
	pre := NewAlbumCreatePresenter()
	usecaseResult := &usecases.AlbumCreateUsecaseResult{
		ID: uint(0),
	}

	expect := &AlbumCreatePresenterResult{
		Res: &gin.H{"message": "success", "id": uint(0)},
	}

	ret, err := pre.Response(usecaseResult)
	assert.NoError(t, err)
	assert.Equal(t, expect, ret)
}
