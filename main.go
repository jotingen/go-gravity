package main

import (
	"fmt"
	"math/rand"
	"time"
)

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

import (
	"github.com/jotingen/go-gravity/gravity"
)

const (
	windowWidth  = 960
	windowHeight = 540
)

var (
	u gravity.Universe
)

func main() {

	fmt.Println("go-gravity")
	u = gravity.Universe{}
	u.Bodies = append(u.Bodies, gravity.Body{
		XPos: 0,
		YPos: 0,
		ZPos: 0,
		XVel: 0,
		YVel: 0,
		ZVel: 0,
		Mass: 20,
	})
	for i := 0; i < 5; i++ {
		u.Bodies = append(u.Bodies, gravity.Body{
			XPos: (rand.Float64()*2 - 1) * 10,
			YPos: (rand.Float64()*2 - 1) * 10,
			ZPos: 0,
			XVel: (rand.Float64()*2 - 1) / 1000,
			YVel: (rand.Float64()*2 - 1) / 1000,
			ZVel: 0,
			Mass: rand.Float64(),
		})
	}

	pixelgl.Run(run)

}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	timestamp := time.Now().UTC()

	for !win.Closed() {
		u.Step()
		//fmt.Printf("%+v\n", u)

		farthest := u.FarthestPointFromOrigin()
		massiest := u.LargestMass()

		win.Clear(colornames.Black)

		for _, b := range u.Bodies {
			x := b.XPos/farthest*.75*windowWidth/2 + windowWidth/2
			y := b.YPos/farthest*.75*windowHeight/2 + windowHeight/2
			r := b.Mass / farthest * massiest
			if r < 1 {
				r = 1
			}
			//fmt.Printf("%d: %f %f   ", i, x, y)
			circle := imdraw.New(nil)
			circle.Color = colornames.White
			circle.Push(pixel.V(x, y))
			circle.Circle(r, 0)
			circle.Draw(win)
		}
		//fmt.Println()

		basicTxt := text.New(pixel.V(100, 500), basicAtlas)
		basicTxt.Color = colornames.Red
		fmt.Fprintf(basicTxt, "%4.1f", 1.0/time.Since(timestamp).Seconds())
		timestamp = time.Now().UTC()
		basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 3))

		win.Update()
	}
}
