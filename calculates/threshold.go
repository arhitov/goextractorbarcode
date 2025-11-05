package calculates

import (
	"image"
	"image/color"
)

// AdaptiveThreshold Рассчитывает порог для бинаризации адаптивным методом
func AdaptiveThreshold(img image.Image) uint8 {
	bounds := img.Bounds()
	var sum int
	count := bounds.Dx() * bounds.Dy()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			sum += int(gray.Y)
		}
	}

	// Используем среднее значение как порог
	return uint8(sum / count)
}

// OtsuThreshold Рассчитывает порог для бинаризации методом Оцу
func OtsuThreshold(img image.Image) uint8 {
	bounds := img.Bounds()

	// Строим гистограмму
	histogram := [256]int{}
	totalPixels := bounds.Dx() * bounds.Dy()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			histogram[gray.Y]++
		}
	}

	// Вычисляем оптимальный порог методом Оцу
	sum := 0
	for i := 0; i < 256; i++ {
		sum += i * histogram[i]
	}

	sumB := 0
	wB := 0
	wF := 0
	maxVariance := 0.0
	threshold := 0

	for i := 0; i < 256; i++ {
		wB += histogram[i]
		if wB == 0 {
			continue
		}

		wF = totalPixels - wB
		if wF == 0 {
			break
		}

		sumB += i * histogram[i]

		mB := float64(sumB) / float64(wB)
		mF := float64(sum-sumB) / float64(wF)

		// Вычисляем дисперсию между классами
		variance := float64(wB) * float64(wF) * (mB - mF) * (mB - mF)

		if variance > maxVariance {
			maxVariance = variance
			threshold = i
		}
	}

	return uint8(threshold)
}
