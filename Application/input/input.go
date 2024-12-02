package input

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Key struct {
	Code glfw.Key
	// Above can be changed to a bitset
	Pressed  bool
	Held     bool
	Released bool
}

var WINDOW glfw.Window

var ESC Key
var KEY1 Key
var KEY2 Key
var K Key
var W Key
var A Key
var S Key
var D Key
var V Key

func AttachWindow(window glfw.Window) {
	WINDOW = window

	ESC.Code = glfw.KeyEscape
	KEY1.Code = glfw.Key1
	KEY2.Code = glfw.Key2
	K.Code = glfw.KeyK
	W.Code = glfw.KeyW
	A.Code = glfw.KeyA
	S.Code = glfw.KeyS
	D.Code = glfw.KeyD
	V.Code = glfw.KeyV
}

func Refresh() {
	RefreshKey(&ESC)
	RefreshKey(&KEY1)
	RefreshKey(&KEY2)
	RefreshKey(&K)
	RefreshKey(&W)
	RefreshKey(&A)
	RefreshKey(&S)
	RefreshKey(&D)
	RefreshKey(&V)
}

func RefreshKey(k *Key) {
	UpdateKey(k, WINDOW.GetKey(k.Code) == glfw.Press)
}

func UpdateKey(k *Key, b bool) {
	if !b {
		if k.Released { // Released only fires once
			k.Released = false
		}
		if k.Held { // Released Fires
			k.Released = true
		}
		k.Pressed = false
		k.Held = false
	} else {
		if k.Pressed && !k.Held {
			k.Pressed = false // Pressed only fires once
			k.Held = true     // Hold Fires
		} else if !k.Held {
			k.Pressed = true // Pressed Fires
		}
		k.Released = false
	}
}
