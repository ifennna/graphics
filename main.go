package main

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"graphics/circles"
	"graphics/pixel"
	"runtime"
)

var position = 0
const distance = 60

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

	midpoint := pixel.Point{float64(position), 0}

	if position > distance {
		position = 0
	}

	run(midpoint)

	position++

	window.SwapBuffers() // swap rendering and drawing buffers
	glfw.PollEvents()    // poll input events
}

func run(midpoint pixel.Point) {
	circles.Draw(midpoint, 19, "midpoint")
	//circles.Draw(pixel.Point{10, 0}, 29, "bresenhams")

	//lines.Draw(pixel.Point{X: 0, Y: 20}, pixel.Point{X: 50, Y: 20}, "bresenhams")
	//lines.Draw(pixel.Point{X: 0, Y: 0}, pixel.Point{X: 50, Y: 0}, "dda")
	//
	//ellipses.Draw(pixel.Point{X: 0, Y: 10}, 3, 10)
	//ellipses.Draw(pixel.Point{X: 50, Y: 10}, 3, 10)
	//
	//arcs.Draw(pixel.Point{X: 25, Y: 45}, 35, true)
	//arcs.Draw(pixel.Point{X: 25, Y: -25}, 35, false)
}
