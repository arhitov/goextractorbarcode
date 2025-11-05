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
