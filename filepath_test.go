package onepagestaticsite

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolveFilepath_Exists(t *testing.T) {

	var (
		expafp string
		fp     = "render.go"
	)
	{
		var wd, err = os.Getwd()
		require.NoError(t, err)
		expafp = filepath.Join(wd, fp)
	}

	// Relative path
	var afp, err = resolveFilepath(fp)
	assert.NoError(t, err)
	assert.Equal(t, expafp, afp)

	// Absolute path
	afp, err = resolveFilepath(expafp)
	assert.NoError(t, err)
	assert.Equal(t, expafp, afp)
}

func TestResolveFilepath_NotExists(t *testing.T) {
	var fp string
	{
		var fpr, err = ioutil.TempDir("", "one-page-static-site-test")
		require.NoError(t, err)
		fp, err = ioutil.TempDir(fpr, "subdir")
		require.NoError(t, err)
		err = os.RemoveAll(fpr)
		require.NoError(t, err)
	}

	var _, err = resolveFilepath(fp)
	assert.Error(t, err)
}

func TestResolveFilepath_NoNormalFile(t *testing.T) {
	var slk = fmt.Sprintf("one-page-static-site-test-%d", time.Now().UnixNano())
	//nolint:errcheck
	defer os.Remove(slk)

	{
		var fp, err = ioutil.TempFile("", "one-page-static-site-test")
		require.NoError(t, err)
		require.NoError(t, fp.Close())
		err = os.Symlink(fp.Name(), slk)
		require.NoError(t, err)

		//nolint:errcheck
		defer os.Remove(fp.Name())
	}

	var _, err = resolveFilepath(slk)
	assert.Error(t, err)
}

func TestResolveFilepath_DirPath(t *testing.T) {
	var dpath = "cmd"
	_, err := resolveFilepath(dpath)
	assert.Error(t, err)
}
