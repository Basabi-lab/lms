package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestArtistGetByIDToJson(t *testing.T) {
	pre := NewArtistGetByIDPresenter()
	usecaseResult := usecases.TestArtistGetByIDUsecaseResult()

	_, err := pre.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
