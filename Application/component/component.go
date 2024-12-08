package component

type Component interface {
	Name() string
	Start()
	Update()
}
