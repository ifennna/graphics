package lines

import (
	"go-graphics/pixel"
)

func drawBresenhams(line Line) {
	if line.slope > 0 {
		if line.slope < 1 {
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
		} else {
			differenceInY, differenceInX := getXAndYDifferences(line.start, line.end)
			decisionVariable := 2*differenceInY - differenceInX

			currentPixel := pixel.Point{X: line.start.Y, Y: line.start.X}
			endPixel := pixel.Point{X: line.end.Y, Y: line.end.X}

			pixel.DrawPixel(currentPixel)

			for currentPixel.X != endPixel.X {
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
	} else {
		if line.slope > -1 {
			differenceInY, differenceInX := getXAndYDifferences(line.start, line.end)
			decisionVariable := 2*differenceInY - differenceInX

			currentPixel := pixel.Point{X: -line.start.X, Y: line.start.Y}
			endPixel := pixel.Point{X: -line.end.X, Y: line.end.Y}

			pixel.DrawPixel(currentPixel)

			for currentPixel.X != endPixel.X {
				currentPixel.X++

				if decisionVariable < 0 {
					decisionVariable += 2 * differenceInY
				} else {
					currentPixel.Y++
					decisionVariable += 2 * (differenceInY - differenceInX)
				}

				pixel.DrawPixel(currentPixel)
			}
		} else {
			differenceInY, differenceInX := getXAndYDifferences(line.start, line.end)
			decisionVariable := 2*differenceInY - differenceInX

			currentPixel := pixel.Point{X: -line.start.Y, Y: line.start.X}
			endPixel := pixel.Point{X: -line.end.Y, Y: line.end.X}

			pixel.DrawPixel(currentPixel)

			for currentPixel.X != endPixel.X {
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
	}
}
