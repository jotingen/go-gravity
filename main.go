package main

import (
	"fmt"
)

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
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
	b1 := gravity.Body{
		XPos: 0,
		YPos: 0,
		ZPos: 0,
		XVel: -.001,
		YVel: 0,
		ZVel: 0,
		Mass: 50,
	}
	b2 := gravity.Body{
		XPos: 0,
		YPos: -10,
		ZPos: 0,
		XVel: .01,
		YVel: 0,
		ZVel: 0,
		Mass: 1,
	}
	b3 := gravity.Body{
		XPos: 0,
		YPos: 10,
		ZPos: 0,
		XVel: -.001,
		YVel: 0,
		ZVel: 0,
		Mass: 10,
	}
	b4 := gravity.Body{
		XPos: 0,
		YPos: -1,
		ZPos: 0,
		XVel: .01,
		YVel: 0,
		ZVel: 0,
		Mass: 1,
	}
	b5 := gravity.Body{
		XPos: 0,
		YPos: -5,
		ZPos: 0,
		XVel: .01,
		YVel: -.01,
		ZVel: 0,
		Mass: 1,
	}
	b6 := gravity.Body{
		XPos: 1,
		YPos: -7,
		ZPos: 0,
		XVel: .01,
		YVel: 0,
		ZVel: 0,
		Mass: 1,
	}
	b7 := gravity.Body{
		XPos: 4,
		YPos: -4,
		ZPos: 0,
		XVel: .01,
		YVel: .01,
		ZVel: 0,
		Mass: 1,
	}
	u = gravity.Universe{}
	u.Bodies = append(u.Bodies, b1)
	u.Bodies = append(u.Bodies, b2)
	u.Bodies = append(u.Bodies, b3)
	u.Bodies = append(u.Bodies, b4)
	u.Bodies = append(u.Bodies, b5)
	u.Bodies = append(u.Bodies, b6)
	u.Bodies = append(u.Bodies, b7)

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
		win.Update()
	}
}
