package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand/v2"
	"os"

	"github.com/PhilippReinke/generative-art/art"
	"github.com/PhilippReinke/generative-art/canvas"
	"github.com/PhilippReinke/generative-art/util"
)

const (
	maxRGBA = 1<<16 - 1
)

func main() {
	if err := example(
		"assets/monet.jpg",
		"out/colored-rects.png",
	); err != nil {
		fmt.Println("Failed to run example:", err)
		os.Exit(1)
	}
}

func example(srcPath, dstPath string) error {
	src, err := util.LoadImage(srcPath)
	if err != nil {
		return fmt.Errorf("failed to load image: %v", err)
	}

	cfg := genConfig{
		src:             src,
		rectSize:        10,
		numOfIterations: 100_000,
	}
	cnv, err := art.Generate(cfg, generator)
	if err != nil {
		return fmt.Errorf("failed to generate art: %v", err)
	}

	if err := util.SavePNG(cnv.Image(), dstPath); err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	fmt.Println("Result stored under", dstPath)

	return nil
}

type genConfig struct {
	src             image.Image
	rectSize        int
	numOfIterations int
}

func (c genConfig) Width() int {
	return c.src.Bounds().Dx()
}

func (c genConfig) Height() int {
	return c.src.Bounds().Dy()
}

func (c genConfig) MaxIterations() int {
	return c.numOfIterations
}

func generator(cnv *canvas.Canvas, cfgInterface art.GenConfig) bool {
	cfg := cfgInterface.(genConfig) // assert type

	x := rand.IntN(cfg.Width())
	y := rand.IntN(cfg.Height())
	r, g, b, _ := cfg.src.At(x, y).RGBA()

	// draw rectangle of that color on dst
	cnv.Rect(
		image.Rect(
			x-cfg.rectSize, y-cfg.rectSize,
			x+cfg.rectSize, y+cfg.rectSize,
		),
		color.NRGBA{
			R: uint8(255 * float64(r) / maxRGBA),
			G: uint8(255 * float64(g) / maxRGBA),
			B: uint8(255 * float64(b) / maxRGBA),
			A: 100,
		},
	)
	return true
}
