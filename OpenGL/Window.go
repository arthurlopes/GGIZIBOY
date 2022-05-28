package OpenGL

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window_struct struct {
	Rows        int
	Columns     int
	Width       int
	Height      int
	Scale       int
	Glfw_window *glfw.Window
}

func WindowFactory(rows int, columns int, scale int) *Window_struct {
	var window = Window_struct{}
	window.Innit(rows, columns, scale)
	return &window
}

func (window *Window_struct) Innit(rows int, columns int, scale int) {
	window.Scale = scale
	window.Rows = rows
	window.Columns = columns
	window.Width = columns * scale
	window.Height = rows * scale

	window.Glfw_window = window.initGlfw()
}

func (window *Window_struct) initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	_window, err := glfw.CreateWindow(window.Width, window.Height, "GG IZI BOY", nil, nil)
	if err != nil {
		panic(err)
	}
	_window.MakeContextCurrent()

	return _window
}
