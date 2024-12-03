package mesh

import (
	"Application/entity"
	"Application/transform"
	"Application/vector"
)

type Mesh struct {
	GameObject          entity.Entity
	Transform           transform.Transform
	Vertices            []vector.Vec3
	Indicies            []uint32
	TransformedVertices []vector.Vec3
}

var DEFAULT_GIZMO Mesh = Mesh{
	Transform: transform.DEFAULT,
	Vertices: []vector.Vec3{
		{X: -0.03, Y: -0.05, Z: 0.0},
		{X: 0.0, Y: 0.05, Z: 0.0},
		{X: 0.03, Y: -0.05, Z: 0.0},
	},
	Indicies: []uint32{0, 1, 2},
	TransformedVertices: []vector.Vec3{
		{X: -0.03, Y: -0.05, Z: 0.0},
		{X: 0.0, Y: 0.05, Z: 0.0},
		{X: 0.03, Y: -0.05, Z: 0.0},
	},
}

var DEFAULT_TRIANGLE Mesh = Mesh{
	Transform: transform.DEFAULT,
	Vertices: []vector.Vec3{
		{X: -0.3, Y: -0.5, Z: 0.0},
		{X: 0.0, Y: 0.5, Z: 0.0},
		{X: 0.3, Y: -0.5, Z: 0.0},
	},
	Indicies: []uint32{0, 1, 2},
	TransformedVertices: []vector.Vec3{
		{X: -0.3, Y: -0.5, Z: 0.0},
		{X: 0.0, Y: 0.5, Z: 0.0},
		{X: 0.3, Y: -0.5, Z: 0.0},
	},
}

func TransformVertices(m *Mesh) {
	m.TransformedVertices = make([]vector.Vec3, len(m.Vertices))
	for i, v := range m.Vertices {
		m.TransformedVertices[i] = vector.Add(v, m.Transform.Position)
	}
}

// Returns closest vertex point of  mesh (m) vertex to point (p)
func GetClosestVertex(m Mesh, p vector.Vec3) (closestVertex vector.Vec3) {
	minDist := float64(1<<64 - 1)
	for _, v := range m.TransformedVertices {
		newDist := vector.Distance(v, p)
		if newDist < minDist {
			minDist = newDist
			closestVertex = v
		}
	}
	return
}

func Get(m Mesh) Mesh {
	return m
}
