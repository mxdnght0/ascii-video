package ascii

import (
	"image"
	"image/color"
	"github.com/nfnt/resize"
)

func ImageToASCII(img image.Image, width int) []string {
	grayRamp := " .'`^,:;OD8MW&%#@"
	bounds := img.Bounds()
	ratio := float64(bounds.Dy()) / float64(bounds.Dx()) * 0.5
	newHeight := int(float64(width) * ratio)

	resized := resize.Resize(uint(width), uint(newHeight), img, resize.Lanczos3)

	var result []string

	for y := 0; y < newHeight; y++ {
		line := ""
		for x := 0; x < width; x++ {
			c := color.GrayModel.Convert(resized.At(x, y)).(color.Gray)
			brightness := float64(c.Y) / 255.0
			index := int(brightness * float64(len(grayRamp)-1))
			line += string(grayRamp[index])
		}
		result = append(result, line)
	}

	return result
}
