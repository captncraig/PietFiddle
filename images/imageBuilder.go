package images

import (
	"github.com/spakin/netpbm"
	"image"
	"image/gif"
	"image/png"
	"io"
)

func BuildImage(w, h int, data string, codelSize, rotation int, writer io.Writer) {
	img := makeImage(w, h, data, codelSize, rotation)
	png.Encode(writer, img)
}

func BuildGif(w, h int, data string, codelSize int, writer io.Writer) {
	g := gif.GIF{}
	g.Image = make([]*image.Paletted, 18)
	g.Delay = make([]int, 18)
	g.LoopCount = -1
	for i := 0; i < 18; i++ {
		g.Image[i] = makeImage(w, h, data, codelSize, i)
		g.Delay[i] = 12
	}
	gif.EncodeAll(writer, &g)
}

func BuildPpm(w, h int, data string, codelSize, rotation int, writer io.Writer) {
	img := makeImage(w, h, data, codelSize, rotation)
	netpbm.Encode(writer, img, &netpbm.EncodeOptions{
		Format:   netpbm.PPM,
		MaxValue: 255,
		//Plain:    true,
	})
}

func makeImage(w, h int, data string, codelSize, rotation int) *image.Paletted {
	hueAdd := uint8(rotation / 3)
	ligAdd := uint8(rotation % 3)
	img := image.NewPaletted(image.Rect(0, 0, w*codelSize, h*codelSize), palette)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			color := data[y*w+x]
			idx := color - 'A'
			idx = rotateColor(idx, hueAdd, ligAdd)
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
	return img
}

func rotateColor(idx, hueAdd, ligAdd uint8) uint8 {
	if idx > 17 {
		return idx
	}
	hue := ((idx / 3) + hueAdd) % 6
	lig := ((idx % 3) + ligAdd) % 3
	return (hue * 3) + lig
}
