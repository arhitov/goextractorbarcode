package calculates

import (
	"github.com/arhitov/goextractorbarcode/helpers"
	"image"
)

func Density(binary *image.Gray, rect image.Rectangle) float64 {
	blackPixels := 0
	totalPixels := rect.Dx() * rect.Dy()

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			if helpers.IsBlack(binary, x, y) {
				blackPixels++
			}
		}
	}

	return float64(blackPixels) / float64(totalPixels)
}
