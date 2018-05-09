package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestArtistAllToJson(t *testing.T) {
	pre := NewArtistAllPresenter()
	usecaseResult := usecases.TestArtistAllUsecaseResult()

	_, err := pre.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
