package goextractorbarcode

import "image"

type extractor struct {
}

func NewExtractor() *extractor {
	return &extractor{}
}

func (e *extractor) Image(img image.Image) *extractorImage {
	return NewExtractorImage(img)
}

func (e *extractor) Images(imgs []image.Image) []*extractorImage {
	res := make([]*extractorImage, len(imgs))
	for i, img := range imgs {
		res[i] = NewExtractorImage(img)
	}
	return res
}

func (e *extractor) Pdf(filePath string) (*extractorPdfReady, error) {
	if ext, err := NewExtractorPdf(filePath).Read(); err != nil {
		return nil, err
	} else {
		return ext, nil
	}
}
