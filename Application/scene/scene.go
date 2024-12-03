package scene

import (
	"Application/component"
	"Application/entity"
	"Application/mesh"
	"Application/transform"
)

type Scene struct {
	Entities []entity.Entity
}

var DEFAULT Scene = Scene{Entities: []entity.Entity{DEFAULT_GIZMO, DEFAULT_TRIANGLE}}

var DEFAULT_GIZMO entity.Entity = entity.Entity{
	Transform:  transform.DEFAULT,
	Components: []component.Component{&mesh.DEFAULT_GIZMO},
}

var DEFAULT_TRIANGLE entity.Entity = entity.Entity{
	Transform:  transform.DEFAULT,
	Components: []component.Component{&mesh.DEFAULT_TRIANGLE},
}
