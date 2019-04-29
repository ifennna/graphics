package pixel

import (
	"github.com/go-gl/gl/all-core/gl"
)

type Point struct {
	X float64
	Y float64
}

func DrawPixel(point Point) {
	x := point.X / 100
	y := point.Y / 100

	gl.PointSize(4)
	gl.Begin(gl.POINTS)
	gl.Color3f(1.0, .5, 1.0)
	gl.Vertex2d(x, y)
	gl.End()
}
