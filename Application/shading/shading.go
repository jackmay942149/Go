package shading

const (
	WIREFRAME = int16(0)
	UNLIT     = int16(1)
)

type Model struct {
	Model int16
}

func MakeModel(m int16) Model {
	return Model{m}
}
