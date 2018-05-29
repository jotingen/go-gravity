package main

import (
	"fmt"
	"math"
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
	//for i := -10; i <= 10; i++ {
	//	for j := -10; j <= 10; j++ {
	//		u.Bodies = append(u.Bodies, gravity.Body{
	//			XPos: float64(i),
	//			YPos: float64(j),
	//			ZPos: 0,
	//			XVel: 0,
	//			YVel: 0,
	//			ZVel: 0,
	//			Mass: 1,
	//		})
	//	}
	//}

	//u.Bodies = append(u.Bodies, gravity.Body{
	//	XPos: 2,
	//	YPos: 0,
	//	ZPos: 0,
	//	XVel: 0,
	//	YVel: .001,
	//	ZVel: 0,
	//	Mass: 10,
	//})
	//u.Bodies = append(u.Bodies, gravity.Body{
	//	XPos: -2,
	//	YPos: 0,
	//	ZPos: 0,
	//	XVel: 0,
	//	YVel: -.001,
	//	ZVel: 0,
	//	Mass: 10,
	//})
	//u.Bodies = append(u.Bodies, gravity.Body{
	//	XPos: 0,
	//	YPos: 10,
	//	ZPos: 0,
	//	XVel: .001,
	//	YVel: 0,
	//	ZVel: 0,
	//	Mass: 20,
	//})

	for i := 0; i < 1000; i++ {
		a := rand.Float64()
		b := rand.Float64()
		if b < a {
			c := a
			a = b
			b = c
		}
		u.Bodies = append(u.Bodies, gravity.Body{
			XPos: b * 100 * math.Cos(2*math.Pi*a/b),
			YPos: b * 100 * math.Sin(2*math.Pi*a/b),
			ZPos: 0,
			XVel: (rand.Float64()*2 - 1) / 10000,
			YVel: (rand.Float64()*2 - 1) / 10000,
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

	zoom := 0.0

	for !win.Closed() {
		u.Step()
		//fmt.Printf("%+v\n", u)

		windowSmallest := float64(windowWidth)
		if windowHeight < windowSmallest {
			windowSmallest = float64(windowHeight)
		}

		//farthest := u.FarthestPointFromOrigin()
		farthestX := u.FarthestXPointFromOrigin()
		farthestY := u.FarthestYPointFromOrigin()
		zoomX := (float64(windowWidth) / 2) / farthestX
		zoomY := (float64(windowHeight) / 2) / farthestY
		//fmt.Printf("%5f  ", zoomX)
		//fmt.Printf("%5f  ", zoomY)
		if zoomX < zoom || zoom == 0.0 {
			zoom = zoomX
		}
		if zoomY < zoom {
			zoom = zoomY
		}
		//fmt.Printf("%5f\n", zoom)
		//massiest := u.LargestMass()

		win.Clear(colornames.Black)

		for _, b := range u.Bodies {
			x := b.XPos*zoom*.95 + windowWidth/2
			y := b.YPos*zoom*.95 + windowHeight/2
			r := b.Mass * zoom * .01
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
