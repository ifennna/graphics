package main

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"go-graphics/lines"
	"go-graphics/pixel"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	window := buildWindow()
	defer glfw.Terminate()

	program := initOpenGL()

	for !window.ShouldClose() {
		draw(window, program)
	}
}

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	run()

	window.SwapBuffers() // swap rendering and drawing buffers
	glfw.PollEvents()    // poll input events
}

func run() {
	start := pixel.Point{0, 0}
	end := pixel.Point{19, 9}

	lines.Draw(start, end, "dda")
}
