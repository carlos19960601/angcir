package main

import (
	"math"
	"os"

	svg "github.com/ajstarks/svgo/float"
)

func main() {
	width := 1010.0
	height := 1010.0
	f, _ := os.OpenFile("x.svg", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 5, "fill:green")
	// canvas.Circle(width/2, height/2, 500, "fill:none;stroke:green;stroke-width:2")

	// for index := 0.0; index < 200; index++ {
	// 	canvas.Line(width/2, height/2, width/2+500*math.Cos(index*math.Pi/100), height/2+500*math.Sin(index*math.Pi/100), "stroke:green;stroke-width:2")
	// }

	for index := 0.0; index < 200; index++ {
		canvas.Circle(width/2+500*math.Cos(index*math.Pi/100), height/2+500*math.Sin(index*math.Pi/100), 2, "fill:green")
	}

	canvas.End()
}
