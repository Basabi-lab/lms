package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestAlbumCreateResponse(t *testing.T) {
	pre := NewAlbumCreatePresenter()
	usecaseResult := usecases.TestAlbumCreateUsecaseResult()
	expect := TestAlbumCreatePresenterResult()

	ret, err := pre.Response(usecaseResult)
	assert.NoError(t, err)
	assert.Equal(t, expect, ret)
}
