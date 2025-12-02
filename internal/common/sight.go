package common

// TODO: Move sight to agent/observe package
import (
	"math"
)

type VisionCell struct {
	Colour   string
	Distance float64
}

type Sight struct {
	Angle    float64
	Distance float64
	cells    []VisionCell
}

func (s *Sight) UpdateVision(cells []VisionCell) {
	s.cells = cells
}

func (s Sight) GetVisionCells() []VisionCell {
	return s.cells
}

type SightBuilder struct {
	angle    float64
	distance float64
}

func (sb *SightBuilder) Angle(angle float64) *SightBuilder {
	sb.angle = angle
	return sb
}

func (sb *SightBuilder) Distance(distance float64) *SightBuilder {
	sb.distance = distance
	return sb
}

func (sb *SightBuilder) Build() *Sight {
	cellNumber := sb.angle / (math.Pi / 180)
	cells := make([]VisionCell, int(cellNumber))

	return &Sight{
		Angle:    sb.angle,
		Distance: sb.distance,
		cells:    cells,
	}
}
