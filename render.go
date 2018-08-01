package onepagestaticsite

import (
	"fmt"
	"os"
	"text/template"
)

// Render generate the HTML file on the outfp file path using the Go HTML
// template on the tplfp file path and placing the HTML content in the
// {{.Content}} of the template, which is obtained from the Markdown file in the
// cntfp file path.
// NOTE that if the outfp file exists, it will be overwritten.
// It returns an errors if the template file doesn't exits, if any of the
// files paths cntfp or outfp don't exists or any error happens during the
// operations performed.
func Render(tplfp string, ctnfp string, outfp string) error {
	tplfp, err := resolveFilepath(tplfp)
	if err != nil {
		return err
	}

	tpl, err := template.ParseFiles(tplfp)
	if err != nil {
		return fmt.Errorf("Error parsing the templage: %+v", err)
	}

	fpa, err := resolveFilepath(ctnfp)
	if err != nil {
		return err
	}

	ctnf, err := os.Open(fpa) // nolint: gosec
	if err != nil {
		return fmt.Errorf("Error opening the content file. %+v", err)
	}

	ctn, err := contentToHTML(ctnf)
	if err != nil {
		return err
	}

	outfp, err = resolveFilepath(outfp)
	if err != nil {
		return err
	}

	outf, err := os.Create(outfp)
	if err != nil {
		return fmt.Errorf("Error creating the output file: %+v", err)
	}

	if err = tpl.Execute(outf, tplData{Content: ctn}); err != nil {
		return fmt.Errorf("Error rendering the template: %+v", err)
	}

	return nil
}

type tplData struct {
	Content string
}
