package main

import (
	"fmt"
)

import (
	"github.com/jotingen/go-gravity/gravity"
)

func main() {
	fmt.Println("go-gravity")
	b1 := gravity.Body{
		XPos: 0,
		YPos: 0,
		ZPos: 0,
		XVel: 0,
		YVel: 0,
		ZVel: 0,
		Mass: 10000,
	}
	b2 := gravity.Body{
		XPos: 0,
		YPos: 1,
		ZPos: 0,
		XVel: .01,
		YVel: 0,
		ZVel: 0,
		Mass: 1,
	}
	u := gravity.Universe{}
	u.Bodies = append(u.Bodies, b1)
	u.Bodies = append(u.Bodies, b2)

	for i := 0; i < 1000; i++ {
		fmt.Printf("%+v\n", u)
		u.Step()
	}
}
