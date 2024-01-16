package fakeopen

import (
	"fmt"
	"io"
	"testing"

	"github.com/aandrew-me/tgpt/v2/structs"
	"github.com/stretchr/testify/assert"
)

// Test for fakeopen NewRequest

func TestRequest(t *testing.T) {
	resp, err := NewRequest("What is 1+1", structs.Params{}, "")
	body, _ := io.ReadAll(resp.Body)

	assert := assert.New(t)

	fmt.Println("Statuscode:", resp.StatusCode)
	fmt.Println("Response:", string(body))	

	assert.Nil(err, "Error should be Nil")
	assert.Equal(200, resp.StatusCode, "Response status code should be 200")
}