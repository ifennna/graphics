package circles

import "go-graphics/pixel"

func drawMidpoint(start pixel.Point) {
	radius := start.Y - start.X
	decisionVariable := 1 - radius

	current := start

	for current.X <= current.Y {
		drawMirroredPixels(current)

		if decisionVariable < 0 {
			decisionVariable += (2 * current.X) + 3
		} else {
			decisionVariable += 2*(current.X-current.Y) + 5
			current.Y--
		}

		current.X++
	}
}
