package onepagestaticsite

import (
	"fmt"
	"io"
	"io/ioutil"

	blackfriday "github.com/russross/blackfriday/v2"
)

func contentToHTML(r io.Reader) (string, error) {
	ctn, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("error reading the entire data of the content file. %+v", err)
	}

	return string(blackfriday.Run(ctn)), nil
}
