package imagegen

import (
	"github.com/iand/salience"
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"
)

type Compositer struct {
	DestWidth int
	TilesPerSide int
	LugDepthRatio float32 // multiplied by the tile width to get lug depth
	LugWidthRatio float32 // multiplied by the tile width to get lug width
}

func (c *Compositer) Composite(imgs []image.Image) (dest image.Image) {
	tileWidth := c.DestWidth / c.TilesPerSide
	//numTiles := c.TilesPerSide * c.TilesPerSide
	//lugDepth := int(c.LugDepthRatio * float32(tileWidth))
	//lugWidth := int(c.LugWidthRatio * float32(tileWidth))
	//scaledWidth := tileWidth + 2*lugDepth
	
	perm := rand.Perm(len(imgs))
	i := 0
	
	/*
	lugsSouth := make([]bool, numTiles)
	lugsEast := make([]bool, numTiles)
	
	for i := 0; i < numTiles; i++ {
		lugsSouth[i] = rand.Float32() < 0.5
		lugsEast[i] = rand.Float32() < 0.5
	}
	*/
	
	dest = image.NewRGBA(image.Rect(0, 0, c.DestWidth, c.DestWidth))
	
	for tileY := 0; tileY < c.TilesPerSide; tileY++ {
		tileTop := tileY * tileWidth
		
		for tileX := 0; tileX < c.TilesPerSide; tileX++ {
			tileLeft := tileX * tileWidth
			
			img := imgs[perm[i]]
			i++
			if i == len(perm) {
				perm = rand.Perm(len(imgs))
				i = 0
			}
			
			/*
			imageBounds := img.Bounds()
			imageWidth := imageBounds.Max.X - imageBounds.Min.X
			imageHeight := imageBounds.Max.Y - imageBounds.Min.Y
			meanDimension := (imageWidth + imageHeight)/2
			croppedWidth := meanDimension/2
			
			croppedImg := salience.Crop(img, croppedWidth, croppedWidth)
			scaledImg := image.NewRGBA(image.Rect(0, 0, scaledWidth, scaledWidth))
			scale(scaledImg, croppedImg)
			*/
			scaledImg := salience.Crop(img, tileWidth, tileWidth)
			
			r := image.Rect(tileLeft, tileTop, tileLeft+tileWidth, tileTop+tileWidth)
			//fmt.Printf("drawing src[%d, %d, %d, %d] onto dest[%d, %d, %d, %d]\n", 0, 0, tileWidth, tileWidth, tileLeft, tileTop, tileLeft+tileWidth, tileTop+tileWidth)
			draw.Draw(dest.(draw.Image), r, scaledImg, image.Pt(0, 0), draw.Src)
		}
	}
	
	return dest
}

func scale(dest draw.Image, src image.Image) {
	srcBounds := src.Bounds()
	srcWidth := srcBounds.Max.X - srcBounds.Min.X
	srcHeight := srcBounds.Max.Y - srcBounds.Min.Y
	
	destBounds := src.Bounds()
	destWidth := destBounds.Max.X - destBounds.Min.X
	destHeight := destBounds.Max.Y - destBounds.Min.Y
	
	srcMaxX := float64(srcWidth - 1)
	srcMaxY := float64(srcHeight - 1)
	destMaxX := float64(destWidth - 1)
	destMaxY := float64(destHeight - 1)
	
	for destY := 0; destY < destHeight; destY++ {
		srcY := (float64(destY) / destMaxY) * srcMaxY
		srcY1 := int(math.Floor(srcY))
		srcY2 := int(math.Ceil(srcY))
		
		for destX := 0; destX < destWidth; destX++ {
			srcX := (float64(destY) / destMaxX) * srcMaxX
			srcX1 := int(math.Floor(srcX))
			srcX2 := int(math.Ceil(srcX))
			
			r1, g1, b1, a1 := src.At(srcX1, srcY1).RGBA()
			r2, g2, b2, a2 := src.At(srcX2, srcY1).RGBA()
			r3, g3, b3, a3 := src.At(srcX1, srcY2).RGBA()
			r4, g4, b4, a4 := src.At(srcX2, srcY2).RGBA()
			
			dest.Set(destX, destY, color.RGBA64{
				R: uint16((r1+r2+r3+r4)/4),
				G: uint16((g1+g2+g3+g4)/4),
				B: uint16((b1+b2+b3+b4)/4),
				A: uint16((a1+a2+a3+a4)/4),
			})
		}
	}
}
