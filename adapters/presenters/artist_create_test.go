package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestArtistCreateResponse(t *testing.T) {
	pre := NewArtistCreatePresenter()
	usecaseResult := usecases.TestArtistCreateUsecaseResult()

	_, err := pre.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
