package circles

import "go-graphics/pixel"

func drawMidpoint(midpoint pixel.Point, radius float64) {
	decisionVariable := 1 - radius

	current := pixel.Point{X: 0, Y: radius}

	for current.X <= current.Y {
		drawMirroredPixels(midpoint, current)

		if decisionVariable < 0 {
			decisionVariable += (2 * current.X) + 3
		} else {
			decisionVariable += 2*(current.X-current.Y) + 5
			current.Y--
		}

		current.X++
	}
}
