package main

import (
	"image"
	"github.com/lucasb-eyer/go-colorful"
)

func clusterImage(k int, img image.Image) image.Image {
	cls := getClusters(k, img)

	// assign each cluster a random color
	for _, c := range cls {
		c.centroid = colorful.WarmColor() // save color in centroid
	}

	// create new image
	newImg := image.NewRGBA(img.Bounds())
	for _, c := range cls {
		for _, p := range c.members {
			newImg.Set(p.X, p.Y, c.centroid)
		}
	}
	return newImg
}
