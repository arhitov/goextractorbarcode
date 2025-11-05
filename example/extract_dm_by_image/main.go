package main

import (
	"fmt"
	"github.com/arhitov/goextractorbarcode"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <image-path>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: file does not exist or cannot be opened: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to decode image: %v\n", err)
		os.Exit(1)
	}

	extractorImage := goextractorbarcode.NewExtractor().Image(img)
	res, err := extractorImage.ExtractDataMatrix()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to extract data matrix: %v\n", err)
		os.Exit(1)
	}

	for _, r := range res {
		fmt.Printf("Format: %s\n", r.Format())
		fmt.Printf("Text: %s\n", r.Text())
	}
}
