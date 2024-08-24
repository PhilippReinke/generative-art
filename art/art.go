package art

import (
	"fmt"

	"github.com/PhilippReinke/generative-art/canvas"
)

// GenConfig represents a generic generator configuration.
type GenConfig interface {
	// Width of target canvas.
	Width() int
	// Height of target canvas.
	Height() int
	// MaxIterations of generative process.
	MaxIterations() int
}

// Generator is meant to perform an iteration in the generative process.
//
// Return false if you wish to end the process.
type Generator func(cnv *canvas.Canvas, cfg GenConfig) bool

// Generate generates the generative art ;)
func Generate(cfg GenConfig, g Generator) (*canvas.Canvas, error) {
	cnv, err := canvas.New(cfg.Width(), cfg.Height())
	if err != nil {
		return nil, fmt.Errorf("failed to create canvas: %v", err)
	}

	for range cfg.MaxIterations() {
		if !g(cnv, cfg) {
			break
		}
	}

	return cnv, nil
}
