package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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

func PromptUserInputConvertion(convertionType string) (string, error) {
	var label string
	switch convertionType {
	case "JPG to PNG":
		label = "Enter image file path(s) to convert to PNG (space-separated):"
	case "PNG to JPG":
		label = "Enter image file path(s) to convert to JPG (space-separated):"
	}

	fmt.Println(label)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func PromptUserInputScaling() (string, uint, uint, error) {

	fmt.Printf("Enter image file path to scale: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", 0, 0, err
	}

	fmt.Printf("Enter the width of the image to scale: ")
	width, err := reader.ReadString('\n')
	if err != nil {
		return "", 0, 0, err
	}

	fmt.Printf("Enter the height of the image to scale: ")
	height, err := reader.ReadString('\n')
	if err != nil {
		return "", 0, 0, err
	}

	uint64Width, err := strconv.ParseUint(strings.TrimSpace(width), 10, 64)
	if err != nil {
		return "", 0, 0, err
	}

	uint64Height, err := strconv.ParseUint(strings.TrimSpace(height), 10, 64)
	if err != nil {
		return "", 0, 0, err
	}

	return strings.TrimSpace(input), uint(uint64Width), uint(uint64Height), nil
}

func PromptOutputPath() (string, error) {
	fmt.Println("Enter the location to store the generated files:")
	reader := bufio.NewReader(os.Stdin)
	outputPath, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	outputPath = strings.TrimSpace(outputPath)
	_, err = os.Stat(outputPath)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("directory does not exist: %s", outputPath)
	} else if err != nil {
		return "", err
	}

	return outputPath, nil
}
