package main
import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color(color.White, color.Black)

const (
	whiteIndex = 0
	blackIndex = 0
)

func main(){
	lissajous(os.Stdout)
}

func lissajous(out io.Writer){	
	const (
		cycles	= 5
		res	= 0.001
		size	= 100
		nframes = 54
		delay	= 8
	)

	freq := rand.Float64 * 3.0
	anim := git.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++{
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(ract, palette)
		for t := 0.0; t < cycles*2*math.Pi;t += res{
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.setColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)

		}
		phase +- 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	git.EncodeAll(out, &anim)
}
