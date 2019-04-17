package circles

import "go-graphics/pixel"

func Draw(start pixel.Point, algorithm string) {
	switch algorithm {
	case "bresenhams":
		drawBresenhams(start)
	}
}

func drawMirroredPixels(point pixel.Point) {
	sectorOne := point
	sectorTwo := pixel.Point{X: point.Y, Y: point.X}
	sectorThree := pixel.Point{X: -point.X, Y: point.Y}
	sectorFour := pixel.Point{X: -point.Y, Y: point.X}
	sectorFive := pixel.Point{X: -point.X, Y: -point.Y}
	sectorSix := pixel.Point{X: -point.Y, Y: -point.X}
	sectorSeven := pixel.Point{X: point.X, Y: -point.Y}
	sectorEight := pixel.Point{X: point.Y, Y: -point.X}

	pixel.DrawPixel(sectorOne)
	pixel.DrawPixel(sectorTwo)
	pixel.DrawPixel(sectorThree)
	pixel.DrawPixel(sectorFour)
	pixel.DrawPixel(sectorFive)
	pixel.DrawPixel(sectorSix)
	pixel.DrawPixel(sectorSeven)
	pixel.DrawPixel(sectorEight)
}
