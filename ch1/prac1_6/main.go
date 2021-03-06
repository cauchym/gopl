package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.RGBA{0x00, 0x00, 0x00, 1}, color.RGBA{0xff, 0xff, 0x00, 1}, color.RGBA{0x00, 0xff, 0xff, 1}, color.RGBA{0xff, 0x00, 0x00, 1}}

const (
	blackIndex = 0 // パレットの最初の色
	greenIndex = 1 // パレットの次の色
)

//[TODO]palette使わず、paramsに応じて動的に変わるみたいなのはできなかった
// func gradateColor(tim int) color.Color{
//	return color.Color{color.RGBA{uint8(tim%32), uint8(tim%32), uint8(tim%32), 1}}
//}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 発振器xが完了する周回の回数
		res     = 0.001 // 回転の分解能
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8((i / 16) % 4 + 1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
