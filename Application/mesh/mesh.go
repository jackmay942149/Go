package mesh

import (
	"Application/transform"
	"Application/vector"
	"math"
)

type Mesh struct {
	Transform           transform.Transform
	Vertices            []vector.Vec3
	Indicies            []uint32
	TransformedVertices []vector.Vec3
}

func GetTransformedIndicies(m *Mesh) []vector.Vec3 {
	vertexPositions := make([]vector.Vec3, len(m.Vertices)*3)
	for i := 0; i < len(m.Vertices); i++ {
		vertexPositions[i].X = m.Vertices[i].X + m.Transform.Position.X
		vertexPositions[i].Y = m.Vertices[i].Y + m.Transform.Position.Y
		vertexPositions[i].Z = m.Vertices[i].Z + m.Transform.Position.Z
	}
	m.TransformedVertices = vertexPositions
	return vertexPositions
}

func GetClosestVertex(m Mesh, point vector.Vec3) vector.Vec3 {
	var closestVertex vector.Vec3
	minDist := float64(1<<64 - 1)
	for i := 0; i < len(m.Vertices); i++ {
		newDist := math.Abs(float64(point.X - m.TransformedVertices[i].X))
		newDist += math.Abs(float64(point.Y - m.TransformedVertices[i].Y))
		if newDist < minDist {
			minDist = newDist
			closestVertex.X = m.TransformedVertices[i].X
			closestVertex.Y = m.TransformedVertices[i].Y
			closestVertex.Z = m.TransformedVertices[i].Z
		}
	}
	return closestVertex
}
