package utils

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
)

// 在原图上绘制文字，直接修改传入的 image.RGBA
func DrawTextOnImage(img image.Image, text string) *image.RGBA {
	// 转为可写的 RGBA 图像
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	col := color.RGBA{255, 0, 0, 255} // 红色字体

	point := fixed.Point26_6{
		X: fixed.I(10),
		Y: fixed.I(30),
	}

	d := &font.Drawer{
		Dst:  rgba,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(text)

	return rgba
}
