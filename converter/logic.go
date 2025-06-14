package converter

import (
	"errors"
	"fmt"
	"github.com/google/shlex"
	"os"
	"path/filepath"
)

func ExtractFiles(input, mode string) ([]string, error) {

	toConverts := []string{}
	paths, err := shlex.Split(input)
	if err != nil {
		return nil, err
	}

	if len(paths) == 0 {
		return nil, errors.New("No path provided")
	}

	for _, path := range paths {
		ext := filepath.Ext(path)
		switch mode {
		case "JPG to PNG":
			if ext == ".jpg" || ext == ".jpeg" {
				toConverts = append(toConverts, path)
			}
		case "PNG to JPG":
			if ext == ".png" || ext == ".PNG" {
				toConverts = append(toConverts, path)
			}
		case "Scale":
			toConverts = append(toConverts, path)
			break
		}
	}

	if len(toConverts) == 0 {
		return nil, errors.New("the required type of image is not provided")
	}

	for _, path := range toConverts {
		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			return []string{}, errors.New(fmt.Sprintf("%s does not exist", path))
		} else if err != nil {
			return []string{}, err
		}
	}

	return toConverts, nil
}
