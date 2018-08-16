package onepagestaticsite

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	"github.com/petergtz/pegomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:generate pegomock generate --package onepagestaticsite -o mock_io_reader_test.go io Reader

func TestContentToHTML(t *testing.T) {
	var ctnr = bytes.NewReader([]byte("# Markdown title"))
	var html, err = contentToHTML(ctnr)
	require.NoError(t, err)
	assert.Equal(t, "<h1>Markdown title</h1>\n", html)
}

func TestContentToHTML_IOError(t *testing.T) {
	var r = NewMockReader()
	pegomock.When(r.Read(AnyByteSlice())).ThenReturn(0, errors.New("reader failed"))

	var _, err = contentToHTML(r)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "reader failed")
	r.VerifyWasCalledOnce()
}

func AnyByteSlice() []byte {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf([]byte{})))
	return []byte{}
}
