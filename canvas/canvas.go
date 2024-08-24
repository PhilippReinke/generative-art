package canvas

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

// Canvas is a RGBA image on which 2D geometric elements can be drawn onto.
type Canvas struct {
	img           *image.RGBA
	width, height int
}

// New creates a canvas of given height and width.
func New(width, height int) (*Canvas, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("width and height must be positive")
	}
	return &Canvas{
		img:    image.NewRGBA(image.Rect(0, 0, width, height)),
		width:  width,
		height: height,
	}, nil
}

// Layer draws an image to the canvas.
func (c *Canvas) Layer(layer image.Image) {
	draw.Draw(
		c.img,
		c.img.Bounds(),
		layer,
		image.Point{},
		draw.Over,
	)
}

// Line draws a line to the canvas.
func (c *Canvas) Line(x0, y0, x1, y1 int, clr color.Color) {
	draw.DrawMask(
		c.img,
		c.img.Bounds(),
		&image.Uniform{clr},
		image.Point{}, // zero point
		LineMask(x0, y0, x1, y1),
		image.Point{},
		draw.Over,
	)
}

// Circle draws a circle to the canvas.
func (c *Canvas) Circle(x, y, r int, clr color.Color) {
	p := image.Point{x, y}
	draw.DrawMask(
		c.img,
		c.img.Bounds(),
		&image.Uniform{clr},
		image.Point{},
		CircleMask(p, r),
		image.Point{},
		draw.Over,
	)
}

// Rect draws a rectangle to the canvas.
func (c *Canvas) Rect(rect image.Rectangle, clr color.Color) {
	draw.Draw(
		c.img,
		rect,
		&image.Uniform{clr},
		image.Point{0, 0},
		draw.Over,
	)
}

// Image exports the canvas as image.
func (c *Canvas) Image() *image.RGBA {
	return c.img
}
