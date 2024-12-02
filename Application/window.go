package main

import (
	"fmt"
	"runtime"

	"Application/input"
	"Application/mesh"
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

var shadingModel shading.Model = shading.MakeModel(0)
var mousePosx float64
var mousePosy float64

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
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

	var tri1 mesh.Mesh
	tri1.Vertices = []vector.Vec3{
		{X: -0.03, Y: -0.05, Z: 0.0},
		{X: 0.0, Y: 0.05, Z: 0.0},
		{X: 0.03, Y: -0.05, Z: 0.0},
	}

	tri1.Indicies = []uint32{0, 1, 2}
	tri1.Transform = transform.DEFAULT
	tri1.TransformedVertices = mesh.GetTransformedIndicies(&tri1)

	var tri2 mesh.Mesh

	tri2.Vertices = []vector.Vec3{
		{X: 0.3, Y: -0.5, Z: 0.0},
		{X: 0.2, Y: 0.5, Z: 0.0},
		{X: 0.1, Y: -0.5, Z: 0.0},
	}

	tri2.Indicies = []uint32{0, 1, 2}
	tri2.Transform = transform.DEFAULT
	tri2.TransformedVertices = mesh.GetTransformedIndicies(&tri2)

	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	vertexShaderSourceRef, free := gl.Strs(shaders.VERTEX_SHADER_SRC)
	gl.ShaderSource(vertexShader, 1, vertexShaderSourceRef, nil)
	free()
	gl.CompileShader(vertexShader)

	fragShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fragShaderSourceRef, free := gl.Strs(shaders.FRAGMENT_SHADER_SRC)
	gl.ShaderSource(fragShader, 1, fragShaderSourceRef, nil)
	free()
	gl.CompileShader(fragShader)

	shaderProgram := gl.CreateProgram()

	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragShader)
	gl.LinkProgram(shaderProgram)

	var VBO1 uint32
	var VAO1 uint32
	var EBO1 uint32

	var VBO2 uint32
	var VAO2 uint32
	var EBO2 uint32

	gl.GenVertexArrays(1, &VAO1)
	gl.GenBuffers(1, &VBO1)
	gl.GenBuffers(1, &EBO1)

	gl.GenVertexArrays(1, &VAO2)
	gl.GenBuffers(1, &VBO2)
	gl.GenBuffers(1, &EBO2)

	gl.BindVertexArray(VAO1)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO1)
	gl.BufferData(gl.ARRAY_BUFFER, len(tri1.Vertices)*12, gl.Ptr(tri1.TransformedVertices), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO1)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(tri1.Indicies)*12, gl.Ptr(tri1.Indicies), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	gl.BindVertexArray(VAO2)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO2)
	gl.BufferData(gl.ARRAY_BUFFER, len(tri2.Vertices)*12, gl.Ptr(tri2.TransformedVertices), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO2)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(tri2.Indicies)*12, gl.Ptr(tri2.Indicies), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	for !window.ShouldClose() {
		// Close window on escape press
		ProcessInput(*window)

		if input.D.Held {
			tri2.Transform.Position.X += 0.001
		} else if input.A.Held {
			tri2.Transform.Position.X -= 0.001
		} else if input.W.Held {
			tri2.Transform.Position.Y += 0.001
		} else if input.S.Held {
			tri2.Transform.Position.Y -= 0.001
		}

		if input.V.Pressed {
			mousePosx, mousePosy = window.GetCursorPos()
			var clickPos vector.Vec3
			clickPos.X = 2*float32(mousePosx/float64(WINDOWWIDTH)) - 1
			clickPos.Y = -2*float32(mousePosy/float64(WINDOWHEIGHT)) + 1
			clickPos.Z = 1
			fmt.Println("ClickPos = ", clickPos)
			tri1.Transform.Position = mesh.GetClosestVertex(tri2, clickPos)
			fmt.Println("Transform is {", tri1.Transform.Position.X, ", ", tri1.Transform.Position.Y, "}")
		}

		if shadingModel.Model == shading.WIREFRAME {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
		} else {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
		}

		mesh.GetTransformedIndicies(&tri1)
		mesh.GetTransformedIndicies(&tri2)

		gl.BindBuffer(gl.ARRAY_BUFFER, VBO1)
		gl.BufferData(gl.ARRAY_BUFFER, len(tri1.Vertices)*12, gl.Ptr(tri1.TransformedVertices), gl.STATIC_DRAW)
		gl.BindBuffer(gl.ARRAY_BUFFER, VBO2)
		gl.BufferData(gl.ARRAY_BUFFER, len(tri2.Vertices)*12, gl.Ptr(tri2.TransformedVertices), gl.STATIC_DRAW)

		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.UseProgram(shaderProgram)

		gl.BindVertexArray(VAO1)
		gl.DrawElements(gl.TRIANGLES, int32(len(tri1.Indicies)), gl.UNSIGNED_INT, nil)

		gl.BindVertexArray(VAO2)
		gl.DrawElements(gl.TRIANGLES, int32(len(tri2.Indicies)), gl.UNSIGNED_INT, nil)

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
