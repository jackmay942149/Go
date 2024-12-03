package vector

import "math"

type Vec2 struct {
	U float32
	V float32
}

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

type Vec4 struct {
	R float32
	G float32
	B float32
	A float32
}

// Adds two vector3 a and b element-wise
func Add(a Vec3, b Vec3) (ret Vec3) {
	ret.X = a.X + b.X
	ret.Y = a.Y + b.Y
	ret.Z = a.Z + b.Z
	return
}

// Finds the 2D distance between vectors a and b
func Distance(a Vec3, b Vec3) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
}
