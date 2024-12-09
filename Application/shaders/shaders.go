package shaders

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
