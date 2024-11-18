package mesh

import (
	"Window/transform"
)

type Mesh struct {
	Transform transform.Transform
	Vertices  []float32
	Indicies  []uint32
}

func GetTransformedIndicies(m Mesh) []float32 {
	vertexPositions := m.Vertices
	for i := 0; i < len(m.Vertices)/3; i++ {
		vertexPositions[3*i] = m.Vertices[3*i] + m.Transform.Position.X
		vertexPositions[3*i+1] = m.Vertices[3*i+1] + m.Transform.Position.Y
		vertexPositions[3*i+2] = m.Vertices[3*i+2] + m.Transform.Position.Z
	}
	return vertexPositions
}
