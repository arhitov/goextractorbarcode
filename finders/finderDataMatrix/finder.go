package finderDataMatrix

import (
	"github.com/arhitov/goextractorbarcode/calculates"
	"image"
	"math"
)

// FindDataMatrixRegions находит регионы DataMatrix в бинарном изображении
// и возвращает их в виде прямоугольников
// minDensity - минимальная плотность региона (например, 0.25 - 25% минимальной плотности)
func FindDataMatrixRegions(binary *image.Gray, minDensity float64) []image.Rectangle {
	// Находим связанные компоненты (blob detection)
	blobs := findConnectedComponents(binary)

	// Фильтруем по размеру, форме и плотности
	var regions []image.Rectangle
	for _, blob := range blobs {
		density := calculates.Density(binary, blob)

		// Используем minDensity для фильтрации
		if isValidDataMatrixRegion(blob, binary, density, minDensity) {
			// Добавляем padding вокруг региона
			padded := expandRectangle(blob, 1.0, binary.Bounds())
			regions = append(regions, padded)
		}
	}

	return regions
}

//func FindDataMatrixRegions(img image.Image, minDensity float64) []image.Rectangle {
//	// Конвертируем в бинарное изображение
//	binary := ToBinaryImage(img)
//
//	// Находим связанные компоненты (blob detection)
//	blobs := findConnectedComponents(binary)
//
//	// Фильтруем по размеру, форме и плотности
//	var regions []image.Rectangle
//	for _, blob := range blobs {
//		density := CalculateDensity(binary, blob)
//
//		// Используем minDensity для фильтрации
//		if isValidDataMatrixRegion(blob, binary, density, minDensity) {
//			// Добавляем padding вокруг региона, но не выходим за границы изображения
//			padded := expandRectangle(blob, 10, img.Bounds())
//			regions = append(regions, padded)
//		}
//	}
//
//	return regions
//}

func isValidDataMatrixRegion(rect image.Rectangle, binary *image.Gray, density float64, minDensity float64) bool {
	width := rect.Dx()
	height := rect.Dy()

	// DataMatrix обычно квадратные или близкие к квадрату
	aspectRatio := float64(width) / float64(height)
	if aspectRatio < 0.7 || aspectRatio > 1.3 {
		return false
	}

	// Проверяем размер (DataMatrix не может быть слишком маленьким)
	if width < 20 || height < 20 || width > binary.Bounds().Dx() || height > binary.Bounds().Dy() {
		return false
	}

	// Используем minDensity для проверки минимальной плотности
	// и устанавливаем максимальную плотность на разумном уровне
	if density < minDensity || density > 0.8 {
		return false
	}

	// Дополнительная проверка: DataMatrix должен иметь равномерное распределение
	if !hasUniformDistribution(binary, rect) {
		return false
	}

	return true
}

// Новая функция для проверки равномерного распределения пикселей
func hasUniformDistribution(binary *image.Gray, rect image.Rectangle) bool {
	width := rect.Dx()
	height := rect.Dy()

	// Разбиваем регион на 4 части и проверяем плотность в каждой
	regions := []image.Rectangle{
		image.Rect(rect.Min.X, rect.Min.Y, rect.Min.X+width/2, rect.Min.Y+height/2),
		image.Rect(rect.Min.X+width/2, rect.Min.Y, rect.Max.X, rect.Min.Y+height/2),
		image.Rect(rect.Min.X, rect.Min.Y+height/2, rect.Min.X+width/2, rect.Max.Y),
		image.Rect(rect.Min.X+width/2, rect.Min.Y+height/2, rect.Max.X, rect.Max.Y),
	}

	var densities []float64
	for _, region := range regions {
		density := calculates.Density(binary, region)
		densities = append(densities, density)
	}

	// Проверяем, что плотности в разных частях не слишком отличаются
	maxDiff := 0.0
	for i := 0; i < len(densities); i++ {
		for j := i + 1; j < len(densities); j++ {
			diff := math.Abs(densities[i] - densities[j])
			if diff > maxDiff {
				maxDiff = diff
			}
		}
	}

	// Если разница плотностей между регионами слишком большая - это не DataMatrix
	return maxDiff < 0.3
}

func expandRectangle(rect image.Rectangle, padding float64, imgBounds image.Rectangle) image.Rectangle {
	paddingX := max(0, int(float64(imgBounds.Max.X-imgBounds.Min.X)/100*padding))
	paddingY := max(0, int(float64(imgBounds.Max.Y-imgBounds.Min.Y)/100*padding))

	return image.Rect(
		max(rect.Min.X-paddingX, imgBounds.Min.X),
		max(rect.Min.Y-paddingY, imgBounds.Min.Y),
		min(rect.Max.X+paddingX, imgBounds.Max.X),
		min(rect.Max.Y+paddingY, imgBounds.Max.Y),
	)
}
