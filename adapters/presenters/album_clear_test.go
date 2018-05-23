package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestAlbumClearResponse(t *testing.T) {
	pre := NewAlbumClearPresenter()
	usecaseResult := usecases.TestAlbumClearUsecaseResult()

	_, err := pre.Response(usecaseResult)
	assert.NoError(t, err)
}
