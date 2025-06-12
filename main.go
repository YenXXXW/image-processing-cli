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

	inputPath, err := input.PromptUserInput(result)
	if err != nil {
		log.Panic("Error taking user input", err)
	}

	toConverts, err := converter.ExtractFiles(inputPath, result)
	if err != nil {
		log.Panic("Validation error:", err)
	}

	outPath, err := input.PromptOutputPath()
	if err != nil {
		log.Panic("Error taking output path", err)
	}

	if result == "JPG to PNG" {
		for _, path := range toConverts {
			err := converter.ConvertToPNG(path, outPath)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else if result == "PNG to JPG" {

		for _, path := range toConverts {
			err := converter.ConvertToJPG(path, outPath)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
