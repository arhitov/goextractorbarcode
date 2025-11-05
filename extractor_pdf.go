package goextractorbarcode

import (
	"fmt"
	"github.com/gen2brain/go-fitz"
	"image"
)

type extractorPdf struct {
	filePath string
	// Порог бинаризации
	binarizationThreshold uint8
	// Минимальная плотность
	minDensity float64
}

type extractorPdfReady struct {
	doc *fitz.Document
}

func NewExtractorPdf(filePath string) *extractorPdf {
	return &extractorPdf{
		filePath:   filePath,
		minDensity: 0.5,
	}
}

func (e *extractorPdf) Read() (*extractorPdfReady, error) {
	doc, err := fitz.New(e.filePath)
	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия PDF: %v\n", err)
	}
	return &extractorPdfReady{doc: doc}, nil
}

func (e *extractorPdfReady) Close() {
	_ = e.doc.Close()
}

func (e *extractorPdfReady) NumPage() int {
	return e.doc.NumPage()
}

// ToImages Преобразуем каждую страницу в изображение
func (e *extractorPdfReady) ToImages() []image.Image {
	// Получаем количество страниц
	pageCount := e.doc.NumPage()

	// Слайс для хранения изображений
	var images []image.Image = make([]image.Image, pageCount)

	for i := 0; i < pageCount; i++ {
		// Извлекаем изображение страницы
		img, err := e.doc.Image(i)
		if err != nil {
			//fmt.Printf("Ошибка преобразования страницы %d: %v\n", i+1, err)
			continue
		}

		images[i] = img
	}

	return images
}

func (e *extractorPdfReady) ExtractDataMatrix() ([]*ResultCode, error) {
	images := e.ToImages()
	results := make([]*ResultCode, 0, len(images))

	for _, img := range images {
		res, err := NewExtractorImage(img).ExtractDataMatrix()
		if err == nil {
			results = append(results, res...)
		}
	}

	return results, nil
}
