package entity

import (
	"Application/component"
)

type Entity struct {
	Components map[string]component.Component
}

func (e *Entity) AddComponent(c component.Component) {
	e.Components[c.Name()] = c
}

func (e *Entity) GetComponent(name string) component.Component {
	return e.Components[name]
}
