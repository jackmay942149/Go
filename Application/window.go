package main

import (
	"fmt"
	"runtime"

	"Application/mesh"
	"Application/shading"
	"Application/transform"

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

const vertexShaderSource string = `#version 330 core\n
    layout (location = 0) in vec3 aPos;\n
    void main()\n
    {\n
       gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);\n
    }\0`

const fragmentShaderSource string = `#version 330 core\n
	out vec4 FragColor;\n
	void main()\n
	{\n
		FragColor = vec4(1.0f, 1.0f, 1.0f, 1.0f);\n
	}\0`

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

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	var tri1 mesh.Mesh
	var tri2 mesh.Mesh

	tri1.Vertices = []float32{-0.03, -0.05, 0.0,
		-0.02, 0.05, 0.0,
		-0.01, -0.05, 0.0}

	tri1.Indicies = []uint32{0, 1, 2}
	tri1.Transform = transform.MakeTransform()

	tri2.Vertices = []float32{0.3, -0.5, 0.0,
		0.2, 0.5, 0.0,
		0.1, -0.5, 0.0}

	tri2.Indicies = []uint32{0, 1, 2}
	tri2.Transform = transform.MakeTransform()

	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	vertexShaderSourceRef, free := gl.Strs(vertexShaderSource)
	gl.ShaderSource(vertexShader, 1, vertexShaderSourceRef, nil)
	free()
	gl.CompileShader(vertexShader)

	fragShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fragShaderSourceRef, free := gl.Strs(fragmentShaderSource)
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
	gl.BufferData(gl.ARRAY_BUFFER, len(tri1.Vertices)*4, gl.Ptr(mesh.GetTransformedIndicies(tri1)), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO1)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(tri1.Indicies)*4, gl.Ptr(tri1.Indicies), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	gl.BindVertexArray(VAO2)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO2)
	gl.BufferData(gl.ARRAY_BUFFER, len(tri2.Vertices)*4, gl.Ptr(tri2.Vertices), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO2)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(tri2.Indicies)*4, gl.Ptr(tri2.Indicies), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	for !window.ShouldClose() {
		// Close window on escape press
		ProcessInput(*window)

		if window.GetKey(glfw.KeyD) == glfw.Press {
			tri1.Transform.Position.X += 0.001
		} else if window.GetKey(glfw.KeyA) == glfw.Press {
			tri1.Transform.Position.X -= 0.001
		} else if window.GetKey(glfw.KeyW) == glfw.Press {
			tri1.Transform.Position.Y += 0.001
		} else if window.GetKey(glfw.KeyS) == glfw.Press {
			tri1.Transform.Position.Y -= 0.001
		}

		if window.GetKey(glfw.KeyV) == glfw.Press {
			mousePosx, mousePosy = window.GetCursorPos()
			tri1.Transform.Position.X = 2*float32(mousePosx/float64(WINDOWWIDTH)) - 1
			tri1.Transform.Position.Y = -2*float32(mousePosy/float64(WINDOWHEIGHT)) + 1
			fmt.Println(mousePosx, mousePosy)
		}

		gl.BindVertexArray(VAO1)
		gl.BindBuffer(gl.ARRAY_BUFFER, VBO1)
		gl.BufferData(gl.ARRAY_BUFFER, len(tri1.Vertices)*4, gl.Ptr(mesh.GetTransformedIndicies(tri1)), gl.STATIC_DRAW)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO1)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(tri1.Indicies)*4, gl.Ptr(tri1.Indicies), gl.STATIC_DRAW)

		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
		gl.EnableVertexAttribArray(0)

		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		if shadingModel.Model == shading.WIREFRAME {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
		} else {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
		}

		gl.UseProgram(shaderProgram)
		gl.BindVertexArray(VAO1)
		gl.DrawElements(gl.TRIANGLES, int32(len(tri1.Indicies)), gl.UNSIGNED_INT, nil)
		gl.BindVertexArray(0)

		gl.BindVertexArray(VAO2)
		//gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangleVert)))
		gl.DrawElements(gl.TRIANGLES, int32(len(tri2.Indicies)), gl.UNSIGNED_INT, nil)
		gl.BindVertexArray(0)

		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func ProcessInput(window glfw.Window) {
	// Close window on escape press
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}

	if window.GetKey(glfw.Key1) == glfw.Press {
		shadingModel.Model = shading.WIREFRAME
	} else if window.GetKey(glfw.Key2) == glfw.Press {
		shadingModel.Model = shading.UNLIT
	}
}
