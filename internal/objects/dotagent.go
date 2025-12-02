package objects

import (
	"fmt"
	"math"
	"time"

	"github.com/quartercastle/vector"
	"github.com/sabaruto/simulator-test/internal/agent"
	"github.com/sabaruto/simulator-test/internal/common"
)

type DotAgent struct {
	BaseObject
	internalState agent.InternalState
	radius        float64
	colour        string
	velocity      vector.Vector
	sight         common.Sight
}

func (d *DotAgent) Move(duration time.Duration) {
	d.observe()
	d.act(duration)
}

func (d *DotAgent) observe() {
	d.updateVision()
}

func (d *DotAgent) act(diffTime time.Duration) {

	d.velocity = d.velocity.Rotate(math.Pi * diffTime.Minutes() * 8)

	d.internalState.UpdateState(diffTime)
}

func (d DotAgent) ReceiveRay(rayPosition vector.Vector, rayDirection vector.Vector) (intersectPoint *vector.Vector, colour *string) {
	rayDirection = rayDirection.Unit()
	intersections := common.GetCircleIntersections(rayPosition, rayDirection, d.position, d.radius)

	// Check an intersection point is found
	if intersections == nil {
		return nil, nil
	}

	var closestPoint *vector.Vector
	smallestDistance := math.Inf(1)

	for _, intersectPoint := range *intersections {
		intersectDistance := intersectPoint.Sub(rayPosition).X() / rayDirection.X()

		if math.IsNaN(intersectDistance) {
			intersectDistance = intersectPoint.Sub(rayPosition).Y() / rayDirection.Y()
		}

		if intersectDistance > 0 && intersectDistance < smallestDistance {
			closestPoint = &intersectPoint
			smallestDistance = intersectDistance
		}
	}

	if closestPoint == nil {
		return nil, nil
	}

	return closestPoint, &d.colour
}

func (d *DotAgent) updateVision() {
	visionCells := make([]common.VisionCell, 0)
	for rayAngle := d.velocity.Angle() - (d.sight.Angle / 2); rayAngle < d.velocity.Angle()+(d.sight.Angle/2); rayAngle += math.Pi / 180 {
		rayEndPosition, rayColour := d.objectManager.SendRay(d.position, vector.Vector{1, 0}.Rotate(rayAngle), []int64{d.GetID()})
		var visionCell common.VisionCell

		if rayEndPosition == nil || rayEndPosition.Sub(d.position).Magnitude() > d.sight.Distance {
			visionCell = common.VisionCell{Colour: "#000000", Distance: d.sight.Distance}
		} else {
			visionCell = common.VisionCell{Colour: *rayColour, Distance: rayEndPosition.Sub(d.position).Magnitude()}
		}
		visionCells = append(visionCells, visionCell)
	}
	d.sight.UpdateVision(visionCells)
}

func (d DotAgent) Draw() {
	if d.debugGetter.GetRaysState() {
		d.drawVision()
	}

	// Draw dot
	d.DrawDot()

	// Draw orientation pointer
	d.drawPointer()

	if d.debugGetter.GetAgentMetricsState() {
		d.drawAgentMetrics()
	}
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
	cv.MoveTo(arrowPosition.X(), arrowPosition.Y())
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
		currentAngle := d.velocity.Angle() - (d.sight.Angle / 2) + (math.Pi/180)*float64(visionIndex)
		rayEndPosition := vector.Vector{1, 0}.Rotate(currentAngle).Scale(visionCell.Distance).Add(d.position)

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

func (d DotAgent) ExpandedString() string {
	return fmt.Sprintf("Dot Agent\n\tid: %d\n\tposition: %v\n\tradius: %f\n\tcolour: %s)", d.GetID(), d.GetPosition(), d.radius, d.colour)
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
	if dab.velocity == nil {
		dab.velocity = vector.Vector{1, 0}
	}

	if dab.vision.Distance == 0 {
		dab.vision = common.Sight{
			Angle:    math.Pi / 3,
			Distance: 160,
		}
	}

	return &DotAgent{
		BaseObject: BaseObject{
			position:    dab.position,
			debugGetter: dab.debugClient,
		},
		internalState: *agent.NewInternalState(),
		radius:        dab.radius,
		colour:        dab.colour,
		velocity:      dab.velocity,
		sight:         dab.vision,
	}
}
