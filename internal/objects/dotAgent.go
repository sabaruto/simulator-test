package objects

import (
	"math"
	"time"

	"github.com/quartercastle/vector"
)

type DotAgent struct {
	Base
	radius   float64
	colour   string
	velocity vector.Vector
}

func NewDotAgent(position vector.Vector, radius float64) *DotAgent {
	return &DotAgent{
		Base: Base{
			position: position,
		},
		radius: radius,
	}
}

func (d *DotAgent) Move(duration time.Duration) {
	d.velocity = d.velocity.Rotate(math.Pi * duration.Minutes() * 4)
}

func (d DotAgent) Draw() {
	// Draw dot
	d.DrawDot()

	// Draw orientation pointer
	d.drawPointer()
}

func (d DotAgent) DrawDot() {
	cv := d.GetCanvas()
	cv.SetFillStyle(d.colour)
	cv.BeginPath()
	cv.MoveTo(d.position.X(), d.position.Y())
	cv.Arc(d.position.X(), d.position.Y(), d.radius, 0, 2*math.Pi, false)
	cv.ClosePath()
	cv.Fill()
}

func (d DotAgent) drawPointer() {
	cv := d.GetCanvas()
	arrowOffset := d.radius * 1.4
	arrowSize := d.radius / 2

	arrowVector := d.velocity.Scale(arrowOffset)
	perpAngle := d.velocity.Rotate(math.Pi / 2)
	arrowPosition := d.position.Add(arrowVector)

	cv.BeginPath()

	cv.MoveTo(
		arrowPosition.X(),
		arrowPosition.Y(),
	)
	cv.LineTo(
		arrowPosition.X()+perpAngle.X()*(arrowSize/2),
		arrowPosition.Y()+perpAngle.Y()*(arrowSize/2),
	)
	cv.LineTo(
		d.position.X()+(d.velocity.X()*(arrowOffset+arrowSize)),
		d.position.Y()+(d.velocity.Y()*(arrowOffset+arrowSize)),
	)
	cv.LineTo(
		arrowPosition.X()-perpAngle.X()*(arrowSize/2),
		arrowPosition.Y()-perpAngle.Y()*(arrowSize/2),
	)
	cv.ClosePath()
	cv.Fill()
}

type DotAgentBuilder struct {
	position vector.Vector
	radius   float64
	colour   string
	velocity vector.Vector
}

func NewDotAgentBuilder() *DotAgentBuilder {
	return &DotAgentBuilder{}
}

func (dab *DotAgentBuilder) Position(x float64, y float64) *DotAgentBuilder {
	dab.position = vector.Vector{x, y}
	return dab
}

func (dab *DotAgentBuilder) Radius(radius float64) *DotAgentBuilder {
	dab.radius = radius
	return dab
}

func (dab *DotAgentBuilder) Colour(colour string) *DotAgentBuilder {
	dab.colour = colour
	return dab
}

func (dab *DotAgentBuilder) Angle(angle float64) *DotAgentBuilder {
	dab.velocity = vector.Vector{1, 0}.Rotate(angle)
	return dab
}

func (dab *DotAgentBuilder) Build() *DotAgent {
	if dab.velocity == nil {
		dab.velocity = vector.Vector{1, 0}
	}

	return &DotAgent{
		Base: Base{
			position: dab.position,
		},
		radius:   dab.radius,
		colour:   dab.colour,
		velocity: dab.velocity,
	}
}
