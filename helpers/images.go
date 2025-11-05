package helpers

import (
	"image"
	"image/draw"
)

//func cropImage(img image.Image, rect image.Rectangle) image.Image {
//	// Создаем новое изображение с нужными размерами
//	cropped := image.NewRGBA(rect)
//	bounds := img.Bounds()
//
//	for y := rect.Min.Y; y < rect.Max.Y; y++ {
//		for x := rect.Min.X; x < rect.Max.X; x++ {
//			if x >= bounds.Min.X && x < bounds.Max.X && y >= bounds.Min.Y && y < bounds.Max.Y {
//				cropped.Set(x, y, img.At(x, y))
//			} else {
//				// Заполняем белым цветом, если выходим за границы
//				cropped.Set(x, y, color.White)
//			}
//		}
//	}
//
//	return cropped
//}

// CropImage Обновленная функция CropImage с проверкой границ
func CropImage(img image.Image, rect image.Rectangle) image.Image {
	//bounds := img.Bounds()
	//
	//// Убеждаемся, что регион не выходит за границы
	//cropRect := image.Rect(
	//	max(rect.Min.X, bounds.Min.X),
	//	max(rect.Min.Y, bounds.Min.Y),
	//	min(rect.Max.X, bounds.Max.X),
	//	min(rect.Max.Y, bounds.Max.Y),
	//)

	//fmt.Printf("CropImage rect %+v %+v\n", rect.Min, rect.Max)
	//fmt.Printf("CropImage cropRect %+v\n", cropRect)

	// Создаем новое изображение с нужными размерами
	croppedBounds := image.Rect(0, 0, rect.Max.X-rect.Min.X, rect.Max.Y-rect.Min.Y)
	cropped := image.NewRGBA(croppedBounds)
	//fmt.Printf("CropImage bounds %+v\n", croppedBounds)

	// Копируем оригинальное изображение
	draw.Draw(cropped, croppedBounds, img, rect.Min, draw.Src)

	//cropped := image.NewRGBA(cropRect)
	//for y := cropRect.Min.Y; y < cropRect.Max.Y; y++ {
	//	for x := cropRect.Min.X; x < cropRect.Max.X; x++ {
	//		cropped.Set(x-cropRect.Min.X, y-cropRect.Min.Y, img.At(x, y))
	//	}
	//}

	return cropped
}
