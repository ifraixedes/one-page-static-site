package onepagestaticsite

// Render generate the HTML file on the outfp file path using the Go HTML
// template on the tplfp file path and placing the HTML content in the
// {{.Content}} of the template, which is obtained from the Markdown file in the
// cntfp file path.
// NOTE that if the outfp file exists, it will be overwritten.
// It returns an errors if the template file doesn't exits, if any of the
// directories paths cntfp or outfp don't exists or any error happens during the
// operations performed.
func Render(tplfp string, cntfp string, outfp string) error {
	// TODO: Implement this function
	return nil
}
