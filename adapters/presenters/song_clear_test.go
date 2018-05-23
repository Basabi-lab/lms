package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestSongClearResponse(t *testing.T) {
	pre := NewSongClearPresenter()
	usecaseResult := usecases.TestSongClearUsecaseResult()

	_, err := pre.Response(usecaseResult)
	assert.NoError(t, err)
}
