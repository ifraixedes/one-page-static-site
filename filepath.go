package onepagestaticsite

import (
	"fmt"
	"os"
	"path/filepath"
)

// resolveFilepath checks if the file with the filepath fp exists, if not exists
// checks that the directory which should contain the file exists, then resolve
// the absolute path of the file path or directory path.
// If fp is a file path to a file which exists and is absolute the same fp is
// returned.
// It returns an error when:
// * file path (or the directory which should contain the file) which  doesn't
//   exist
// * file path isn't of a regular file
// * any unexpected error returned by the called function of the os package.
//   This errors will be wrapped by message indicating about the operation
//   which has failed plus the stringified form of the original error.
func resolveFilepath(fp string) (string, error) {
	var fi, err = os.Lstat(fp)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", fmt.Errorf("error reading the file information: %+v", err)
		}

		var dirp = filepath.Dir(fp)
		fi, err = os.Lstat(dirp)
		if err != nil {
			if !os.IsNotExist(err) {
				return "", fmt.Errorf("error reading the directory information: %+v", err)
			}

			return "", fmt.Errorf("the directory (%s) doesn't exist", dirp)
		}
	} else {
		var fm = fi.Mode()
		if !fm.IsRegular() {
			return "", fmt.Errorf(
				"file path (%s) is from a non-regular, mode: %s", fp, fm.String(),
			)
		}
	}

	if filepath.IsAbs(fp) {
		return fp, nil
	}

	var dirp string
	dirp, err = os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting the root directory: %+v", err)
	}

	return filepath.Join(dirp, fp), nil
}
