package converter

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func ConvertToPNG(path, outPath string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	image, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	var buff bytes.Buffer

	err = png.Encode(&buff, image)
	if err != nil {
		return err
	}

	fileName := strings.Split(filepath.Base(path), ".")[0]

	file, err = os.Create(fmt.Sprintf("%s/%s.png", outPath, fileName))
	if err != nil {
		return err
	}
	_, err = file.Write(buff.Bytes())
	if err != nil {
		return err
	}
	fmt.Println("successfully converted the image from jpeg to png")

	return nil
}

func ConvertToJPG(path, outPath string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	image, err := png.Decode(file)
	if err != nil {
		return err
	}

	var buff bytes.Buffer

	err = jpeg.Encode(&buff, image, nil)
	if err != nil {
		return err
	}

	fileName := strings.Split(filepath.Base(path), ".")[0]

	file, err = os.Create(fmt.Sprintf("%s/%s.jpeg", outPath, fileName))
	if err != nil {
		return err
	}
	_, err = file.Write(buff.Bytes())
	if err != nil {
		return err
	}
	fmt.Println("successfully converted the image from png to jpeg")

	return nil
}
