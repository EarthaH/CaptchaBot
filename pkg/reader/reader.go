package reader

import (
	"path/filepath"

	"github.com/otiai10/gosseract/v2"
)

var homepath = "./resources/images"

func ReadImage(filename string) (string, error) {
	fpath := filepath.Join(homepath, filename)

	client := gosseract.NewClient()
	defer client.Close()
	err := client.SetImage(fpath)
	if err != nil {
		return "", err
	}

	text, err := client.Text()
	if err != nil {
		return "", err
	}

	return text, nil
}
