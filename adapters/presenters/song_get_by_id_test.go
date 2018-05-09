package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestSongGetByIDToJson(t *testing.T) {
	pre := NewSongGetByIDPresenter()
	usecaseResult := usecases.TestSongGetByIDUsecaseResult()

	_, err := pre.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
