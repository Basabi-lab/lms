package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/test"
	"github.com/Basabi-lab/lms/usecases"
)

func TestAlbumGetByIDToJson(t *testing.T) {
	pre := NewAlbumGetByIDPresenter()
	album := test.TestAlbumData()
	usecaseResult := &usecases.AlbumGetByIDUsecaseResult{
		Album: album,
	}

	_, err := pre.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
