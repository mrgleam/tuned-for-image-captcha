package main

import (
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type ImageSet interface {
	Set(x, y int, c color.Color)
}

func main() {
	r, err := os.Open("example.png")
	check(err)
	defer r.Close()

	img, _, err := image.Decode(r)
	check(err)

	b := img.Bounds()

	imgSet := img.(ImageSet)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			oldPixel := img.At(x, y)

			r, g, b, a := oldPixel.RGBA()
			//fmt.Println(r, g, b, a)
			r = 65535 - r
			g = 65535 - g
			b = 65535 - b
			pixel := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			imgSet.Set(x, y, pixel)
		}
	}
	m := resize.Resize(uint(b.Size().X*5), uint(b.Size().Y*5), img, resize.Lanczos3)

	w, err := os.Create("test.png")
	check(err)
	defer w.Close()

	err = png.Encode(w, m)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
