package transform

type Transform struct {
	Position Position
	Rotation Rotation
	Scale    Scale
}

type Position struct {
	X float32
	Y float32
	Z float32
}

type Rotation struct {
	Pitch float32
	Yaw   float32
	Roll  float32
}

type Scale struct {
	X float32
	Y float32
	Z float32
}

func MakeTransform() Transform {
	var newTransform Transform
	newTransform.Position.X = 0
	newTransform.Position.Y = 0
	newTransform.Position.Z = 0
	newTransform.Rotation.Pitch = 0
	newTransform.Rotation.Yaw = 0
	newTransform.Rotation.Roll = 0
	newTransform.Scale.X = 1
	newTransform.Scale.Y = 1
	newTransform.Scale.Z = 1
	return newTransform
}
