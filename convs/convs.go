package convs

import (
	"image"
	"image/color"
)

// ToBinaryImage Конвертируем в бинарное изображение
func ToBinaryImage(img image.Image, threshold uint8) *image.Gray {
	bounds := img.Bounds()
	binary := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			if gray.Y < threshold {
				binary.SetGray(x, y, color.Gray{Y: 0}) // Черный
			} else {
				binary.SetGray(x, y, color.Gray{Y: 255}) // Белый
			}
		}
	}

	return binary
}
