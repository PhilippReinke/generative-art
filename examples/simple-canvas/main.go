package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/PhilippReinke/generative-art/canvas"
	"github.com/PhilippReinke/generative-art/util"
)

func main() {
	if err := example("out/simple-canvas.png"); err != nil {
		fmt.Println("Failed to run example:", err)
		os.Exit(1)
	}
}

func example(dstPath string) error {
	cnv, err := canvas.New(600, 600)
	if err != nil {
		return fmt.Errorf("failed to create canvas: %v", err)
	}

	cnv.Rect(image.Rect(0, 0, 400, 600), color.NRGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 100,
	})
	cnv.Rect(image.Rect(200, 0, 600, 600), color.NRGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 100,
	})

	cnv.Circle(300, 300, 40, color.NRGBA{
		255, 255, 255, 170,
	})

	cnv.Line(300, 0, 300, 600, color.Black)
	cnv.Line(0, 300, 600, 300, color.Black)
	cnv.Line(0, 0, 600, 600, color.Black)
	cnv.Line(600, 0, 0, 600, color.Black)

	if err := util.SavePNG(cnv.Image(), dstPath); err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	fmt.Println("Result stored under", dstPath)

	return nil
}
