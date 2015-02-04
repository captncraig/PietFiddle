package images

import (
	"errors"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

var palette color.Palette = buildPalette()

var colors = make([]color.Color, 20)

func buildPalette() color.Palette {
	p := colors
	p[0], _ = colorful.Hex("#FFC0C0")
	p[1], _ = colorful.Hex("#FF0000")
	p[2], _ = colorful.Hex("#C00000")
	p[3], _ = colorful.Hex("#FFFFC0")
	p[4], _ = colorful.Hex("#FFFF00")
	p[5], _ = colorful.Hex("#C0C000")
	p[6], _ = colorful.Hex("#C0FFC0")
	p[7], _ = colorful.Hex("#00FF00")
	p[8], _ = colorful.Hex("#00C000")
	p[9], _ = colorful.Hex("#C0FFFF")
	p[10], _ = colorful.Hex("#00FFFF")
	p[11], _ = colorful.Hex("#00C0C0")
	p[12], _ = colorful.Hex("#C0C0FF")
	p[13], _ = colorful.Hex("#0000FF")
	p[14], _ = colorful.Hex("#0000C0")
	p[15], _ = colorful.Hex("#FFC0FF")
	p[16], _ = colorful.Hex("#FF00FF")
	p[17], _ = colorful.Hex("#C000C0")
	p[18], _ = colorful.Hex("#FFFFFF")
	p[19], _ = colorful.Hex("#000000")
	return color.Palette(p)
}

func LoadImage(input io.Reader, codelSize int) (width int, height int, data string, err error) {
	i, _, err := image.Decode(input)
	if err != nil {
		return 0, 0, "", err
	}

	width = i.Bounds().Max.X
	height = i.Bounds().Max.Y

	if width%codelSize != 0 || height%codelSize != 0 {
		return 0, 0, "", errors.New("codel size mismatch")
	}
	width /= codelSize
	height /= codelSize
	fmt.Println(width, height)
	b := make([]byte, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := i.At(x*codelSize, y*codelSize)
			idx := palette.Index(c)
			//TODO: Check found color against actual color.
			// Apply rule on mismatch. Take closest? White? Black?
			b[y*width+x] = byte(idx + 'A')
		}
	}
	fmt.Println(string(b))
	return width, height, string(b), nil
}
