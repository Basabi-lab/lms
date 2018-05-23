package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestArtistClearResponse(t *testing.T) {
	pre := NewArtistClearPresenter()
	usecaseResult := usecases.TestArtistClearUsecaseResult()

	_, err := pre.Response(usecaseResult)
	assert.NoError(t, err)
}
