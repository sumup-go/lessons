package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/Pitrified/go-turtle"
)

func main() {
	rand.Seed(time.Now().Unix())

	w := turtle.NewWorld(500, 500)
	defer w.Close()

	td := turtle.NewTurtleDraw(w)

	draw(td)

	// save the current image
	err := w.SaveImage("example.png")
	if err != nil {
		fmt.Println("Could not save the image: ", err)
	}
}

func draw(td *turtle.TurtleDraw) {
	td.SetPos(250, 250)
	td.SetHeading(turtle.East)
	td.PenDown()
	td.SetSize(2)

	for i := 0; i < 50; i++ {
		td.SetColor(color.RGBA{randColor(), randColor(), randColor(), 255})
		td.Forward(float64(20 * i))
		td.Right(90)
	}
}

func randColor() uint8 {
	return uint8(255 * rand.Float64())
}
