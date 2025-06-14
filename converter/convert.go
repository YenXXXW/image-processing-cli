package converter

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func ConvertToPNG(path, outPath string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()
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

	outFile, err := os.Create(fmt.Sprintf("%s/%s.png", outPath, fileName))
	if err != nil {
		return err
	}

	defer outFile.Close()

	_, err = outFile.Write(buff.Bytes())
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

	defer file.Close()

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

	outFile, err := os.Create(fmt.Sprintf("%s/%s.jpeg", outPath, fileName))
	if err != nil {
		return err
	}

	defer outFile.Close()
	_, err = outFile.Write(buff.Bytes())
	if err != nil {
		return err
	}
	fmt.Println("successfully converted the image from png to jpeg")

	return nil
}

func ScaleImage(path, outPath string, width, height uint) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	file.Close()

	base := filepath.Base(path)
	ext := filepath.Ext(base)
	fileName := strings.TrimSuffix(base, ext)

	var buff bytes.Buffer

	if ext == "jpeg" || ext == "jpg" {

		img, err := jpeg.Decode(file)

		scaledImage := resize.Resize(width, height, img, resize.Lanczos3)

		err = jpeg.Encode(&buff, scaledImage, nil)
		if err != nil {
			return err
		}
	}

	if ext == "PNG" || ext == "png" {
		img, err := png.Decode(file)

		scaledImage := resize.Resize(width, height, img, resize.Lanczos3)

		err = png.Encode(&buff, scaledImage)
		if err != nil {
			return err
		}
	}

	outFile, err := os.Create(fmt.Sprintf("%s/%sScaled.%s", outPath, fileName, ext))
	if err != nil {
		return err
	}

	defer outFile.Close()
	_, err = outFile.Write(buff.Bytes())
	if err != nil {
		return err
	}

	return nil

}
