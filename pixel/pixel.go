package pixel

import "github.com/go-gl/gl/all-core/gl"

func DrawPixel(x float32, y float32) {
	gl.PointSize(10)
	gl.Begin(gl.POINTS)
	gl.Color3f(1.0, 1.0, 1.0)
	gl.Vertex2f(x, y)
	gl.End()
}
