package main

import (
	"generativeart/sketch"
	"github.com/fogleman/gg"
	"log"
)

var (
	sourceImgName   = "source.jpg"
	outputImgName   = "out.png"
	totalCycleCount = 5000
)

func main() {
	image, err := gg.LoadImage(sourceImgName)
	if err != nil {
		log.Panicln(err)
	}

	destWidth := 2000
	params := sketch.UserParams{
		DestWidth:                destWidth,
		DestHeight:               2000,
		StrokeRatio:              0.75,
		StrokeReduction:          0.002,
		StrokeInversionThreshold: 0.05,
		StrokeJitter:             int(0.1 * float64(destWidth)),
		InitialAlpha:             0.1,
		AlphaIncrease:            0.06,
		MinEdgeCount:             3,
		MaxEdgeCount:             4,
	}

	s := sketch.NewSketch(image, params)
	for i := 0; i < totalCycleCount; i++ {
		s.Update()
	}
	err = gg.SavePNG(outputImgName, s.Output())
	if err != nil {
		log.Panicln(err)
	}
}
