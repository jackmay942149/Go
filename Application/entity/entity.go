package entity

import (
	"Application/component"
	"Application/transform"
)

type Entity struct {
	Transform  transform.Transform
	Components []component.Component
}
