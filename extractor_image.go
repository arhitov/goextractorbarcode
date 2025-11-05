package goextractorbarcode

import (
	"fmt"
	"github.com/arhitov/goextractorbarcode/calculates"
	"github.com/arhitov/goextractorbarcode/convs"
	"github.com/arhitov/goextractorbarcode/finders/finderDataMatrix"
	"github.com/arhitov/goextractorbarcode/helpers"
	"github.com/arhitov/goextractorbarcode/types"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/datamatrix"
	"image"
)

type extractorImage struct {
	img image.Image
	// Порог бинаризации
	binarizationThreshold uint8
	// Минимальная плотность
	minDensity float64
}

func NewExtractorImage(img image.Image) *extractorImage {
	return &extractorImage{
		img:        img,
		minDensity: 0.5,
	}
}

func (e *extractorImage) ExtractDataMatrix() ([]*ResultCode, error) {
	results := make([]*ResultCode, 0)

	// Рассчитываем порог бинаризации, если он не указан
	if e.binarizationThreshold == 0 {
		e.binarizationThreshold = calculates.AdaptiveThreshold(e.img)
		//e.binarizationThreshold = calculates.OtsuThreshold(e.img)
	}

	binary := convs.ToBinaryImage(e.img, e.binarizationThreshold)

	regions := finderDataMatrix.FindDataMatrixRegions(binary, e.minDensity)

	// Создаем декодер DataMatrix
	reader := datamatrix.NewDataMatrixReader()
	hints := map[gozxing.DecodeHintType]interface{}{
		gozxing.DecodeHintType_TRY_HARDER: true,
		//gozxing.DecodeHintType_PURE_BARCODE: false,
	}

	for _, region := range regions {
		// Вырезаем регион из изображения
		cropped := helpers.CropImage(binary, region)

		// Пробуем декодировать
		bmp, err := gozxing.NewBinaryBitmapFromImage(cropped)
		if err != nil {
			fmt.Printf("  Ошибка создания битмапа: %v\n", err)
			continue
		}

		result, err := reader.Decode(bmp, hints)
		if result == nil {
			continue
		}

		if result.GetBarcodeFormat() != gozxing.BarcodeFormat_DATA_MATRIX {
			continue
		}

		results = append(results, NewDataCode(
			types.CodeFormatDataMatrix,
			result.GetText(),
		))
	}

	return results, nil
}
