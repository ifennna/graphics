package circles

import "go-graphics/pixel"

func drawBresenhams(start pixel.Point) {
	radius := start.Y - start.X
	decisionVariable := 3 - 2*radius

	current := start

	for current.X <= current.Y {
		drawMirroredPixels(current)

		if decisionVariable < 0 {
			decisionVariable += (4 * current.X) + 6
		} else {
			decisionVariable += 4*(current.X-current.Y) + 10
			current.Y--
		}

		current.X++
	}
}
