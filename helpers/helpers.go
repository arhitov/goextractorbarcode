package helpers

import "image"

func IsBlack(img *image.Gray, x, y int) bool {
	return img.GrayAt(x, y).Y == 0
}
