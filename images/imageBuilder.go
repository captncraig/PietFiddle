package images

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"golang.org/x/image/bmp"
	"hash/crc32"
	"image"
	"image/png"
	"io"
)

func BuildImage(w, h int, data string, codelSize int, writer io.Writer) {
	img := image.NewPaletted(image.Rect(0, 0, w*codelSize, h*codelSize), palette)
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

func BuildImageRaw(w, h int, data string, codelSize int, writer io.Writer) {
	hdr := []byte{137, 80, 78, 71, 13, 10, 26, 10}
	writer.Write(hdr)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, int32(w*codelSize))
	binary.Write(buf, binary.BigEndian, int32(h*codelSize))
	buf.WriteByte(8)
	buf.WriteByte(3)
	buf.WriteByte(0)
	buf.WriteByte(0)
	buf.WriteByte(0)
	writeChunk("IHDR", buf.Bytes(), writer)
	writeChunk("PLTE", []byte{255, 192, 192, 255, 0, 0, 192, 0, 0, 255, 255, 192, 255, 255, 0, 192, 192, 0, 192, 255, 192, 0, 255, 0, 0, 192, 0, 192, 255, 255, 0, 255, 255, 0, 192, 192, 192, 192, 255, 0, 0, 255, 0, 0, 192, 255, 192, 255, 255, 0, 255, 192, 0, 192, 255, 255, 255, 0, 0, 0}, writer)
	buf = new(bytes.Buffer)
	z, _ := zlib.NewWriterLevel(buf, zlib.BestCompression)

	row := make([]byte, w*codelSize+1)
	noopRow := make([]byte, w*codelSize+1)
	noopRow[0] = 2
	row[0] = 1
	for y := 0; y < h; y++ {
		lastColor := uint8(0)
		for x := 0; x < w; x++ {
			color := data[y*w+x]
			idx := color - 'A'
			row[(x*codelSize)+1] = idx - lastColor
			lastColor = idx

		}
		z.Write(row)
		for i := 1; i < codelSize; i++ {
			z.Write(noopRow)
		}
	}
	z.Flush()
	writeChunk("IDAT", buf.Bytes(), writer)
	writeChunk("IEND", []byte{}, writer)
}

func writeChunk(typ string, data []byte, w io.Writer) {
	binary.Write(w, binary.BigEndian, int32(len(data)))
	w.Write([]byte(typ))
	w.Write(data)
	hash := crc32.NewIEEE()
	hash.Write([]byte(typ))
	hash.Write(data)
	binary.Write(w, binary.BigEndian, hash.Sum32())
}

func BuildImageBmpRaw(w, h int, data string, codelSize int, writer io.Writer) {
	//using rle: size of data potion = rows * (cols * 2) + 2?
	rowSize := w * codelSize * 1
	if rowSize%4 != 0 {
		rowSize += 4 - (rowSize % 4)
	}
	imgSize := h * codelSize * rowSize
	palletSize := 20 * 4
	fileSize := imgSize + 54 + palletSize
	writer.Write([]byte("BM"))
	binary.Write(writer, binary.LittleEndian, int32(fileSize))
	//reserved
	writer.Write([]byte{0, 0, 0, 0})
	//offset
	binary.Write(writer, binary.LittleEndian, int32(54+palletSize)) //+pallet size
	//size of header
	writer.Write([]byte{0x28, 0, 0, 0})
	//width
	binary.Write(writer, binary.LittleEndian, int32(w*codelSize))
	//height
	binary.Write(writer, binary.LittleEndian, int32(h*codelSize))
	//color planes / bpp/ compression
	writer.Write([]byte{1, 0, 24, 0, 0, 0, 0, 0})
	//bitmap data size
	binary.Write(writer, binary.LittleEndian, int32(imgSize))
	//resolution / resolution / pallette size / important colors
	writer.Write([]byte{0x13, 0xb, 0, 0, 0x13, 0xb, 0, 0 /*pallete size*/, 20, 0, 0, 0, 0, 0, 0, 0})
	//pallette
	writer.Write([]byte{255, 192, 192, 0, 255, 0, 0, 0, 192, 0, 0, 0, 255, 255, 192, 0, 255, 255, 0, 0, 192, 192, 0, 0, 192, 255, 192, 0, 0, 255, 0, 0, 0, 192, 0, 0, 192, 255, 255, 0, 0, 255, 255, 0, 0, 192, 192, 0, 192, 192, 255, 0, 0, 0, 255, 0, 0, 0, 192, 0, 255, 192, 255, 0, 255, 0, 255, 0, 192, 0, 192, 0, 255, 255, 255, 0, 0, 0, 0, 0})
	row := make([]byte, rowSize)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			color := data[i*w+j]
			idx := color - 'A'
			c := colors[idx]
			fmt.Println(c.RGBA())
		}
		writer.Write(row)
	}
}

func BuildImageBmp(w, h int, data string, codelSize int, writer io.Writer) {
	img := image.NewPaletted(image.Rect(0, 0, w*codelSize, h*codelSize), palette)
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
	bmp.Encode(writer, img)
}
