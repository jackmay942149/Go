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

func AttachWindow(window glfw.Window) {
	WINDOW = window

	ESC = key(glfw.KeyEscape)
	KEY1 = key(glfw.Key1)
	KEY2 = key(glfw.Key2)
	K = key(glfw.KeyK)

}

func Refresh() {
	RefreshKey(&ESC)
	RefreshKey(&KEY1)
	RefreshKey(&KEY2)
	RefreshKey(&K)
}

func RefreshKey(k *Key) {
	if WINDOW.GetKey(k.Code) == glfw.Press {
		UpdateKey(k, true)
	} else {
		UpdateKey(k, false)
	}
}

// Update key should
/* 	Trigger Pressed Once
Trigger Held Always
Trigger Resleased Once
5 moments (No Press, First Press, Hold, Release, No Press)
*/

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

func key(glfwKey glfw.Key) Key {
	var k Key
	k.Code = glfwKey
	k.Pressed = false
	k.Held = false
	k.Released = false
	return k
}
