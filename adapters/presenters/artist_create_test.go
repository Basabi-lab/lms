package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestArtistCreateResponse(t *testing.T) {
	pre := NewArtistCreatePresenter()
	usecaseResult := usecases.TestArtistCreateUsecaseResult()
	expect := TestArtistCreatePresenterResult()

	ret, err := pre.Response(usecaseResult)
	assert.NoError(t, err)
	assert.Equal(t, expect, ret)
}
