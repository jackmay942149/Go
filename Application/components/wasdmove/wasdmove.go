package wasdmove

import (
	"Application/entity"
	"Application/input"
	"Application/transform"
)

type Wasdmove struct {
	Entity *entity.Entity

	transform *transform.Transform
}

func (this *Wasdmove) Name() string {
	return "Wasdmove"
}

func (this *Wasdmove) Start() {
	this.transform = this.Entity.GetComponent("Transform").(*transform.Transform)
}

func (this *Wasdmove) Update() {
	if input.D.Held {
		this.transform.Position.X += 0.001
	} else if input.A.Held {
		this.transform.Position.X -= 0.001
	} else if input.W.Held {
		this.transform.Position.Y += 0.001
	} else if input.S.Held {
		this.transform.Position.Y -= 0.001
	}

}
