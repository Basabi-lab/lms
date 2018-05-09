package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestAlbumGetByIDToJson(t *testing.T) {
	pre := NewAlbumGetByIDPresenter()
	usecaseResult := usecases.TestAlbumGetByIDUsecaseResult()

	_, err := pre.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
