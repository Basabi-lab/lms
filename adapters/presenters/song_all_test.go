package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestSongAllToJson(t *testing.T) {
	pre := NewSongAllPresenter()
	usecaseResult := usecases.TestSongAllUsecaseResult()

	_, err := pre.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
