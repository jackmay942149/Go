package shaders

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

const VERTEX_SHADER_SRC = `#version 330 core
layout (location = 0) in vec3 aPos;
void main()
{
	gl_Position = vec4(aPos.x, -aPos.y, aPos.z, 1.0);
}`

const FRAGMENT_SHADER_SRC string = `#version 330 core
out vec4 FragColor;
uniform vec4 ourColor;
void main()
{
	FragColor = ourColor;
}`

func MakeShaderProgram(vertFilepath string, fragFilepath string) uint32 {
	// Make Vertex Shader
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	vecSrc := VERTEX_SHADER_SRC
	vertexShaderSourceRef, free := gl.Strs(vecSrc)
	vecLen := int32(len(VERTEX_SHADER_SRC))
	gl.ShaderSource(vertexShader, 1, vertexShaderSourceRef, &vecLen)
	gl.CompileShader(vertexShader)
	free()

	// Test Vertex Shader
	var vertexSuccess int32
	gl.GetShaderiv(vertexShader, gl.COMPILE_STATUS, &vertexSuccess)
	if vertexSuccess == 0 {
		error := make([]uint8, 512)
		var length int32
		gl.GetShaderInfoLog(vertexShader, 512, &length, &error[0])
		fmt.Println("ERROR::SHADER::VERTEX::COMPILATION_FAILED", string(error[:length]))
		fmt.Println(VERTEX_SHADER_SRC)
	}

	// Make Fragment Shader
	fragShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fragShaderSourceRef, free := gl.Strs(FRAGMENT_SHADER_SRC)
	fragLen := int32(len(FRAGMENT_SHADER_SRC))
	gl.ShaderSource(fragShader, 1, fragShaderSourceRef, &fragLen)
	gl.CompileShader(fragShader)
	free()

	// Test Fragment Shader
	var fragSuccess int32
	gl.GetShaderiv(fragShader, gl.COMPILE_STATUS, &fragSuccess)
	if fragSuccess == 0 {
		error := make([]uint8, 512)
		var length int32
		gl.GetShaderInfoLog(fragShader, 512, &length, &error[0])
		fmt.Println("ERROR::SHADER::FRAG::COMPILATION_FAILED", string(error[:length]))
		fmt.Println(FRAGMENT_SHADER_SRC)
	}

	// Create Shader Program
	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragShader)
	gl.LinkProgram(shaderProgram)

	// Test Program Linking
	var linkSuccess int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &linkSuccess)
	if linkSuccess == 0 {
		error := make([]uint8, 512)
		var length int32
		gl.GetProgramInfoLog(shaderProgram, 512, &length, &error[0])
		fmt.Println("ERROR::SHADER::PROGRAM::LINKING_FAILED", string(error[:length]))
	}

	// Cleanup
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragShader)
	return shaderProgram
}
