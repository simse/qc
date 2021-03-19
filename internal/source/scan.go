package source

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Scan returns all files in directory
func Scan(root string, recursive bool, filter []string, filesChannel chan File) error {
	files := make(chan File)
	var fileError error

	if recursive {
		go func() { fileError = scanRecursively(root, files) }()
	} else {
		go func() { fileError = scanNonRecursively(root, files) }()
	}

	// Listen for files sent on channel
	for file := range files {
		// Filter files according to filter
		for _, extension := range filter {
			if extension[0] == '-' {
				if trimFirstRune(extension) == file.Extension {
					continue
				}
			}

			if extension == file.Extension || filter[0] == "*" {
				filesChannel <- file
			}
		}
	}

	close(filesChannel)

	return fileError
}

func scanNonRecursively(root string, filesChannel chan File) error {
	fileArray, scanError := ioutil.ReadDir(root)

	if scanError != nil {
		return scanError
	}

	//workdir, _ := os.Getwd()
	//fullPath := filepath.Join(workdir, root)

	for _, f := range fileArray {
		if !f.IsDir() {
			filesChannel <- File{
				Path:      filepath.Join(root, f.Name()),
				Extension: GetExtension(f.Name(), true),
				Key:       filepath.Base(f.Name()),
			}
		}
	}

	close(filesChannel)

	return nil
}

func scanRecursively(root string, filesChannel chan File) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		filesChannel <- File{
			Path:      filepath.Join(root, info.Name()),
			Extension: GetExtension(info.Name(), true),
			Key:       filepath.Base(path),
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	close(filesChannel)
	return nil
}

// GetExtension returns file extension given just file name or full path.
// hey, it's a free country (maybe?) do what you want
func GetExtension(name string, lowercase bool) string {
	var extension string
	if len(filepath.Ext(name)) < 1 {
		extension = ""
	} else {
		extension = filepath.Ext(name)[1:]
	}

	if lowercase {
		return strings.ToLower(extension)
	}

	return extension
}

func trimFirstRune(s string) string {
	for i := range s {
		if i > 0 {
			// The value i is the index in s of the second
			// rune.  Slice to remove the first rune.
			return s[i:]
		}
	}
	// There are 0 or 1 runes in the string.
	return ""
}
