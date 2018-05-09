package presenters

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestSongCreateResponse(t *testing.T) {
	aap := NewSongCreatePresenter()
	usecaseResult := &usecases.SongCreateUsecaseResult{
		ID: uint(0),
	}

	expect := &SongCreatePresenterResult{
		Res: &gin.H{"message": "success", "id": uint(0)},
	}

	ret, err := aap.Response(usecaseResult)
	assert.NoError(t, err)
	assert.Equal(t, expect, ret)
}
