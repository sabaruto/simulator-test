package agent

import "time"

type InternalState struct {
	fuel   Attribute
	energy Attribute
}

func NewInternalState() *InternalState {
	return &InternalState{
		fuel: *NewAttributeBuilder().
			InitialValue(0.6).
			MaxIdealValue(0.7).
			MinIdealValue(0.4).
			UpdaterFunction(func(diffTime time.Duration, f ...float64) float64 {
				newValue := -0.4
				if len(f) > 1 {
					newValue += f[0]
				}

				return newValue * diffTime.Minutes()
			}).
			Build(),
		energy: *NewAttributeBuilder().
			InitialValue(0.6).
			MaxIdealValue(0.8).
			MinIdealValue(0.5).
			UpdaterFunction(func(diffTime time.Duration, f ...float64) float64 {
				newValue := -0.2

				if len(f) > 1 {
					newValue += f[0]
				}

				return newValue * diffTime.Minutes()
			}).
			Build(),
	}
}

func (i *InternalState) UpdateState(diffTime time.Duration) {
	i.fuel.Update(diffTime)
	i.energy.Update(diffTime)
}

func (i *InternalState) GetEnergy() float64 {
	return i.energy.current
}

func (i *InternalState) GetFuel() float64 {
	return i.fuel.current
}
