package scene

import (
	"Application/component"
	"Application/entity"
)

type Scene struct {
	Entities []entity.Entity
}

var DEFAULT Scene = Scene{Entities: []entity.Entity{DEFAULT_GIZMO, DEFAULT_TRIANGLE}}

var DEFAULT_GIZMO entity.Entity = entity.Entity{
	Components: make(map[string]component.Component),
}

var DEFAULT_TRIANGLE entity.Entity = entity.Entity{
	Components: make(map[string]component.Component),
}
