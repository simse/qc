package source

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Scan returns all files in directory
func Scan(root string, recursive bool, filter []string) ([]File, error) {
	var files []File
	var fileError error
	var filteredFiles []File

	if recursive {
		files, fileError = scanRecursively(root)
	} else {
		files, fileError = scanNonRecursively(root)
	}

	for _, file := range files {
		for _, extension := range filter {
			if extension == file.Extension || filter[0] == "*" {
				filteredFiles = append(filteredFiles, file)
			}
		}
	}

	return filteredFiles, fileError
}

func scanNonRecursively(root string) ([]File, error) {
	var files []File

	fileArray, scanError := ioutil.ReadDir(root)

	if scanError != nil {
		return files, scanError
	}

	//workdir, _ := os.Getwd()
	//fullPath := filepath.Join(workdir, root)

	for _, f := range fileArray {
		if !f.IsDir() {
			// Get extensions
			var extension string
			if len(filepath.Ext(f.Name())) < 1 {
				extension = ""
			} else {
				extension = filepath.Ext(f.Name())[1:]
			}

			files = append(files, File{
				Path:      filepath.Join(root, f.Name()),
				Extension: extension,
				Key:       filepath.Base(f.Name()),
			})
		}
	}

	return files, nil
}

func scanRecursively(root string) ([]File, error) {
	var files []File

	//workdir, _ := os.Getwd()
	//fullPath := filepath.Join(workdir, root)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		// Get extensions
		var extension string
		if len(filepath.Ext(info.Name())) < 1 {
			extension = ""
		} else {
			extension = filepath.Ext(info.Name())[1:]
		}

		files = append(files, File{
			Path:      filepath.Join(root, info.Name()),
			Extension: extension,
			Key:       filepath.Base(path),
		})

		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}

	return files, nil
}
