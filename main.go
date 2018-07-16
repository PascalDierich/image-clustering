package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
)

var inPath = flag.String("in", "", "Path to input image")
var outPath = flag.String("out", "${HOME}/img_clustered.png", "Path to save clustered image")
var k = flag.Int("k", 10, "Number of cluster")

func main() {
	flag.Parse()

	if *inPath == "" {
		fmt.Println("require image to process")
		os.Exit(0)
	}
	if *k < 1 {
		fmt.Println("require positive amount of clusters")
	}

	imgFile, err := os.Open(*inPath)
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()
	defer imgFile.Seek(0, 0)

	imgData, err := png.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}

	img := clusterImage(*k, imgData)

	cImg, err := os.Create(*outPath)
	if err != nil {
		log.Fatal(err)
	}
	defer cImg.Close()
	defer cImg.Seek(0, 0)

	png.Encode(cImg, img)
}
