package main

import (
	"GGIZIBOY/OpenGL"
	"GGIZIBOY/gameboy"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
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
	points []float32
	vao    uint32
	vbo    uint32
}

func draw(cur_screen *screen, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindBuffer(gl.ARRAY_BUFFER, cur_screen.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(cur_screen.points), gl.Ptr(cur_screen.points), gl.DYNAMIC_DRAW)

	gl.BindVertexArray(cur_screen.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(cur_screen.points)/4))

	glfw.PollEvents()
	window.SwapBuffers()
}

func updateScreen(new_screen [][]uint8, cur_screen *screen, window *OpenGL.Window_struct) {
	for y := 0; y < window.Rows; y++ {
		for x := 0; x < window.Columns; x++ {
			color := new_screen[y][x]
			for p := 3; p < len(square); p += 4 {
				cur_screen.points[(y*window.Columns*len(square))+(x*len(square))+p] = 1 - float32(color)/3
			}
		}
	}
}

func makeScreen(window *OpenGL.Window_struct) *screen {
	points := make([][][]float32, window.Rows)
	points_flat := make([]float32, window.Rows*window.Columns*len(square))
	for y := 0; y < window.Rows; y++ {
		for x := 0; x < window.Columns; x++ {
			c := newPixel(x, y, window)
			points[y] = append(points[y], c)
			for p := 0; p < len(square); p++ {
				points_flat[(y*window.Columns*len(square))+(x*len(square))+p] = points[y][x][p]
			}
		}
	}

	vao, vbo := makeVao(points_flat)
	return &screen{
		points_flat,
		vao,
		vbo,
	}
}

func newPixel(x, y int, window *OpenGL.Window_struct) []float32 {
	points := make([]float32, len(square))
	copy(points, square)

	for i := 0; i < len(points); i++ {
		var position float32
		var size float32
		switch i % 4 {
		case 0:
			size = 1.0 / float32(window.Columns)
			position = float32(x) * size
		case 1:
			size = 1.0 / float32(window.Rows)
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

	return points
}

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

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	return vao, vbo
}

func initOpenGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

}

func main() {
	runtime.LockOSThread()

	window := OpenGL.WindowFactory(256, 256, 4)
	defer glfw.Terminate()
	initOpenGL()

	vertexShader := OpenGL.ShaderFactory("./OpenGL/shaders/vertex.glsl", gl.VERTEX_SHADER)
	vertexShader.CompileShader()
	fragmentShader := OpenGL.ShaderFactory("./OpenGL/shaders/fragment.glsl", gl.FRAGMENT_SHADER)
	fragmentShader.CompileShader()

	gl.CreateProgram()
	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader.Compiled_shader)
	gl.AttachShader(program, fragmentShader.Compiled_shader)
	gl.LinkProgram(program)

	hblank_channel := make(chan bool)

	var gameboy = gameboy.GameboyFactory(hblank_channel)
	go gameboy.Run(-1)

	_screen := makeScreen(window)
	draw(_screen, window.Glfw_window, program)

	for !window.Glfw_window.ShouldClose() {
		draw_flag := <-hblank_channel
		if draw_flag {
			updateScreen(gameboy.GPU.Background, _screen, window)
			draw(_screen, window.Glfw_window, program)
		}
	}
}
