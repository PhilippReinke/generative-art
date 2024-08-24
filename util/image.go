package util

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"slices"
)

var (
	SupportedFileExt = []string{".jpg", ".jpeg", ".png"}
)

func LoadImage(path string) (image.Image, error) {
	fileExt := filepath.Ext(path)
	if !slices.Contains(SupportedFileExt, fileExt) {
		return nil, fmt.Errorf("unsported file extension: %v", fileExt)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode file: %v", err)
	}

	return img, nil
}

func SavePNG(img image.Image, path string) error {
	if filepath.Ext(path) != ".png" {
		return fmt.Errorf("file extension must be .png")
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return fmt.Errorf("failed to encode png: %v", err)
	}

	return nil
}

func SaveJPEG(img image.Image, path string, bgColor color.Color) error {
	fmt.Println("SaveJPEG needs to be implemented")
	return nil
}
