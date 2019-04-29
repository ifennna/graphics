package arcs

import (
	"go-graphics/pixel"
)

func Draw(midpoint pixel.Point, radius float64, upper bool) {
	decisionVariable := 1 - radius

	current := pixel.Point{X: 0, Y: radius}

	for current.X <= current.Y {
		if upper == true {
			drawUpperMirroredPixels(midpoint, current)
		} else {
			drawLowerMirroredPixels(midpoint, current)
		}

		if decisionVariable < 0 {
			decisionVariable += (2 * current.X) + 3
		} else {
			decisionVariable += 2*(current.X-current.Y) + 5
			current.Y--
		}

		current.X++
	}
}

func drawUpperMirroredPixels(midpoint pixel.Point, point pixel.Point) {
	sectorFive := pixel.Point{X: -point.X + midpoint.X, Y: -point.Y + midpoint.Y}
	//sectorSix := pixel.Point{X: -point.Y + midpoint.X, Y: -point.X + midpoint.Y}
	sectorSeven := pixel.Point{X: point.X + midpoint.X, Y: -point.Y + midpoint.Y}
	//sectorEight := pixel.Point{X: point.Y + midpoint.X, Y: -point.X + midpoint.Y}

	pixel.DrawPixel(sectorFive)
	//pixel.DrawPixel(sectorSix)
	pixel.DrawPixel(sectorSeven)
	//pixel.DrawPixel(sectorEight)
}

func drawLowerMirroredPixels(midpoint pixel.Point, point pixel.Point) {
	sectorOne := pixel.Point{X: point.X + midpoint.X, Y: point.Y + midpoint.Y}
	//sectorTwo := pixel.Point{X: point.Y + midpoint.X, Y: point.X + midpoint.Y}
	sectorThree := pixel.Point{X: -point.X + midpoint.X, Y: point.Y + midpoint.Y}
	//sectorFour := pixel.Point{X: -point.Y + midpoint.X, Y: point.X + midpoint.Y}

	pixel.DrawPixel(sectorOne)
	//pixel.DrawPixel(sectorTwo)
	pixel.DrawPixel(sectorThree)
	//pixel.DrawPixel(sectorFour)
}
