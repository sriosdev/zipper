package zipper

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func addFiles(directory *os.File, w *zip.Writer, baseInZip string) error {
	rootDirectory := directory.Name()
	files, err := directory.Readdir(-1)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			f, err := ioutil.ReadFile(rootDirectory + "/" + file.Name())
			if err != nil {
				return err
			}

			fw, err := w.Create(baseInZip + file.Name())
			if err != nil {
				return err
			}
			_, err = fw.Write(f)
			if err != nil {
				return err
			}
		} else if file.IsDir() {
			// Recursive call to add files in nested folders
			nestedFolder, err := os.Open(rootDirectory + file.Name() + "/")
			defer nestedFolder.Close()
			if err != nil {
				return err
			}

			// New relative path of nested folder inside ZIP
			newBaseInZip := baseInZip + file.Name() + "/"

			if err := addFiles(nestedFolder, w, newBaseInZip); err != nil {
				return err
			}
		}
	}

	return nil
}

// ZipFolder makes a ZIP file including all nested files and folders inside a directory given.
// Empty folders will be ignored.
func ZipFolder(directory *os.File) (*os.File, error) {
	// Check path is a directory
	fi, err := os.Stat(directory.Name())
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, errors.New("Path must be a directory not a file")
	}

	zipfile, err := os.Create(filepath.Base(directory.Name()) + ".zip")
	defer zipfile.Close()
	if err != nil {
		return nil, err
	}

	w := zip.NewWriter(zipfile)
	defer w.Close()

	if err := addFiles(directory, w, ""); err != nil {
		return nil, err
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return zipfile, nil
}
