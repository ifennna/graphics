package lines

import "go-graphics/pixel"

func drawBresenhams(line Line) {
	differenceInY, differenceInX := getXAndYDifferences(line.start, line.end)
	decisionVariable := 2*differenceInY - differenceInX

	currentPixel := line.start

	pixel.DrawPixel(currentPixel)

	for currentPixel.X != line.end.X {
		currentPixel.X++

		if decisionVariable < 0 {
			decisionVariable += 2 * differenceInY
		} else {
			currentPixel.Y++
			decisionVariable += 2 * (differenceInY - differenceInX)
		}

		pixel.DrawPixel(currentPixel)
	}
}

//func setPixelForLine(line Line, point pixel.Point) {
//	x := point.X
//	y := point.Y
//
//	if line.slope < 0 {
//		if line.slope > 1 {
//			pixel.DrawPixel(point)
//		} else {
//			newPoint := pixel.Point{X: x, Y: -y}
//			pixel.DrawPixel(newPoint)
//		}
//	} else {
//
//	} else if line.slope > 1 {
//		newPoint := pixel.Point{X: y, Y: x}
//		pixel.DrawPixel(newPoint)
//	} else if line.slope < 0 {
//		newPoint := pixel.Point{X: y, Y: x}
//		pixel.DrawPixel(newPoint)
//	} else if line.slope < -1 {
//		newPoint := pixel.Point{X: -x, Y: y}
//		pixel.DrawPixel(newPoint)
//	}
//}
