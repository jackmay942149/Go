package transform

import (
	"Application/vector"
)

type Transform struct {
	Position vector.Vec3
	Rotation vector.Vec3
	Scale    vector.Vec3
}

var DEFAULT Transform = Transform{
	vector.Vec3{X: 0, Y: 0, Z: 0},
	vector.Vec3{X: 0, Y: 0, Z: 0},
	vector.Vec3{X: 1, Y: 1, Z: 1},
}
