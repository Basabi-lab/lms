package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestSongCreateResponse(t *testing.T) {
	pre := NewSongCreatePresenter()
	usecaseResult := usecases.TestSongCreateUsecaseResult()

	_, err := pre.ToByteList(usecaseResult)
	assert.NoError(t, err)
}
