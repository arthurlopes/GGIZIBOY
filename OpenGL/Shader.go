package OpenGL

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
)

var (
	square = []float32{
		-0.5, 0.5, 0, 1,
		-0.5, -0.5, 0, 1,
		0.5, -0.5, 0, 1,

		-0.5, 0.5, 0, 1,
		0.5, 0.5, 0, 1,
		0.5, -0.5, 0, 1,
	}
)

type Shader_struct struct {
	code            string
	shaderType      uint32
	Compiled_shader uint32
}

func ShaderFactory(path string, shaderType uint32) *Shader_struct {
	var shader = Shader_struct{}
	shader.Innit(path, shaderType)
	return &shader
}

func (shader *Shader_struct) Innit(path string, shaderType uint32) {
	shader.Load_shader_from_file(path)
	shader.shaderType = shaderType
}

func (shader *Shader_struct) Load_shader_from_file(path string) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
	shader.code = string(content) + "\x00"
}

func (shader *Shader_struct) CompileShader() (uint32, error) {
	_shader := gl.CreateShader(shader.shaderType)

	csources, free := gl.Strs(shader.code)
	gl.ShaderSource(_shader, 1, csources, nil)
	free()
	gl.CompileShader(_shader)

	var status int32
	gl.GetShaderiv(_shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(_shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(_shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", shader.code, log)
	}

	shader.Compiled_shader = _shader

	return _shader, nil
}
