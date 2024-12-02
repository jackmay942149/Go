package shaders

const VERTEX_SHADER_SRC string = `#version 330 core\n
    layout (location = 0) in vec3 aPos;\n
    void main()\n
    {\n
       gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);\n
    }\0`

const FRAGMENT_SHADER_SRC string = `#version 330 core\n
	out vec4 FragColor;\n
	void main()\n
	{\n
		FragColor = vec4(1.0f, 1.0f, 1.0f, 1.0f);\n
	}\0`
