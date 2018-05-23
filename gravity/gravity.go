package gravity

import (
	//"fmt"
	"math"
)

const G = 6.67408e-11
const Time = 1

type Body struct {
	XPos float64
	YPos float64
	ZPos float64
	XVel float64
	YVel float64
	ZVel float64
	Mass float64
}

type Universe struct {
	Bodies []Body
}

func (u *Universe) Step() {
	//fmt.Println("Step")
	//TODO convert to be concurrent
	var XForces []float64
	var YForces []float64
	var ZForces []float64
	for i := range u.Bodies {
		var XForce float64
		var YForce float64
		var ZForce float64
		for j := range u.Bodies {
			if i == j {
				continue
			}
			GMass := -G * (u.Bodies[i].Mass * u.Bodies[j].Mass)
			if u.Bodies[i].XPos-u.Bodies[j].XPos != 0 {
				if u.Bodies[i].XPos-u.Bodies[j].XPos > 0 {
					XForce += GMass / math.Pow(u.Bodies[i].XPos-u.Bodies[j].XPos, 2)
				} else {
					XForce -= GMass / math.Pow(u.Bodies[i].XPos-u.Bodies[j].XPos, 2)
				}
			}
			if u.Bodies[i].YPos-u.Bodies[j].YPos != 0 {
				if u.Bodies[i].YPos-u.Bodies[j].YPos > 0 {
					YForce += GMass / math.Pow(u.Bodies[i].YPos-u.Bodies[j].YPos, 2)
				} else {
					YForce -= GMass / math.Pow(u.Bodies[i].YPos-u.Bodies[j].YPos, 2)
				}
			}
			if u.Bodies[i].ZPos-u.Bodies[j].ZPos != 0 {
				if u.Bodies[i].ZPos-u.Bodies[j].ZPos > 0 {
					ZForce += GMass / math.Pow(u.Bodies[i].ZPos-u.Bodies[j].ZPos, 2)
				} else {
					ZForce -= GMass / math.Pow(u.Bodies[i].ZPos-u.Bodies[j].ZPos, 2)
				}
			}
			//fmt.Printf("GMass  %d %d : %e\n", i, j, GMass)
			//fmt.Printf("XPos   %d %d : %e %e\n", i, j, u.Bodies[i].XPos, u.Bodies[j].XPos)
			//fmt.Printf("r      %d %d : %e\n", i, j, u.Bodies[i].XPos-u.Bodies[j].XPos)
			//fmt.Printf("XForce %d %d : %e\n", i, j, XForce)
		}
		XForces = append(XForces, XForce)
		YForces = append(YForces, YForce)
		ZForces = append(ZForces, ZForce)
	}
	//fmt.Println("Force")
	//fmt.Println(XForces)
	//fmt.Println(YForces)
	//fmt.Println(ZForces)

	for i := range u.Bodies {
		if XForces[i] > 0 {
			u.Bodies[i].XVel += math.Sqrt(2 * math.Abs(XForces[i]) / u.Bodies[i].Mass)
		} else {
			u.Bodies[i].XVel -= math.Sqrt(2 * math.Abs(XForces[i]) / u.Bodies[i].Mass)
		}
		if YForces[i] > 0 {
			u.Bodies[i].YVel += math.Sqrt(2 * math.Abs(YForces[i]) / u.Bodies[i].Mass)
		} else {
			u.Bodies[i].YVel -= math.Sqrt(2 * math.Abs(YForces[i]) / u.Bodies[i].Mass)
		}
		if ZForces[i] > 0 {
			u.Bodies[i].ZVel += math.Sqrt(2 * math.Abs(ZForces[i]) / u.Bodies[i].Mass)
		} else {
			u.Bodies[i].ZVel -= math.Sqrt(2 * math.Abs(ZForces[i]) / u.Bodies[i].Mass)
		}
	}
	for i := range u.Bodies {
		u.Bodies[i].XPos += u.Bodies[i].XVel * Time
		u.Bodies[i].YPos += u.Bodies[i].YVel * Time
		u.Bodies[i].ZPos += u.Bodies[i].ZVel * Time
	}
}
