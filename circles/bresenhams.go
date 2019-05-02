package circles

import "graphics/pixel"

func drawBresenhams(midpoint pixel.Point, radius float64) {
	decisionVariable := 3 - 2*radius

	current := pixel.Point{X: 0, Y: radius}

	for current.X <= current.Y {
		drawMirroredPixels(midpoint, current)

		if decisionVariable < 0 {
			decisionVariable += (4 * current.X) + 6
		} else {
			decisionVariable += 4*(current.X-current.Y) + 10
			current.Y--
		}

		current.X++
	}
}
