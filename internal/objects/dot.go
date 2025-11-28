package objects

import (
	"math"
	"time"

	"github.com/quartercastle/vector"
)

type Dot struct {
	Base
	radius float64
}

func NewDot(position vector.Vector, radius float64) *Dot {
	return &Dot{
		Base: Base{
			position: position,
		},

		radius: radius,
	}
}

func (d *Dot) Move(duration time.Duration) {
	d.radius += duration.Seconds()
}

func (d Dot) Draw() {
	cv := d.GetCanvas()
	cv.SetFillStyle(128, 128, 128)
	cv.BeginPath()
	cv.MoveTo(d.position.X(), d.position.Y())
	cv.Arc(d.position.X(), d.position.Y(), d.radius, 0, 2*math.Pi, false)
	cv.ClosePath()
	cv.Fill()
}
