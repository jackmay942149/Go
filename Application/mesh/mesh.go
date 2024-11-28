package mesh

import (
	"Application/transform"
	"math"
)

type Mesh struct {
	Transform           transform.Transform
	Vertices            []float32
	Indicies            []uint32
	TransformedVertices []float32
}

func GetTransformedIndicies(m *Mesh) []float32 {
	vertexPositions := make([]float32, len(m.Vertices))
	for i := 0; i < len(m.Vertices)/3; i++ {
		vertexPositions[3*i] = m.Vertices[3*i] + m.Transform.Position.X
		vertexPositions[3*i+1] = m.Vertices[3*i+1] + m.Transform.Position.Y
		vertexPositions[3*i+2] = m.Vertices[3*i+2] + m.Transform.Position.Z
	}
	m.TransformedVertices = vertexPositions
	return vertexPositions
}

func GetClosestVertex(m Mesh, point transform.Position) transform.Position {
	var closestVertex transform.Position
	minDist := float64(1<<64 - 1)
	for i := 0; i < len(m.Vertices)/3; i++ {
		newDist := math.Abs(float64(point.X - m.TransformedVertices[3*i]))
		newDist += math.Abs(float64(point.Y - m.TransformedVertices[3*i+1]))
		if newDist < minDist {
			minDist = newDist
			closestVertex.X = m.TransformedVertices[3*i]
			closestVertex.Y = m.TransformedVertices[3*i+1]
			closestVertex.Z = m.TransformedVertices[3*i+2]
		}
	}
	return closestVertex
}
