package objects

import (
	"fmt"
	"math"
	"time"

	"github.com/quartercastle/vector"
	"github.com/sabaruto/simulator-test/internal/agent"
	"github.com/sabaruto/simulator-test/internal/agent/observe"
	"github.com/sabaruto/simulator-test/internal/common"
	"github.com/sabaruto/simulator-test/internal/ray"
	"github.com/sabaruto/simulator-test/internal/shapes"
)

type DotAgent struct {
	object
	internalState agent.InternalState
	radius        float64
	colour        string
	velocity      *vector.Vector
	sight         common.Sight
}

func (d *DotAgent) Move(duration time.Duration) {
	d.observe()
	d.act(duration)
}

func (d *DotAgent) observe() {
	d.updateVision()
}

func (d DotAgent) GetVelocity() vector.Vector {
	return *d.velocity
}

func (d *DotAgent) act(diffTime time.Duration) {

	*d.velocity = d.GetVelocity().Rotate(math.Pi * diffTime.Minutes() * 8)

	d.internalState.UpdateState(diffTime)
}

func (d DotAgent) ReceiveRay(rayPosition vector.Vector, rayDirection vector.Vector) (*vector.Vector, *string) {
	return ray.RecieveRayToCircle(rayPosition, rayDirection, d.position, d.radius, d.colour)
}

func (d *DotAgent) updateVision() {
	d.sight.UpdateVision()
}

func (d *DotAgent) SetObjectManager(objectManager common.ObjectManager) {
	d.object.SetObjectManager(objectManager)
	d.sight.SetObjectManager(objectManager)
}

func (d DotAgent) Draw() {
	if d.debugGetter.GetRaysState() {
		d.drawVision()
	}

	// Draw dot
	shapes.DrawCircle(d.GetCanvas(), d.position, d.colour, d.radius)

	// Draw orientation pointer
	d.drawPointer()

	if d.debugGetter.GetAgentMetricsState() {
		d.drawAgentMetrics()
	}
}

func (d DotAgent) drawPointer() {
	arrowOffset := d.radius * 1.6
	arrowSize := d.radius / 2

	arrowVector := d.velocity.Scale(arrowOffset)
	arrowPosition := arrowVector.Add(*d.position)

	shapes.DrawIsoscelesTriangle(
		d.GetCanvas(),
		&arrowPosition,
		d.colour,
		arrowSize,
		arrowSize,
		d.velocity.Angle(),
	)
}

func (d DotAgent) drawAgentMetrics() {
	fontSize := 14
	cv := d.GetCanvas()
	cv.SetFillStyle("#FFFFFF")
	cv.SetFont("assets/Righteous-Regular.ttf", float64(fontSize))

	debugMetrics := []string{
		fmt.Sprintf("id: %d", d.GetID()),
		fmt.Sprintf("position: %v", d.GetPosition()),
		fmt.Sprintf("radius: %.2f", d.radius),
		fmt.Sprintf("velocity: [%.2f, %.2f] ", d.velocity.X(), d.velocity.Y()),
		fmt.Sprintf("colour: %s", d.colour),
		"Internal State",
		fmt.Sprintf("    energy: %.2f", d.internalState.GetEnergy()),
		fmt.Sprintf("    fuel: %.2f", d.internalState.GetFuel()),
	}

	for metricIndex, metric := range debugMetrics {
		cv.FillText(metric, d.GetPosition().X(), d.GetPosition().Y()+float64(fontSize)*float64(metricIndex))
	}
}

func (d DotAgent) drawVision() {
	cv := d.GetCanvas()
	for visionIndex, visionCell := range d.sight.GetVisionCells() {
		currentAngle := d.velocity.Angle() - (d.sight.GetViewAngle() / 2) + (math.Pi/180)*float64(visionIndex)
		rayEndPosition := vector.Vector{1, 0}.Rotate(currentAngle).Scale(visionCell.Distance).Add(d.GetPosition())

		cv.BeginPath()
		cv.SetStrokeStyle(visionCell.Colour)
		cv.MoveTo(d.position.X(), d.position.Y())
		cv.LineTo(rayEndPosition.X(), rayEndPosition.Y())
		cv.Stroke()
	}
}

func (d DotAgent) String() string {
	return fmt.Sprintf("Dot Agent(id: %d, position: %v, radius: %f, colour: %s)", d.GetID(), d.GetPosition(), d.radius, d.colour)
}

type DotAgentBuilder struct {
	position    vector.Vector
	radius      float64
	colour      string
	velocity    vector.Vector
	vision      common.Sight
	debugClient common.DebugGetter
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

func (dab *DotAgentBuilder) Vision(vision common.Sight) *DotAgentBuilder {
	dab.vision = vision
	return dab
}

func (dab *DotAgentBuilder) DebugClient(debugClient common.DebugGetter) *DotAgentBuilder {
	dab.debugClient = debugClient
	return dab
}

func (dab *DotAgentBuilder) Build() *DotAgent {
	id := createID()
	if dab.velocity == nil {
		dab.velocity = vector.Vector{1, 0}
	}

	if dab.vision == nil {
		dab.vision = observe.NewSightBuilder().
			Angle(math.Pi / 3).
			Distance(160).
			ObjectPosition(&dab.position).
			ObjectVelocity(&dab.velocity).
			ObjectID(id).
			Build()
	}

	return &DotAgent{
		object: object{
			position:    &dab.position,
			debugGetter: dab.debugClient,
			id:          id,
		},
		internalState: *agent.NewInternalState(),
		radius:        dab.radius,
		velocity:      &dab.velocity,
		sight:         dab.vision,
		colour:        dab.colour,
	}
}
