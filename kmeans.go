package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"sync"
)

type cluster struct {
	centroid   color.Color
	lCentroid  color.Color
	plCentroid color.Color
	members    []image.Point
}

func getClusters(k int, set image.Image) []*cluster {

	// init clusters
	clr := []*cluster{}
	for i := 0; i < k; i++ {
		c := new(cluster)
		c.centroid = getRandColor()
		c.lCentroid = color.RGBA{0, 0, 0, 0}
		c.plCentroid = c.lCentroid
		clr = append(clr, c)
	}

	// run k means algorithm
	for {
		partition(clr, set)

		if converged(clr) {
			return clr
		}
	}
}

func partition(clr []*cluster, set image.Image) {
	var w sync.WaitGroup
	w.Add(set.Bounds().Max.Y * set.Bounds().Max.X)
	var mtx sync.Mutex

	// assign all data points to the nearest cluster
	for y := set.Bounds().Min.Y; y < set.Bounds().Max.Y; y++ {
		go func(y int) {
			for x := set.Bounds().Min.X; x < set.Bounds().Max.X; x++ {
				i := indexNewCentroid(clr, set.At(x, y))
				mtx.Lock()
				clr[i].members = append(clr[i].members, image.Pt(x, y))
				mtx.Unlock()
				w.Done()
			}
		}(y)
	}
	w.Wait()

	var getAverage = func(pts []image.Point) color.Color {
		lPts := len(pts)
		if lPts < 1 {
			lPts = 1
		}

		var rSum, gSum, bSum uint32
		for _, p := range pts {
			r, g, b, _ := set.At(p.X, p.Y).RGBA()
			rSum += r
			gSum += g
			bSum += b
		}

		i := uint8(rSum / uint32(lPts))
		j := uint8(gSum / uint32(lPts))
		k := uint8(bSum / uint32(lPts))
		return color.RGBA{R: i, G: j, B: k, A: 100}
	}

	// update each centroid per cluster
	for _, c := range clr {
		c.plCentroid = c.lCentroid
		c.lCentroid = c.centroid
		c.centroid = getAverage(c.members)
	}
}

const stop = 1000000000

func converged(set []*cluster) bool {
	for _, c := range set {
		if euclidDis(c.centroid, c.lCentroid) > stop {
			return false
		}
	}
	return true
}

func getRandColor() color.Color {
	r := rand.Uint32()
	return color.RGBA{R: uint8(r), G: uint8(r >> 8), B: uint8(r >> 16), A: 0}
}

func indexNewCentroid(clr []*cluster, p color.Color) (idx int) {
	min := uint32(math.MaxUint32)

	for i, c := range clr {
		// find cluster with lowest color-distance
		d := euclidDis(c.centroid, p)
		if d < min {
			min = d
			idx = i
		}
	}

	return
}

// see image/color/color.go
func euclidDis(a, b color.Color) uint32 {
	ar, ag, ab, _ := a.RGBA()
	br, bg, bb, _ := b.RGBA()

	var sqDiff = func(x, y uint32) uint32 {
		return ((x - y) * (x - y)) >> 2
	}

	return sqDiff(ar, br) + sqDiff(ag, bg) + sqDiff(ab, bb)
}
