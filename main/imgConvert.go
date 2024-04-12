package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/esimov/stackblur-go"
)

func open(file string) image.Image {

	r, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	decodedImg, _, err := image.Decode(r)
	if err != nil {
		log.Fatal(err)
	}

	return decodedImg

}

func blur(decodedImg image.Image) {
	w := 200
	h := 100

	background := image.NewRGBA(image.Rect(0, 0, w, h))
	// Filling a Rectangle from https://go.dev/blog/image-draw
	draw.Draw(background, background.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)

}

func convertToPng(decodedImg image.Image) {

	w, err := os.Create("writetest.png")
	if err != nil {
		log.Fatal(err)
	}

	bluredImage, err := stackblur.Process(decodedImg, 12)
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(w, bluredImage)
	if err != nil {
		log.Fatal(err)
	}

	defer w.Close()

}
