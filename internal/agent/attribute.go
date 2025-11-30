package agent

import (
	"math"
	"time"
)

// Attributes are always between 0 and 1
type Attribute struct {
	current   float64
	idealZone []float64
	updater   func(time.Duration, ...float64) float64
}

// How far from the ideal zone is the Attribute
func (a Attribute) DiscomfortAmount() float64 {
	if a.IsInIdealZone() {
		return 0
	}

	return math.Min(
		math.Abs(a.current-a.idealZone[0]),
		math.Abs(a.current-a.idealZone[1]),
	)
}

func (a Attribute) IsInIdealZone() bool {
	return a.current > a.idealZone[0] && a.current < a.idealZone[1]
}

func (a *Attribute) Update(diffTime time.Duration, updateValues ...float64) {
	a.current += a.updater(diffTime, updateValues...)
}

type AttributeBuilder struct {
	initialValue float64
	minIdeal     float64
	maxIdeal     float64
	updater      func(time.Duration, ...float64) float64
}

func NewAttributeBuilder() *AttributeBuilder {
	return &AttributeBuilder{}
}

func (ab *AttributeBuilder) InitialValue(initialValue float64) *AttributeBuilder {
	ab.initialValue = initialValue
	return ab
}

func (ab *AttributeBuilder) MaxIdealValue(maxIdealValue float64) *AttributeBuilder {
	ab.maxIdeal = maxIdealValue
	return ab
}

func (ab *AttributeBuilder) MinIdealValue(minIdealValue float64) *AttributeBuilder {
	ab.minIdeal = minIdealValue
	return ab
}

func (ab *AttributeBuilder) UpdaterFunction(updater func(time.Duration, ...float64) float64) *AttributeBuilder {
	ab.updater = updater
	return ab
}

func (ab AttributeBuilder) Build() *Attribute {
	return &Attribute{
		current: ab.initialValue,
		idealZone: []float64{
			ab.minIdeal,
			ab.maxIdeal,
		},
		updater: ab.updater,
	}
}
