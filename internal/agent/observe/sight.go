package observe

// TODO: Move sight to agent/observe package
import (
	"math"

	"github.com/quartercastle/vector"
	"github.com/sabaruto/simulator-test/internal/common"
)

const ONE_DEGREE_RADIAN float64 = math.Pi / 180

type Sight struct {
	Observe
	ViewAngle     float64
	halfViewAngle float64
	Distance      float64
	cells         []common.VisionCell
}

func (s *Sight) UpdateVision() {
	visionCells := make([]common.VisionCell, 0)
	for rayAngle := s.objectVelocity.Angle() - s.halfViewAngle; rayAngle < s.objectVelocity.Angle()+s.halfViewAngle; rayAngle += ONE_DEGREE_RADIAN {
		rayEndPosition, rayColour := s.objectManager.SendRay(*s.objectPosition, vector.Vector{1, 0}.Rotate(rayAngle), []int64{*s.objectID})
		var visionCell common.VisionCell

		if rayEndPosition == nil || rayEndPosition.Sub(*s.objectPosition).Magnitude() > s.Distance {
			visionCell = common.VisionCell{Colour: "#000000", Distance: s.Distance}
		} else {
			visionCell = common.VisionCell{Colour: *rayColour, Distance: rayEndPosition.Sub(*s.objectPosition).Magnitude()}
		}
		visionCells = append(visionCells, visionCell)
	}
	s.cells = visionCells
}

func (s Sight) GetViewAngle() float64 {
	return s.ViewAngle
}

func (s Sight) GetVisionCells() []common.VisionCell {
	return s.cells
}

func (s *Sight) SetObjectManager(objectManager common.ObjectManager) {
	s.objectManager = objectManager
}

type SightBuilder struct {
	objectPosition *vector.Vector
	objectVelocity *vector.Vector
	objectID       *int64
	angle          float64
	distance       float64
}

func NewSightBuilder() *SightBuilder {
	return &SightBuilder{}
}

func (sb *SightBuilder) Angle(angle float64) *SightBuilder {
	sb.angle = angle
	return sb
}

func (sb *SightBuilder) Distance(distance float64) *SightBuilder {
	sb.distance = distance
	return sb
}

func (sb *SightBuilder) ObjectPosition(objectPosition *vector.Vector) *SightBuilder {
	sb.objectPosition = objectPosition
	return sb
}

func (sb *SightBuilder) ObjectVelocity(objectVelocity *vector.Vector) *SightBuilder {
	sb.objectVelocity = objectVelocity
	return sb
}

func (sb *SightBuilder) ObjectID(id *int64) *SightBuilder {
	sb.objectID = id
	return sb
}

func (sb *SightBuilder) Build() common.Sight {
	cellNumber := sb.angle / ONE_DEGREE_RADIAN
	cells := make([]common.VisionCell, int(cellNumber))

	return &Sight{
		Observe: Observe{
			objectPosition: sb.objectPosition,
			objectVelocity: sb.objectVelocity,
			objectID:       sb.objectID,
		},
		ViewAngle:     sb.angle,
		halfViewAngle: sb.angle / 2,
		Distance:      sb.distance,
		cells:         cells,
	}
}
