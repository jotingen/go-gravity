package gravity

import (
	//"fmt"
	"math"
	"sync"
)

//import (
//	"github.com/jinzhu/copier"
//)

const G = 6.67408e-11
const Time = 2

type Body struct {
	XPos   float64
	YPos   float64
	ZPos   float64
	XVel   float64
	YVel   float64
	ZVel   float64
	Radius float64
	Mass   float64

	xForce float64
	yForce float64
	zForce float64
}

type Universe struct {
	Bodies []Body
}

//func (u *Universe) Copy() (copy Universe) {
//	copier.Copy(&copy, u)
//	return
//}

func (u *Universe) Step() {
	//fmt.Println("Step")
	var wg sync.WaitGroup

	wg.Add(len(u.Bodies))
	for i := range u.Bodies {
		go func(i int) {
			defer wg.Done()
			var XForce float64
			var YForce float64
			var ZForce float64
			for j := range u.Bodies {
				if i == j {
					continue
				}
				GMass := -G * (u.Bodies[i].Mass * u.Bodies[j].Mass)

				RMag := math.Sqrt(math.Pow(u.Bodies[i].XPos-u.Bodies[j].XPos, 2) +
					math.Pow(u.Bodies[i].YPos-u.Bodies[j].YPos, 2) +
					math.Pow(u.Bodies[i].ZPos-u.Bodies[j].ZPos, 2))

				F := GMass / math.Pow(RMag, 2)

				XForce += F * (u.Bodies[i].XPos - u.Bodies[j].XPos) / RMag
				YForce += F * (u.Bodies[i].YPos - u.Bodies[j].YPos) / RMag
				ZForce += F * (u.Bodies[i].ZPos - u.Bodies[j].ZPos) / RMag

				//fmt.Printf("%d %d\n", i, j)
				//fmt.Printf("GMass  %d %d : %e\n", i, j, GMass)
				//fmt.Printf("RMag   %d %d : %e\n", i, j, RMag)
				//fmt.Printf("F      %d %d : %e\n", i, j, F)
				//fmt.Printf("XForce %d %d : %e\n", i, j, XForce)
				//fmt.Printf("YForce %d %d : %e\n", i, j, YForce)
			}
			u.Bodies[i].xForce = XForce
			u.Bodies[i].yForce = YForce
			u.Bodies[i].zForce = ZForce
		}(i)
	}
	wg.Wait()
	//for i := range u.Bodies {
	//	fmt.Printf("%d %10e %10e  ", i, u.Bodies[i].xForce, u.Bodies[i].yForce)
	//}
	//fmt.Println()
	//fmt.Println("Force")
	//fmt.Println(XForces)
	//fmt.Println(YForces)
	//fmt.Println(ZForces)

	wg.Add(len(u.Bodies))
	for i := range u.Bodies {
		go func(i int) {
			defer wg.Done()
			u.Bodies[i].XVel += u.Bodies[i].xForce / u.Bodies[i].Mass * Time
			u.Bodies[i].YVel += u.Bodies[i].yForce / u.Bodies[i].Mass * Time
			u.Bodies[i].ZVel += u.Bodies[i].zForce / u.Bodies[i].Mass * Time

			u.Bodies[i].XPos += u.Bodies[i].XVel * Time
			u.Bodies[i].YPos += u.Bodies[i].YVel * Time
			u.Bodies[i].ZPos += u.Bodies[i].ZVel * Time
		}(i)
	}
	wg.Wait()

	//Check for collisions
	for i := 0; i < len(u.Bodies); i++ {
		for j := i + 1; j < len(u.Bodies); j++ {
			distance := math.Sqrt(math.Pow(u.Bodies[i].XPos-u.Bodies[j].XPos, 2) +
				math.Pow(u.Bodies[i].YPos-u.Bodies[j].YPos, 2) +
				math.Pow(u.Bodies[i].ZPos-u.Bodies[j].ZPos, 2))
			if distance <= u.Bodies[i].Radius+u.Bodies[j].Radius {
				u.Bodies[i] = Body{
					XPos:   ((u.Bodies[i].Mass * u.Bodies[i].XPos) + (u.Bodies[j].Mass * u.Bodies[j].XPos)) / (u.Bodies[i].Mass + u.Bodies[j].Mass),
					YPos:   ((u.Bodies[i].Mass * u.Bodies[i].YPos) + (u.Bodies[j].Mass * u.Bodies[j].YPos)) / (u.Bodies[i].Mass + u.Bodies[j].Mass),
					ZPos:   ((u.Bodies[i].Mass * u.Bodies[i].ZPos) + (u.Bodies[j].Mass * u.Bodies[j].ZPos)) / (u.Bodies[i].Mass + u.Bodies[j].Mass),
					XVel:   ((u.Bodies[i].Mass * u.Bodies[i].XVel) + (u.Bodies[j].Mass * u.Bodies[j].XVel)) / (u.Bodies[i].Mass + u.Bodies[j].Mass),
					YVel:   ((u.Bodies[i].Mass * u.Bodies[i].YVel) + (u.Bodies[j].Mass * u.Bodies[j].YVel)) / (u.Bodies[i].Mass + u.Bodies[j].Mass),
					ZVel:   ((u.Bodies[i].Mass * u.Bodies[i].ZVel) + (u.Bodies[j].Mass * u.Bodies[j].ZVel)) / (u.Bodies[i].Mass + u.Bodies[j].Mass),
					Radius: math.Sqrt(math.Pow(u.Bodies[i].Radius, 2) + math.Pow(u.Bodies[j].Radius, 2)),
					Mass:   (u.Bodies[i].Mass + u.Bodies[j].Mass),
				}
				u.Bodies = append(u.Bodies[:j], u.Bodies[j+1:]...)
				j--
			}
		}
	}

}

func (u *Universe) FarthestPointFromOrigin() float64 {
	var farthest float64
	for i := range u.Bodies {
		r := math.Sqrt(math.Pow(u.Bodies[i].XPos, 2) + math.Pow(u.Bodies[i].YPos, 2) + math.Pow(u.Bodies[i].ZPos, 2))
		if math.Abs(r) > math.Abs(farthest) {
			farthest = r
		}
	}
	return farthest
}

func (u *Universe) FarthestXPointFromOrigin() float64 {
	var farthest float64
	for i := range u.Bodies {
		r := u.Bodies[i].XPos
		if math.Abs(r) > math.Abs(farthest) {
			farthest = r
		}
	}
	return farthest
}

func (u *Universe) FarthestYPointFromOrigin() float64 {
	var farthest float64
	for i := range u.Bodies {
		r := u.Bodies[i].YPos
		if math.Abs(r) > math.Abs(farthest) {
			farthest = r
		}
	}
	return farthest
}

func (u *Universe) FarthestZPointFromOrigin() float64 {
	var farthest float64
	for i := range u.Bodies {
		r := u.Bodies[i].ZPos
		if math.Abs(r) > math.Abs(farthest) {
			farthest = r
		}
	}
	return farthest
}

func (u *Universe) LargestMass() float64 {
	var mass float64
	for i := range u.Bodies {
		m := u.Bodies[i].Mass
		if m > mass {
			mass = m
		}
	}
	return mass
}

func (u *Universe) TotalMass() float64 {
	var mass float64
	for i := range u.Bodies {
		mass += u.Bodies[i].Mass
	}
	return mass
}

func (u *Universe) XCenterOfMass() float64 {
	var moments float64
	for i := range u.Bodies {
		moments += u.Bodies[i].XPos * u.Bodies[i].Mass
	}
	return moments / u.TotalMass()
}

func (u *Universe) YCenterOfMass() float64 {
	var moments float64
	for i := range u.Bodies {
		moments += u.Bodies[i].YPos * u.Bodies[i].Mass
	}
	return moments / u.TotalMass()
}

func (u *Universe) ZCenterOfMass() float64 {
	var moments float64
	for i := range u.Bodies {
		moments += u.Bodies[i].ZPos * u.Bodies[i].Mass
	}
	return moments / u.TotalMass()
}
