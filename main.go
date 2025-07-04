package main

import (
	"fmt"
	"log"

	"github.com/yenxxxw/image-processing-cli/converter"
	"github.com/yenxxxw/image-processing-cli/input"
)

func main() {

	fmt.Println("Welcome! Please enter the file path(s) to be converted.")
	fmt.Println("You can also drag and drop the files into the terminal.")

	result, err := input.PromptType()
	if err != nil {
		log.Panic("Error taking convertion type", err)
	}

	var inputPath string
	var height uint
	var width uint

	var toConverts []string

	if result == "Scale" {
		inputPath, width, height, err = input.PromptUserInputScaling()
		if err != nil {
			log.Panic("Error taking user input for scaling", err)

		}
	} else {

		toConverts, err = input.PromptUserInputConvertion(result)
		if err != nil {
			log.Panic("Error taking user input", err)
		}
	}

	outPath, err := input.PromptOutputPath()
	if err != nil {
		log.Println("Error taking output path", err)
	}

	if result == "Scale" {
		err := converter.ScaleImage(inputPath, outPath, width, height)
		if err != nil {
			log.Printf("Error scaling image %v", err)
		}

	} else {

		wp := converter.WorkerPool{
			Images:      toConverts,
			Concurrency: 5,
		}

		wp.Run(outPath, result)
	}

}
