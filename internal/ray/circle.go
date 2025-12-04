package ray

import (
	"math"

	"github.com/quartercastle/vector"
	"github.com/sabaruto/simulator-test/internal/common"
)

func RecieveRayToCircle(rayPosition vector.Vector, rayDirection vector.Vector, position *vector.Vector, radius float64, colour string) (*vector.Vector, *string) {
	rayDirection = rayDirection.Unit()
	intersections := common.GetCircleIntersections(rayPosition, rayDirection, *position, radius)

	// Check an intersection point is found
	if intersections == nil {
		return nil, nil
	}

	var closestPoint *vector.Vector
	smallestDistance := math.Inf(1)
	returnColour := colour

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

	return closestPoint, &returnColour

}
