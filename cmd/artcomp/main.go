package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"

	"github.com/iand/mixalist/pkg/imagegen"
)

func main() {
	if len(os.Args) < 2 {
		println("missing input image filename(s)")
		os.Exit(1)
	}

	imgs := []image.Image{}

	imgListIndex := 1
	for _, imgFilename := range os.Args[imgListIndex:] {
		fimg, err := os.Open(imgFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open image %s: %v", imgFilename, err)
			os.Exit(1)
		}

		img, _, err := image.Decode(fimg)
		fimg.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to read image data %s: %v", imgFilename, err)
			os.Exit(1)
		}

		imgs = append(imgs, img)
	}

	compositer := &imagegen.Compositer{
		DestWidth:     128,
		TilesPerSide:  3,
		LugDepthRatio: 0.15, // multiplied by the tile width to get lug depth
		LugWidthRatio: 0.15, // multiplied by the tile width to get lug width
	}

	outImg := compositer.Composite(imgs)
	_ = outImg
	foutName := "/tmp/composite.png"
	fout, err := os.OpenFile(foutName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to write output image %s: %v", foutName, err)
		os.Exit(1)
	}

	if err = png.Encode(fout, outImg); err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode output image %s: %v", foutName, err)
		os.Exit(1)
	}
	fout.Close()
}
