package main
import(
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const(
	whiteIndex = 0
	blackIndex = 1
)

func main(){
	lissajous(os.Stdout)
}

func lissajous(out io.writer){
	const{
		cycle = 5
		res  = 0.001
		size = 100
		nframes = 64
		delay = 8
	}
	
	freq := rand.Float64() * 3.0
	anim := gif.Gif{LoopCount: nframes}
	phase := 0.0
	for i:=0; i < nframes;i++{
		rect := image.Rect(0,0,2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t<cycle*2*math.Pi; t += res{
			x := math.Sin(t)
		}
		
	}
}