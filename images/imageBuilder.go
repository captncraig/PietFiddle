package images

import (
	"fmt"
	"image"
	"image/png"
	"io"
)

func BuildImage(w, h int, data string, codelSize int, writer io.Writer) {
	img := image.NewPaletted(image.Rect(0, 0, w*codelSize, h*codelSize), palette)
	fmt.Println(w*h, len(data))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			color := data[y*w+x]
			idx := color - 'A'
			startX := x * codelSize
			startY := y * codelSize
			endX := x*codelSize + codelSize
			endY := y*codelSize + codelSize
			for xx := startX; xx <= endX; xx++ {
				for yy := startY; yy < endY; yy++ {
					img.SetColorIndex(xx, yy, idx)
				}
			}
		}
	}
	png.Encode(writer, img)
}
