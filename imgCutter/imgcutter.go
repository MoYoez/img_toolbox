package imgCutter

import (
	"github.com/nfnt/resize"
	"image"
)

// use resize to do.
func CropImage(img image.Image, targetWidth, targetHeight int) (destimg image.Image) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	if targetWidth > width || targetHeight > height {
		var scale float64
		if float64(targetWidth)/float64(width) > float64(targetHeight)/float64(height) {
			scale = float64(targetWidth) / float64(width)
		} else {
			scale = float64(targetHeight) / float64(height)
		}
		img = resize.Resize(uint(scale*float64(width)), uint(scale*float64(height)), img, resize.Lanczos3)
		bounds = img.Bounds()
		width = bounds.Dx()
		height = bounds.Dy()
	}
	x := (width - targetWidth) / 2
	y := (height - targetHeight) / 2
	rect := image.Rect(x, y, x+targetWidth, y+targetHeight)
	croppedImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(rect)
	return croppedImg
}
