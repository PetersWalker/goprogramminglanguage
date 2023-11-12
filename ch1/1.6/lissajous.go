// Exercis 1.6: Mo dif y the Lissajous program to pro duce images in multiple color s by adding
// more values to palette and then displ aying them by chang ing the third argument of Set-
// ColorIndex in som e interesting way.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

const backgroundIndex = 1 // next color in palette
var palletIndex = uint8(0)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	var continuousPalette [][]color.Color
	// generate continutes palette
	for i := 0 ; i<=255; i++ {		
		// randColor:=rand.Int()
		continuousPalette = append(continuousPalette, []color.Color{color.RGBA{uint8(i*2), uint8(i) , uint8(i*4), 255}, color.White})
	}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		palette := continuousPalette[i]

		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func nextColorIndex(colorIndex uint8) uint8 {
	if (colorIndex == 254){
		return 0
	}
	
	return colorIndex + 1
}
