package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructUrl(t *testing.T) {
	urlExpected := "https://host/path?field=value"

	urlGot := constructUrl(
		"https",
		"host",
		"path",
		map[string]string{
			"field": "value"})

	assert.Equal(t, urlExpected, urlGot.String())
}
