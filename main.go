package main

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
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

	pixel.DrawPixel(0, 0)

	window.SwapBuffers() // swap rendering and drawing buffers
	glfw.PollEvents()    // poll input events
}
