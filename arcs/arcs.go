package arcs

import "go-graphics/pixel"

func Draw(start pixel.Point, end pixel.Point) {
	radius := start.Y - start.X
	decisionVariable := 3 - 2*radius

	current := start

	for current.X <= current.Y {
		pixel.DrawPixel(current)

		if decisionVariable < 0 {
			decisionVariable += (4 * current.X) + 6
		} else {
			decisionVariable += 4*(current.X-current.Y) + 10
			current.Y--
		}

		current.X++
	}
}
