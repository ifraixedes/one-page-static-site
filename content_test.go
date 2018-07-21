package onepagestaticsite

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContentToHTML(t *testing.T) {
	var ctnr = bytes.NewReader([]byte("# Markdown title"))
	var html, err = contentToHTML(ctnr)
	require.NoError(t, err)
	assert.Equal(t, "<h1>Markdown title</h1>\n", html)
}
