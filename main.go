package main

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"graphics/arcs"
	"graphics/ellipses"
	"graphics/lines"
	"graphics/pixel"
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
	//circles.Draw(pixel.Point{0, 0}, 19, "midpoint")
	//circles.Draw(pixel.Point{10, 0}, 29, "bresenhams")

	lines.Draw(pixel.Point{X: 0, Y: 20}, pixel.Point{X: 50, Y: 20}, "bresenhams")
	lines.Draw(pixel.Point{X: 0, Y: 0}, pixel.Point{X: 50, Y: 0}, "dda")

	ellipses.Draw(pixel.Point{X: 0, Y: 10}, 3, 10)
	ellipses.Draw(pixel.Point{X: 50, Y: 10}, 3, 10)

	arcs.Draw(pixel.Point{X: 25, Y: 45}, 35, true)
	arcs.Draw(pixel.Point{X: 25, Y: -25}, 35, false)
}
