# Image Processing CLI

A fast and simple command-line tool for image conversion and resizing, built with Go.

## Features
- **Convert images between JPEG and PNG formats** using Go's standard `image/jpeg` and `image/png` packages.
- **Resize images** with high quality using [`github.com/nfnt/resize`](https://github.com/nfnt/resize).
- **Concurrent processing**: Convert multiple images in parallel for faster batch operations.
- **Cross-platform binaries**: Built with [goreleaser](https://goreleaser.com/) for Windows, macOS, and Linux.
- **Interactive CLI**: User-friendly prompts for file paths, output directories, and scaling dimensions.

## Installation

### Download Pre-built Binaries
You can download the latest release for your operating system from the [GitHub Releases page](https://github.com/yenxxxw/image-processing-cli/releases).

### Build from Source
1. Clone the repository:
   ```sh
   git clone https://github.com/yenxxxw/image-processing-cli.git
   cd image-processing-cli
   ```
2. Build the binary:
   ```sh
   go build -o ./bin/image-processing-cli
   ```
3. The binary will be available in the `bin` folder. Run it with:
   ```sh
   ./bin/image-processing-cli
   ```

## Usage
- **Convert images:**
  - Enter the file path(s) to convert from JPEG to PNG or PNG to JPEG.
  - The tool supports batch conversion with concurrent processing.
- **Scale images:**
  - Enter the image path and specify the desired width and height.
- **Output:**
  - Specify the output directory for the converted or scaled images.

## Dependencies
- [Go standard library](https://golang.org/pkg/): `image/jpeg`, `image/png`
- [github.com/nfnt/resize](https://github.com/nfnt/resize) for image resizing
- [goreleaser](https://goreleaser.com/) for building cross-platform binaries

## License
MIT 