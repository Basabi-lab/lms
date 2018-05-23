package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestAlbumCreateResponse(t *testing.T) {
	pre := NewAlbumCreatePresenter()
	usecaseResult := usecases.TestAlbumCreateUsecaseResult()

	_, err := pre.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
