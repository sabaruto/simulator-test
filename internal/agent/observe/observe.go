package observe

import (
	"github.com/quartercastle/vector"
	"github.com/sabaruto/simulator-test/internal/common"
)

type Observe struct {
	objectPosition *vector.Vector
	objectVelocity *vector.Vector
	objectID       *int64
	objectManager  common.ObjectManager
}
