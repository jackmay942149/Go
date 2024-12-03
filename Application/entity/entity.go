package entity

import (
	"Application/component"
	"Application/transform"
)

type Entity struct {
	Transform  transform.Transform
	Components map[string]component.Component
}

func (e *Entity) AddComponent(c component.Component) {
	e.Components[c.Name()] = c
}

func (e *Entity) GetComponent(name string) component.Component {
	return e.Components[name]
}
