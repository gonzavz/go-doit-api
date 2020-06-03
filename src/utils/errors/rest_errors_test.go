package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBadRequestError(t *testing.T) {
	assert := assert.New(t)
	badRequest := NewBadRequestError("Error Message")
	assert.Equal(400, badRequest.Status())
	assert.Equal("Error Message", badRequest.Message())
	assert.Equal("message: Error Message - status: 400 - error: bad_request - causes: []", badRequest.Error())
}
