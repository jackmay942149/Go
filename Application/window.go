package main

import (
	"fmt"
	"math"
	"runtime"

	"Application/components/wasdmove"
	"Application/input"
	"Application/mesh"
	"Application/scene"
	"Application/shaders"
	"Application/shading"
	"Application/transform"
	"Application/vector"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	WINDOWWIDTH  int32 = 800
	WINDOWHEIGHT int32 = 600
)

var window *glfw.Window
var currentscene scene.Scene
var shadingModel shading.Model = shading.Model{Model: shading.UNLIT}
var mousePosx float64
var mousePosy float64

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	// Start glfw
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(int(WINDOWWIDTH), int(WINDOWHEIGHT), "Triangle", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	input.AttachWindow(*window)

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	// Make Scene
	currentscene = scene.DEFAULT
	currentscene.Entities[0].AddComponent(&transform.DEFAULT)
	currentscene.Entities[1].AddComponent(&transform.DEFAULT)

	meshGizmo := mesh.Mesh{
		Entity:    &currentscene.Entities[0],
		Transform: nil,
		Vertices: []vector.Vec3{
			{X: -0.03, Y: -0.05, Z: 0.0},
			{X: 0.0, Y: 0.05, Z: 0.0},
			{X: 0.03, Y: -0.05, Z: 0.0},
		},
		Indicies: []uint32{0, 1, 2},
		TransformedVertices: []vector.Vec3{
			{X: -0.03, Y: -0.05, Z: 0.0},
			{X: 0.0, Y: 0.05, Z: 0.0},
			{X: 0.03, Y: -0.05, Z: 0.0},
		},
	}
	meshTriangle := mesh.Mesh{
		Entity:    &currentscene.Entities[1],
		Transform: nil,
		Vertices: []vector.Vec3{
			{X: -0.3, Y: -0.5, Z: 0.0},
			{X: 0.0, Y: 0.5, Z: 0.0},
			{X: 0.3, Y: -0.5, Z: 0.0},
		},
		Indicies: []uint32{0, 1, 2},
		TransformedVertices: []vector.Vec3{
			{X: -0.3, Y: -0.5, Z: 0.0},
			{X: 0.0, Y: 0.5, Z: 0.0},
			{X: 0.3, Y: -0.5, Z: 0.0},
		},
	}
	currentscene.Entities[0].AddComponent(&meshGizmo)
	currentscene.Entities[1].AddComponent(&meshTriangle)

	movement := wasdmove.Wasdmove{Entity: &currentscene.Entities[0]}
	currentscene.Entities[0].AddComponent(&movement)

	shaderProgram := shaders.MakeShaderProgram()

	var VAO []uint32 = make([]uint32, len(currentscene.Entities))
	var VBO []uint32 = make([]uint32, len(currentscene.Entities))
	var EBO []uint32 = make([]uint32, len(currentscene.Entities))

	for i := range VAO {
		gl.GenVertexArrays(1, &VAO[i])
		gl.GenBuffers(1, &VBO[i])
		gl.GenBuffers(1, &EBO[i])
	}

	for i := range VAO {
		meshToDraw := currentscene.Entities[i].GetComponent("Mesh").(*mesh.Mesh)
		gl.BindVertexArray(VAO[i])
		gl.BindBuffer(gl.ARRAY_BUFFER, VBO[i])
		gl.BufferData(gl.ARRAY_BUFFER, len(meshToDraw.Vertices)*12, gl.Ptr(meshToDraw.TransformedVertices), gl.DYNAMIC_DRAW)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO[i])
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(meshToDraw.Indicies)*12, gl.Ptr(meshToDraw.Indicies), gl.DYNAMIC_DRAW)

		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
		gl.EnableVertexAttribArray(0)
	}

	for _, e := range currentscene.Entities {
		for _, c := range e.Components {
			c.Start()
		}
	}

	var maxVertexsAbs int32
	gl.GetIntegerv(gl.MAX_VERTEX_ATTRIBS, &maxVertexsAbs)
	fmt.Println(maxVertexsAbs)

	for !window.ShouldClose() {
		// Close window on escape press
		ProcessInput(*window)

		tri1 := currentscene.Entities[0].GetComponent("Mesh").(*mesh.Mesh)
		tri2 := currentscene.Entities[1].GetComponent("Mesh").(*mesh.Mesh)

		timeValue := glfw.GetTime()
		greenValue := float32((math.Sin(timeValue) / 2.0) + 0.5)
		vertexColorLocation := gl.GetUniformLocation(shaderProgram, gl.Str("ourColor\x00"))
		gl.UseProgram(shaderProgram)
		gl.Uniform4f(vertexColorLocation, 0.0, greenValue, 0.0, 1.0)

		/*
			if input.D.Held {
				tri2.Transform.Position.X += 0.001
			} else if input.A.Held {
				tri2.Transform.Position.X -= 0.001
			} else if input.W.Held {
				tri2.Transform.Position.Y += 0.001
			} else if input.S.Held {
				tri2.Transform.Position.Y -= 0.001
			}
		*/

		for _, e := range currentscene.Entities {
			for _, c := range e.Components {
				c.Update()
			}
		}
		if input.V.Pressed {
			mousePosx, mousePosy = window.GetCursorPos()
			var clickPos vector.Vec3
			clickPos.X = 2*float32(mousePosx/float64(WINDOWWIDTH)) - 1
			clickPos.Y = -2*float32(mousePosy/float64(WINDOWHEIGHT)) + 1
			clickPos.Z = 1
			fmt.Println("ClickPos = ", clickPos)
			tri1.Transform.Position = mesh.GetClosestVertex(*tri2, clickPos)
			fmt.Println("Transform is {", tri1.Transform.Position.X, ", ", tri1.Transform.Position.Y, "}")
		}

		if shadingModel.Model == shading.WIREFRAME {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
		} else {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
		}

		mesh.TransformVertices(tri1)
		mesh.TransformVertices(tri2)

		for i := range VAO {
			meshToDraw := currentscene.Entities[i].GetComponent("Mesh").(*mesh.Mesh)
			gl.BindBuffer(gl.ARRAY_BUFFER, VBO[i])
			gl.BufferData(gl.ARRAY_BUFFER, len(meshToDraw.Vertices)*12, gl.Ptr(meshToDraw.TransformedVertices), gl.STATIC_DRAW)
		}

		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.UseProgram(shaderProgram)

		for i := range VAO {
			meshToDraw := currentscene.Entities[i].GetComponent("Mesh").(*mesh.Mesh)
			gl.BindVertexArray(VAO[i])
			gl.DrawElements(gl.TRIANGLES, int32(len(meshToDraw.Indicies)), gl.UNSIGNED_INT, nil)
		}

		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func ProcessInput(window glfw.Window) {
	input.Refresh()

	// Close window on escape press
	if input.ESC.Released {
		window.SetShouldClose(true)
	}

	if input.KEY1.Pressed {
		shadingModel.Model = shading.WIREFRAME
	} else if input.KEY2.Pressed {
		shadingModel.Model = shading.UNLIT
	}
}
