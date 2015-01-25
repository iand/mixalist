package imagegen

/*
import (
    "image"
    "image/color"
    //"image/draw"
)

var (
    Transparent = color.Alpha{0x00}
    Opaque = color.Alpha{0xff}
)

func MakeLugMask(outerWidth, outerHeight, lugDepth, lugWidth int) image.Image {
    innerWidth := outerWidth - 2*lugDepth
    innerHeight := outerHeight - 2*lugDepth
    
    innerLeft := lugDepth
    innerTop := lugDepth
    innerRight := innerLeft + innerWidth
    innerBottom := innerTop + innerHeight
    
    imgRect := image.Rect(0, 0, outerWidth, outerHeight)
    img := image.NewAlpha(imgRect)
    
    // Fill background transparent
    draw.Draw(img, imgRect, image.NewUniform(Transparent), image.Pt(0, 0))
    
    // Fill inner section opaque
    innerRect := image.Rect(innerLeft, innerTop, innerRight, innerBottom)
    draw.Draw(img, innerRect, image.NewUniform(Opaque), image.Pt(0, 0))
    
    // // Top lug
    // topLugRect := image.Rect((outerWidth - lugWidth)/2, innerTop, (outerWidth + lugWidth)/2, innerTop + lugDepth)
    // topLug := makeVertLug(topLugRect, Transparent, Opaque)
    // draw.Draw(img, topLugRect, topLug, topLugRect.Min)
    
    return img
}

func makeVertLug(rect *image.Rectangle, fg, bg color.Color) image.Image {
    img := image.NewAlpha(rect)
    draw.Draw(img, rect, image.NewUniform(bg), image.Pt(0, 0))
    
    for s := -2.985; s <= 2.985; s += 0.005 {
        secs := math.Sech(s)
        // 0 < fx < 1, 0 < fy < 1
        fx := 0.9554*s*((1.33-math.Sech(2(s-2.2))-math.Sech(2(s+2.2)))/1.33)/2 + 0.5
        fy := ((((2*secs*(2-secs))*math.Sech(math.Pow(math.Abs(s)/2,1.4))))*1.06875-0.1375)/2
        x := int(fx * float32(rect.Dx()))
        y := int(fy * float32(rect.Dy()))
        img.Set(rect.Min.X + x, rect.Min.Y + y, fg)
    }
    
    return img
}
*/
