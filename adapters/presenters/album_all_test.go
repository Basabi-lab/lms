package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestAlbumAllToJson(t *testing.T) {
	pre := NewAlbumAllPresenter()
	usecaseResult := usecases.TestAlbumAllUsecaseResult()

	_, err := pre.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
