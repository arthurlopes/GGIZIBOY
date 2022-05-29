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
	title       string
	Glfw_window *glfw.Window
}

func WindowFactory(rows int, columns int, scale int, title string, shared_window *glfw.Window) *Window_struct {
	var window = Window_struct{}
	window.Innit(rows, columns, scale, title, shared_window)
	return &window
}

func (window *Window_struct) Innit(rows int, columns int, scale int, title string, shared_window *glfw.Window) {
	window.Scale = scale
	window.Rows = rows
	window.Columns = columns
	window.Width = columns * scale
	window.Height = rows * scale

	window.title = title

	window.Glfw_window = window.initGlfw(shared_window)
}

func (window *Window_struct) initGlfw(shared_window *glfw.Window) *glfw.Window {
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	_window, err := glfw.CreateWindow(window.Width, window.Height, "GG IZI BOY - "+window.title, nil, shared_window)
	if err != nil {
		panic(err)
	}
	// _window.MakeContextCurrent()

	return _window
}
