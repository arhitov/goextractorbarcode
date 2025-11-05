package finderDataMatrix

import (
	"github.com/arhitov/goextractorbarcode/helpers"
	"image"
)

func findConnectedComponents(binary *image.Gray) []image.Rectangle {
	bounds := binary.Bounds()
	visited := make([][]bool, bounds.Dy())
	for i := range visited {
		visited[i] = make([]bool, bounds.Dx())
	}

	var blobs []image.Rectangle

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if !visited[y][x] && helpers.IsBlack(binary, x, y) {
				// Нашли новый blob, начинаем обход
				blob := floodFill(binary, visited, x, y)
				if blob.Dx() > 5 && blob.Dy() > 5 { // Игнорируем слишком маленькие blob'ы
					blobs = append(blobs, blob)
				}
			}
		}
	}

	return blobs
}

//// Улучшенная функция поиска связанных компонентов
//func findConnectedComponents(binary *image.Gray) []image.Rectangle {
//	bounds := binary.Bounds()
//	visited := make([][]bool, bounds.Dy())
//	for i := range visited {
//		visited[i] = make([]bool, bounds.Dx())
//	}
//
//	var blobs []image.Rectangle
//
//	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
//		for x := bounds.Min.X; x < bounds.Max.X; x++ {
//			if !visited[y][x] && isBlack(binary, x, y) {
//				// Нашли новый blob, начинаем обход
//				blob := floodFill(binary, visited, x, y)
//
//				// Фильтруем слишком маленькие и слишком большие blob'ы
//				width := blob.Dx()
//				height := blob.Dy()
//
//				if width >= 20 && height >= 20 &&
//					width <= bounds.Dx()-10 && height <= bounds.Dy()-10 {
//					blobs = append(blobs, blob)
//				}
//			}
//		}
//	}
//
//	return blobs
//}

func floodFill(img *image.Gray, visited [][]bool, startX, startY int) image.Rectangle {
	bounds := img.Bounds()
	minX, minY := startX, startY
	maxX, maxY := startX, startY

	stack := [][2]int{{startX, startY}}
	visited[startY][startX] = true

	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // Диагонали для лучшего связывания
	}

	for len(stack) > 0 {
		point := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		x, y := point[0], point[1]

		// Обновляем границы blob'а
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}

		// Проверяем соседей
		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= bounds.Min.X && nx < bounds.Max.X &&
				ny >= bounds.Min.Y && ny < bounds.Max.Y &&
				!visited[ny][nx] && helpers.IsBlack(img, nx, ny) {
				visited[ny][nx] = true
				stack = append(stack, [2]int{nx, ny})
			}
		}
	}

	return image.Rect(minX, minY, maxX+1, maxY+1)
}
