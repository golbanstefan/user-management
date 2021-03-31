package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckError(t *testing.T) {
	err := errors.New("test")
	assert.Panics(t, func() {
		CheckError(err)
	})

}
