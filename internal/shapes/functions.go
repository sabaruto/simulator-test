package shapes

import (
	"math"

	"github.com/quartercastle/vector"
	"github.com/tfriedel6/canvas"
)

func DrawCircle(cv *canvas.Canvas, position *vector.Vector, colour string, radius float64) {
	cv.SetFillStyle(colour)
	cv.BeginPath()
	cv.MoveTo(position.X(), position.Y())
	cv.Arc(position.X(), position.Y(), radius, 0, 2*math.Pi, false)
	cv.ClosePath()
	cv.Fill()
}

// Draw an isosceles triangle, at an angle where 0 has DrawIsoscelesTriangle
// facing right
func DrawIsoscelesTriangle(cv *canvas.Canvas, position *vector.Vector, colour string, length float64, height float64, angle float64) {
	halfHeight := height / 2
	halfLength := length / 2

	cosAngle := math.Cos(angle)
	sinAngle := math.Sin(angle)

	halfHeightCos := halfHeight * cosAngle
	halfHeightSin := halfHeight * sinAngle
	halfLengthCos := halfLength * cosAngle
	halfLengthSin := halfLength * sinAngle

	pointOne := &vector.Vector{
		halfHeightCos + position.X(),
		halfHeightSin + position.Y(),
	}

	pointTwo := &vector.Vector{
		-halfLengthSin - halfHeightCos + position.X(),
		halfLengthCos - halfHeightSin + position.Y(),
	}

	pointThree := &vector.Vector{
		halfLengthSin - halfHeightCos + position.X(),
		-halfLengthCos - halfHeightSin + position.Y(),
	}

	cv.BeginPath()
	cv.MoveTo(pointOne.X(), pointOne.Y())
	cv.LineTo(pointTwo.X(), pointTwo.Y())
	cv.LineTo(pointThree.X(), pointThree.Y())
	cv.ClosePath()
	cv.Fill()
}
