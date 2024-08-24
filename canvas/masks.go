package canvas

import (
	"image"
	"image/color"

	"github.com/PhilippReinke/generative-art/util"
)

/*
	Recommended read
	https://go.dev/blog/image
	https://go.dev/blog/image-draw

*/

// Circle represents a circle mask.
type Circle struct {
	p image.Point
	r int
}

func CircleMask(p image.Point, r int) *Circle {
	return &Circle{p, r}
}

func (c *Circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *Circle) Bounds() image.Rectangle {
	return image.Rect(
		c.p.X-c.r, c.p.Y-c.r,
		c.p.X+c.r, c.p.Y+c.r,
	)
}

func (c *Circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

// Line represents a line mask.
//
// TODO: implement thickness of line
func LineMask(x0, y0, x1, y1 int) *image.Alpha {
	rect := image.Rect(
		min(x0, x1), min(y0, y1),
		max(x0, x1)+1, max(y0, y1)+1,
	)
	mask := image.NewAlpha(rect)

	// Bresenham
	dx := util.Abs(x1 - x0)
	dy := util.Abs(y1 - y0)
	sx := -1
	if x0 < x1 {
		sx = 1
	}
	sy := -1
	if y0 < y1 {
		sy = 1
	}
	err := dx - dy

	for {
		mask.SetAlpha(x0, y0, color.Alpha{255})
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}

	return mask
}
