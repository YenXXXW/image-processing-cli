package input

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"

	"github.com/google/shlex"
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

func PromptUserInputConvertion(convertionType string) ([]string, error) {
	var label string
	switch convertionType {
	case "JPG to PNG":
		label = "Enter file paths (one per line). Press Enter on an empty line to finish:"
	case "PNG to JPG":
		label = "Enter file paths (one per line). Press Enter on an empty line to finish:"
	}

	fmt.Println(label)
	scanner := bufio.NewScanner(os.Stdin)
	var paths []string

	for {
		fmt.Printf("file path: ")
		scanner.Scan()

		raw := strings.TrimSpace(scanner.Text())

		if raw == "" {
			break
		}
		espcatedPaths, err := shlex.Split(raw)
		if err != nil {
			fmt.Printf("Error escaping the shell path %v", err)
			return []string{}, nil
		}
		filepath := espcatedPaths[0]

		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			fmt.Println("❌ File does not exist. Try again.")
			continue
		}

		base := path.Base(filepath)
		ext := path.Ext(base)

		if convertionType == "JPG to PNG" {
			if !slices.Contains([]string{".jpg", ".jpeg"}, ext) {
				fmt.Println("provided file is not a jpg file")
				continue
			}
		} else {

			if !slices.Contains([]string{".png", ".PNG"}, ext) {
				fmt.Println("provided file is not a jpg file")
				continue
			}
		}

		paths = append(paths, filepath)
	}

	return paths, nil
}

func PromptUserInputScaling() (string, uint, uint, error) {

	fmt.Printf("Enter image file path to scale: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", 0, 0, err
	}

	var uint64Width uint64
	var uint64Height uint64

	var min uint64
	var max uint64
	min = 1
	max = 10000

	fmt.Printf("Enter the width of the image to scale: ")
	uint64Width = promptUintInRange("Enter the width of the image", min, max)
	uint64Height = promptUintInRange("Enter the height of the image", min, max)

	return strings.TrimSpace(input), uint(uint64Width), uint(uint64Height), nil
}

func PromptOutputPath() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	var outputPath string
	for {

		fmt.Print("Enter the location to store the generated files: ")
		rawString, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("error reading the output path: %v", err)
		}

		rawString = strings.TrimSpace(rawString)

		espcatedPaths, err := shlex.Split(rawString)
		if err != nil {
			fmt.Printf("Error escaping the shell path %v\n", err)
			return "", nil
		}
		outputPath = espcatedPaths[0]
		info, err := os.Stat(outputPath)
		if os.IsNotExist(err) {
			fmt.Printf("directory does not exist\n")
			continue
		} else if err != nil {
			continue
		}

		if !info.IsDir() {
			fmt.Println("Given path is not a dir")
			continue
		}

		break
	}

	return outputPath, nil
}

func promptUintInRange(label string, min, max uint64) uint64 {
	for {
		fmt.Printf("Enter %s (%d–%d): ", label, min, max)
		var input string
		fmt.Scanln(&input)

		value, err := strconv.ParseUint(strings.TrimSpace(input), 10, 64)
		if err != nil {
			fmt.Println("❌ Invalid number. Please enter a valid positive integer.")
			continue
		}

		if value < min || value > max {
			fmt.Printf("❌ Value out of range. Must be between %d and %d.\n", min, max)
			continue
		}

		return value
	}
}
