package square

import (
	"math"
)

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

type SidesAmount int

// Define custom int type to hold sides number and
// update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3),
// SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum SidesAmount) float64 {
	switch sidesNum {
	case SidesCircle:
		return sideLen * sideLen * math.Pi
	case SidesTriangle:
		return (sideLen * sideLen * math.Sqrt(3)) / 4
	case SidesSquare:
		return sideLen * sideLen
	}
}
