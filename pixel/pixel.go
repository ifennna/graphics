package pixel

import (
	"github.com/go-gl/gl/all-core/gl"
)

type Point struct {
	X float32
	Y float32
}

func DrawPixel(point Point) {
	x := point.X / 100
	y := point.Y / 100

	gl.PointSize(5)
	gl.Begin(gl.POINTS)
	gl.Color3f(1.0, .5, 1.0)
	gl.Vertex2f(x, y)
	gl.End()
}
