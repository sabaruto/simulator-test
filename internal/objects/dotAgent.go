package objects

import (
	"math"
	"time"

	"github.com/sabaruto/simulator-test/internal/common"
)

type DotAgent struct {
	Base
	radius float64
	colour string
	// TODO: Create an angle type to quickly get cos / sin of values
	orientation float64
}

func NewDotAgent(position common.Position, radius float64) *DotAgent {
	return &DotAgent{
		Base: Base{
			position: position,
		},
		radius: radius,
	}
}

func (d *DotAgent) Move(duration time.Duration) {
	d.orientation += math.Pi * duration.Minutes() * 4
}

func (d DotAgent) Draw() {
	cv := d.GetCanvas()

	// Draw dot
	cv.SetFillStyle(d.colour)
	cv.BeginPath()
	cv.MoveTo(d.position.X, d.position.Y)
	cv.Arc(d.position.X, d.position.Y, d.radius, 0, 2*math.Pi, false)
	cv.ClosePath()
	cv.Fill()

	// Draw orientation pointer
	arrowOffset := d.radius * 1.4
	arrowSize := d.radius / 2

	arrowPosition := common.Position{
		X: d.position.X + math.Cos(d.orientation)*arrowOffset,
		Y: d.position.Y + math.Sin(d.orientation)*arrowOffset,
	}

	cv.BeginPath()
	cv.MoveTo(
		arrowPosition.X,
		arrowPosition.Y,
	)
	cv.LineTo(
		d.position.X+(math.Cos(d.orientation)*arrowOffset)+math.Cos(d.orientation+math.Pi/2)*(arrowSize/2),
		d.position.Y+(math.Sin(d.orientation)*arrowOffset)+math.Sin(d.orientation+math.Pi/2)*(arrowSize/2),
	)
	cv.LineTo(
		d.position.X+(math.Cos(d.orientation)*(arrowOffset+arrowSize)),
		d.position.Y+(math.Sin(d.orientation)*(arrowOffset+arrowSize)),
	)
	cv.LineTo(
		d.position.X+(math.Cos(d.orientation)*arrowOffset)-math.Cos(d.orientation+math.Pi/2)*(arrowSize/2),
		d.position.Y+(math.Sin(d.orientation)*arrowOffset)-math.Sin(d.orientation+math.Pi/2)*(arrowSize/2),
	)
	cv.ClosePath()
	cv.Fill()
}

type DotAgentBuilder struct {
	position common.Position
	radius   float64
	colour   string
}

func NewDotAgentBuilder() *DotAgentBuilder {
	return &DotAgentBuilder{}
}

func (dab *DotAgentBuilder) Position(x float64, y float64) *DotAgentBuilder {
	dab.position = common.Position{X: x, Y: y}
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

func (dab *DotAgentBuilder) Build() *DotAgent {
	return &DotAgent{
		Base: Base{
			position: dab.position,
		},
		radius:      dab.radius,
		colour:      dab.colour,
		orientation: 0,
	}
}
