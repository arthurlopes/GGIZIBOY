package main

import (
	"GGIZIBOY/gameboy"
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	rows    = 256
	columns = 256

	width  = columns * 4
	height = rows * 4

	vertexShaderSource = `
		#version 460
		in vec3 vp;
		in float cp;
		out float c;
		void main() {
			gl_Position = vec4(vp, 1.0);
			c = cp;
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 460
		out vec4 frag_colour;
		in float c;
		void main() {
			frag_colour = vec4(c, c, c, 1.0);
		}
	` + "\x00"
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

type screen struct {
	points [][][]float32
	vaos   [][]uint32
	vbos   [][]uint32
}

func draw(cur_screen *screen, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for y := 0; y < rows; y++ {
		for x := 0; x < columns; x++ {
			// vao := makeVao(cur_screen.points[y][x])
			gl.BindBuffer(gl.ARRAY_BUFFER, cur_screen.vbos[y][x])
			gl.BufferData(gl.ARRAY_BUFFER, 4*len(cur_screen.points[y][x]), gl.Ptr(cur_screen.points[y][x]), gl.STATIC_DRAW)

			gl.BindVertexArray(cur_screen.vaos[y][x])
			// gl.BindVertexArray(vao)
			gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/4))
		}
	}

	glfw.PollEvents()
	window.SwapBuffers()
}

func updateScreen(new_screen [][]uint8, cur_screen *screen) {
	for y := 0; y < rows; y++ {
		for x := 0; x < columns; x++ {
			color := new_screen[y][x]
			cur_screen.points[y][x][3] = 1 - float32(color)/3
			cur_screen.points[y][x][7] = 1 - float32(color)/3
			cur_screen.points[y][x][11] = 1 - float32(color)/3
			cur_screen.points[y][x][15] = 1 - float32(color)/3
			cur_screen.points[y][x][19] = 1 - float32(color)/3
			cur_screen.points[y][x][23] = 1 - float32(color)/3
		}
	}
}

func makeScreen() *screen {
	points := make([][][]float32, rows)
	vaos := make([][]uint32, rows)
	vbos := make([][]uint32, rows)
	for y := 0; y < rows; y++ {
		for x := 0; x < columns; x++ {
			c, vao, vbo := newPixel(x, y)
			points[y] = append(points[y], c)
			vaos[y] = append(vaos[y], vao)
			vbos[y] = append(vbos[y], vbo)
		}
	}

	return &screen{
		points,
		vaos,
		vbos,
	}
}

func newPixel(x, y int) ([]float32, uint32, uint32) {
	points := make([]float32, len(square))
	copy(points, square)

	for i := 0; i < len(points); i++ {
		var position float32
		var size float32
		switch i % 4 {
		case 0:
			size = 1.0 / float32(columns)
			position = float32(x) * size
		case 1:
			size = 1.0 / float32(rows)
			position = float32(y) * size
		default:
			continue
		}

		if points[i] < 0 {
			points[i] = (position * 2) - 1
		} else {
			points[i] = ((position + size) * 2) - 1
		}

		if i%4 == 1 {
			points[i] *= -1
		}
	}

	vaos, vbos := makeVao(points)
	return points, vaos, vbos
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) (uint32, uint32) {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	// position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 4*4, nil)
	gl.EnableVertexAttribArray(0)
	// color
	gl.VertexAttribPointer(1, 1, gl.FLOAT, false, 4*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	// gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	return vao, vbo
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "GG IZI BOY", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	hblank_channel := make(chan bool)

	var gameboy = gameboy.GameboyFactory(hblank_channel)
	go gameboy.Run(-1)

	_screen := makeScreen()
	draw(_screen, window, program)

	for !window.ShouldClose() {
		draw_flag := <-hblank_channel
		if draw_flag {
			updateScreen(gameboy.GPU.Background, _screen)
			draw(_screen, window, program)
		}
	}
}
