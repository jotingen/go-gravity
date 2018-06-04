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
	//	"github.com/jinzhu/copier"
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
	u       gravity.Universe
	history []gravity.Universe
)

func main() {

	fmt.Println("go-gravity")
	u = gravity.Universe{}

	//Grid
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

	//3 Body
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

	//Circles
	//for i := 0; i < 400; i++ {
	//	a := rand.Float64()
	//	b := rand.Float64()
	//	if b < a {
	//		c := a
	//		a = b
	//		b = c
	//	}
	//	u.Bodies = append(u.Bodies, gravity.Body{
	//		XPos:   11 + b*10*math.Cos(2*math.Pi*a/b),
	//		YPos:   b * 10 * math.Sin(2*math.Pi*a/b),
	//		ZPos:   0,
	//		XVel:   .001 * (rand.Float64()*2 - 1),
	//		YVel:   -.015 + .001*(rand.Float64()*2-1),
	//		ZVel:   0,
	//		Radius: .025,
	//		Mass:   100000,
	//	})
	//}

	//for i := 0; i < 400; i++ {
	//	a := rand.Float64()
	//	b := rand.Float64()
	//	if b < a {
	//		c := a
	//		a = b
	//		b = c
	//	}
	//	u.Bodies = append(u.Bodies, gravity.Body{
	//		XPos:   -11 + b*10*math.Cos(2*math.Pi*a/b),
	//		YPos:   b * 10 * math.Sin(2*math.Pi*a/b),
	//		ZPos:   0,
	//		XVel:   .001 * (rand.Float64()*2 - 1),
	//		YVel:   .015 + .001*(rand.Float64()*2-1),
	//		ZVel:   0,
	//		Radius: .025,
	//		Mass:   100000,
	//	})
	//}

	//for i := 0; i < 400; i++ {
	//	a := rand.Float64()
	//	b := rand.Float64()
	//	if b < a {
	//		c := a
	//		a = b
	//		b = c
	//	}
	//	u.Bodies = append(u.Bodies, gravity.Body{
	//		XPos:   b * 10 * math.Cos(2*math.Pi*a/b),
	//		YPos:   11 + b*10*math.Sin(2*math.Pi*a/b),
	//		ZPos:   0,
	//		XVel:   .015 + .001*(rand.Float64()*2-1),
	//		YVel:   .001 * (rand.Float64()*2 - 1),
	//		ZVel:   0,
	//		Radius: .025,
	//		Mass:   100000,
	//	})
	//}

	//for i := 0; i < 400; i++ {
	//	a := rand.Float64()
	//	b := rand.Float64()
	//	if b < a {
	//		c := a
	//		a = b
	//		b = c
	//	}
	//	u.Bodies = append(u.Bodies, gravity.Body{
	//		XPos:   b * 10 * math.Cos(2*math.Pi*a/b),
	//		YPos:   -11 + b*10*math.Sin(2*math.Pi*a/b),
	//		ZPos:   0,
	//		XVel:   -.015 + .001*(rand.Float64()*2-1),
	//		YVel:   .001 * (rand.Float64()*2 - 1),
	//		ZVel:   0,
	//		Radius: .025,
	//		Mass:   100000,
	//	})
	//}

	//Collision
	//u.Bodies = append(u.Bodies, gravity.Body{
	//	XPos:   10,
	//	YPos:   -10,
	//	ZPos:   0,
	//	XVel:   -.01,
	//	YVel:   .01,
	//	ZVel:   0,
	//	Radius: 1,
	//	Mass:   10,
	//})
	//u.Bodies = append(u.Bodies, gravity.Body{
	//	XPos:   -10,
	//	YPos:   -10,
	//	ZPos:   0,
	//	XVel:   .01,
	//	YVel:   .01,
	//	ZVel:   0,
	//	Radius: 1,
	//	Mass:   10,
	//})

	//Earth Moon
	//u.Bodies = append(u.Bodies, gravity.Body{
	//	XPos:   0,
	//	YPos:   0,
	//	ZPos:   0,
	//	XVel:   0,
	//	YVel:   0,
	//	ZVel:   0,
	//	Radius: 6371e3,
	//	Mass:   597237e24,
	//})
	//u.Bodies = append(u.Bodies, gravity.Body{
	//	XPos:   384399e3,
	//	YPos:   0,
	//	ZPos:   0,
	//	XVel:   0,
	//	YVel:   1.022e3,
	//	ZVel:   0,
	//	Radius: 1.7371e3,
	//	Mass:   7342e22,
	//})

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

	counter := 0
	for !win.Closed() {
		for len(u.Bodies) < 10 {
			a := rand.Float64()
			b := rand.Float64()
			if b < a {
				c := a
				a = b
				b = c
			}
			u.Bodies = append(u.Bodies, gravity.Body{
				XPos:   100 * b * math.Cos(2*math.Pi*a/b),
				YPos:   100 * b * math.Sin(2*math.Pi*a/b),
				ZPos:   0,
				XVel:   .05 * (rand.Float64()*2 - 1),
				YVel:   .05 * (rand.Float64()*2 - 1),
				ZVel:   0,
				Radius: 1,
				Mass:   100000000,
			})
		}

		//History
		if counter%30 == 0 {
			h := u
			h.Bodies = make([]gravity.Body, len(u.Bodies))
			copy(h.Bodies, u.Bodies)
			history = append(history, h)
			if len(history) > 100 {
				history = history[len(history)-100 : len(history)-1]
			}
		}
		counter++

		u.Step()

		updateZoom(&zoom, &u)

		win.Clear(colornames.Black)

		//Objects
		circle := imdraw.New(nil)
		circle.Color = colornames.Grey
		for _, h := range history {
			for _, b := range h.Bodies {
				x := zoom*.95*(b.XPos-u.XCenterOfMass()) + windowWidth/2
				y := zoom*.95*(b.YPos-u.YCenterOfMass()) + windowHeight/2
				r := zoom * b.Radius
				if r < 1 {
					r = 1
				}
				circle.Push(pixel.V(x, y))
			}
		}
		circle.Circle(1, 0)
		circle.Color = colornames.White
		for _, b := range u.Bodies {
			x := zoom*.95*(b.XPos-u.XCenterOfMass()) + windowWidth/2
			y := zoom*.95*(b.YPos-u.YCenterOfMass()) + windowHeight/2
			r := zoom * b.Radius
			if r < 1 {
				r = 2
			}
			circle.Push(pixel.V(x, y))
			circle.Circle(r, 0)
		}
		circle.Draw(win)

		//FPS
		fpsTxt := text.New(pixel.V(0+40, windowHeight-40), basicAtlas)
		fpsTxt.Color = colornames.Red
		fmt.Fprintf(fpsTxt, "%2.0f", 1.0/time.Since(timestamp).Seconds())
		timestamp = time.Now().UTC()
		fpsTxt.Draw(win, pixel.IM.Scaled(fpsTxt.Orig, 1))

		//Zoom
		zoomTxt := text.New(pixel.V(0+40, 0+40), basicAtlas)
		zoomTxt.Color = colornames.Red
		fmt.Fprintf(zoomTxt, "%4.2f m/pixel", 1/zoom)
		timestamp = time.Now().UTC()
		zoomTxt.Draw(win, pixel.IM.Scaled(zoomTxt.Orig, 1))

		//Body Count
		bodiesTxt := text.New(pixel.V(windowWidth-40, 0+40), basicAtlas)
		bodiesTxt.Color = colornames.Red
		fmt.Fprintf(bodiesTxt, "%d", len(u.Bodies))
		timestamp = time.Now().UTC()
		bodiesTxt.Draw(win, pixel.IM.Scaled(bodiesTxt.Orig, 1))
		win.Update()
	}
}

func updateZoom(zoom *float64, u *gravity.Universe) {
	//farthest := u.FarthestPointFromOrigin()
	//fmt.Println(u.XCenterOfMass())
	farthestX := math.Abs(u.FarthestXPointFromOrigin())
	farthestY := math.Abs(u.FarthestYPointFromOrigin())
	zoomX := (float64(windowWidth) / 2) / farthestX
	zoomY := (float64(windowHeight) / 2) / farthestY
	//fmt.Printf("%5f ", farthestX)
	//fmt.Printf("%5f  ", zoomX)
	//fmt.Printf("%5f ", u.FarthestYPointFromOrigin())
	//fmt.Printf("%5f ", farthestY)
	//fmt.Printf("%5f  ", zoomY)
	if zoomX < *zoom || *zoom == 0.0 {
		*zoom = zoomX
	}
	if zoomY < *zoom {
		*zoom = zoomY
	}
	//fmt.Printf("%5f\n", zoom)
	//massiest := u.LargestMass()
}
