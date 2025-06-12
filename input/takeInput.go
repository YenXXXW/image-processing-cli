package input

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
)

func PromptType() (string, error) {

	prompt := promptui.Select{
		Label: "Choose type",
		Items: []string{"JPG to PNG", "PNG to JPG", "Scale"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	fmt.Printf("You choose %q\n", result)
	return result, nil
}

func PromptUserInput(convertionType string) (string, error) {

	var label string
	switch convertionType {
	case "JPG to PNG":
		label = "Enter image file path(s) to convert to PNG"
	case "PNG to JPG":
		label = "Enter image file path(s) to convert to JPG"
	case "Scale":
		label = "Enter image file path to scale"
	}

	promptPath := promptui.Prompt{
		Label: label,
	}

	input, err := promptPath.Run()
	if err != nil {
		log.Println(err)
	}

	return input, nil
}

func PromptOutputPath() (string, error) {

	promptOutpath := promptui.Prompt{
		Label: "Enter the location to store the generated files",
	}
	outputPath, err := promptOutpath.Run()

	if err != nil {
		log.Println(err)
	}

	_, err = os.Stat(outputPath)
	if os.IsNotExist(err) {
		return "", err
	} else if err != nil {
		return "", err
	}

	return outputPath, nil

}
