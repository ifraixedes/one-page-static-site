package onepagestaticsite_test

import (
	"io/ioutil"
	"os"
	"testing"

	onepagestaticsite "github.com/ifraixedes/one-page-static-site"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRender(t *testing.T) {
	// nolint: lll
	const expOutput = `<html>
  <head>
    <title>Example</title>
  </head>
  <body>
    <h1>Example</h1>

<p>This is an example of <a href="https://github.com/ifraixedes/one-page-static-site">one-page-static-site</a> for testing the tool.</p>

  </body>
</html>
`
	var err = onepagestaticsite.Render(
		"testdata/layout.html", "testdata/content.md", "testdata/index.html",
	)
	require.NoError(t, err)

	outr, err := os.Open("testdata/index.html")
	require.NoError(t, err)
	output, err := ioutil.ReadAll(outr)
	require.NoError(t, err)

	assert.Equal(t, expOutput, string(output))
}

func TestRender_InvalidPaths(t *testing.T) {
	var err = onepagestaticsite.Render(
		"invalid-template.html", "testdata/content.md", "testdata/index.html",
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid-template.html")

	err = onepagestaticsite.Render(
		"testdata/layout.html", "invalid-content.md", "testdata/index.html",
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid-content.md")

	err = onepagestaticsite.Render(
		"testdata/layout.html", "testdata/content.md", "testdata/not-exist/index.html",
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "testdata/not-exist")
}
