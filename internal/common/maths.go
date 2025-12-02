package common

import (
	"math"

	"github.com/quartercastle/vector"
)

func sgn(x float64) float64 {
	if x < 0 {
		return -1
	} else {
		return 1
	}
}

func GetCircleIntersections(
	linePosition vector.Vector,
	lineDirection vector.Vector,
	circlePosition vector.Vector,
	radius float64,
) *[]vector.Vector {
	pointOne := linePosition.Sub(circlePosition)
	pointTwo := pointOne.Add(lineDirection)

	directionSquared := lineDirection.Magnitude() * lineDirection.Magnitude()
	D := pointOne.X()*pointTwo.Y() - pointTwo.X()*pointOne.Y()

	discriminant := radius*radius*directionSquared - D*D

	if discriminant < 0 {
		return nil
	}

	workingIntersection := vector.Vector{
		(D * lineDirection.Y()) / (directionSquared),
		-D * lineDirection.X() / (directionSquared),
	}

	workingIntersectionTwo := vector.Vector{
		(sgn(lineDirection.Y()) * lineDirection.X() * math.Sqrt(discriminant)) / (directionSquared),
		(math.Abs(lineDirection.Y()) * math.Sqrt(discriminant)) / (directionSquared),
	}

	intersectOne := workingIntersection.Add(workingIntersectionTwo).Add(circlePosition)
	intersectTwo := workingIntersection.Sub(workingIntersectionTwo).Add(circlePosition)

	return &[]vector.Vector{intersectOne, intersectTwo}
}

func Distance(a, b vector.Vector) float64 {
	return b.Sub(a).Magnitude()
}
