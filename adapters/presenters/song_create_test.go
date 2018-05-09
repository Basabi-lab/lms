package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/usecases"
)

func TestSongCreateResponse(t *testing.T) {
	pre := NewSongCreatePresenter()
	usecaseResult := usecases.TestSongCreateUsecaseResult()
	expect := TestSongCreatePresenterResult()

	ret, err := pre.Response(usecaseResult)
	assert.NoError(t, err)
	assert.Equal(t, expect, ret)
}
