package common

type VisionCell struct {
	Colour   string
	Distance float64
}

type Sight interface {
	GetViewAngle() float64
	GetVisionCells() []VisionCell
	SetObjectManager(ObjectManager)
	UpdateVision()
}
