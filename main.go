package main

import (
	"GGIZIBOY/OpenGL"
	"GGIZIBOY/gameboy"
	"log"
	"runtime"
	"time"

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

func initGlfw() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
}

func check_input(window *glfw.Window, input_channel chan uint8) {
	for {
		state := uint8(0)
		state |= uint8(window.GetKey(glfw.KeyRight))
		state |= uint8(window.GetKey(glfw.KeyLeft)) << 1
		state |= uint8(window.GetKey(glfw.KeyUp)) << 2
		state |= uint8(window.GetKey(glfw.KeyDown)) << 3
		state |= uint8(window.GetKey(glfw.KeyA)) << 4
		state |= uint8(window.GetKey(glfw.KeyS)) << 5
		state |= uint8(window.GetKey(glfw.KeyBackspace)) << 6
		state |= uint8(window.GetKey(glfw.KeyEnter)) << 7
		state = (uint8(state) ^ 0xff)

		select {
		case input_channel <- state:
		default:
		}
	}
}

func main() {
	runtime.LockOSThread()

	initGlfw()

	// bg_window := OpenGL.WindowFactory(256, 256, 4, "Background", nil)
	// bg_window.Glfw_window.MakeContextCurrent()

	screen_window := OpenGL.WindowFactory(144, 160, 4, "Main screen", nil)
	screen_window.Glfw_window.MakeContextCurrent()

	// tile_window := OpenGL.WindowFactory(196, 128, 4, "Tile Data", bg_window.Glfw_window)
	// tile_window.Glfw_window.MakeContextCurrent()
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

	bg_screen := makeScreen(screen_window)
	// tile_screen := makeScreen(tile_window)

	hblank_channel := make(chan bool)
	input_channel := make(chan uint8)

	var gameboy = gameboy.GameboyFactory(hblank_channel, input_channel)
	go gameboy.Run(-1)
	go check_input(screen_window.Glfw_window, input_channel)

	// for !bg_window.Glfw_window.ShouldClose() && !tile_window.Glfw_window.ShouldClose() {
	for !screen_window.Glfw_window.ShouldClose() {
		draw_flag := <-hblank_channel

		ticker := time.NewTicker(1 * time.Millisecond)
		done := make(chan bool)

		if draw_flag {
			select {
			case <-done:
				return
			case <-ticker.C:

				updateScreen(gameboy.GPU.Screen, bg_screen, screen_window)
				// updateScreen(gameboy.GPU.Tile_data, tile_screen, tile_window)

				screen_window.Glfw_window.MakeContextCurrent()
				draw(bg_screen, screen_window.Glfw_window, program)

				// tile_window.Glfw_window.MakeContextCurrent()
				// draw(tile_screen, tile_window.Glfw_window, program)

				glfw.PollEvents()
			}
		}

	}
}
