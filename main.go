package main

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"go-graphics/arcs"
	"go-graphics/circles"
	"go-graphics/ellipses"
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
	end1 := pixel.Point{19, 9}
	end2 := pixel.Point{9, 19}

	lines.Draw(start, end1, "dda")
	lines.Draw(start, end2, "bresenhams")

	circles.Draw(pixel.Point{0, 0}, 19, "midpoint")
	circles.Draw(pixel.Point{0, 0}, 29, "bresenhams")

	ellipse := pixel.Point{0, 0}
	ellipses.Draw(ellipse, 31, 15)

	arcStart := pixel.Point{0, 70}
	arcEnd := pixel.Point{50, 20}

	arcs.Draw(arcStart, arcEnd)
}
