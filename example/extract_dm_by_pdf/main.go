package main

import (
	"fmt"
	"github.com/arhitov/goextractorbarcode"
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

	extractor, err := goextractorbarcode.NewExtractor().Pdf(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to create extractor: %v\n", err)
		os.Exit(1)
	}
	defer extractor.Close()

	// Получаем количество страниц
	pageCount := extractor.NumPage()
	fmt.Printf("Количество страниц: %d\n", pageCount)

	res, err := extractor.ExtractDataMatrix()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to extract data matrix: %v\n", err)
		os.Exit(1)
	}

	for _, r := range res {
		fmt.Printf("Format: %s\n", r.Format())
		fmt.Printf("Text: %s\n", r.Text())
	}
}
