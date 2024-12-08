package mesh

import (
	"Application/entity"
	"Application/transform"
	"Application/vector"
)

type Mesh struct {
	Entity              *entity.Entity
	Transform           *transform.Transform
	Vertices            []vector.Vec3
	Indicies            []uint32
	TransformedVertices []vector.Vec3
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

func (m *Mesh) Name() string {
	return "Mesh"
}
func (m *Mesh) Start() {
	m.Transform = m.Entity.GetComponent("Transform").(*transform.Transform)
}
func (m *Mesh) Update() {
}
