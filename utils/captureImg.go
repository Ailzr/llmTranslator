package utils

import (
	"errors"
	"fmt"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
	"os"
)

func CaptureAllScreen() (*image.RGBA, error) {
	// 获取当前活跃的显示器数量
	n := screenshot.NumActiveDisplays()
	// 如果没有找到显示器，则记录日志
	if n <= 0 {
		return nil, errors.New("没有找到显示器")
	}
	// 获取第一个显示器的边界
	bounds := screenshot.GetDisplayBounds(0)
	// 截取整个屏幕的图像
	img, err := screenshot.CaptureRect(bounds)
	// 如果截取图像时发生错误，则记录日志
	if err != nil {
		return nil, err
	}
	return img, nil
}

// CaptureImg函数用于截取屏幕上的指定区域
func CaptureImg(min, max image.Point) (*image.RGBA, error) {
	// 获取当前活跃的显示器数量
	n := screenshot.NumActiveDisplays()
	// 如果没有找到显示器，则记录日志
	if n <= 0 {
		return nil, errors.New("没有找到显示器")
	}
	// 获取第一个显示器的边界
	bounds := screenshot.GetDisplayBounds(0)

	// 如果截取区域的宽度超过了显示器的宽度，则将宽度调整为显示器宽度减去x坐标
	if min.X < 0 || max.X > bounds.Dx() {
		return nil, errors.New("截屏宽度越界")
	}
	// 如果截取区域的高度超过了显示器的高度，则将高度调整为显示器高度减去y坐标
	if min.Y < 0 || max.Y > bounds.Dy() {
		return nil, errors.New("截屏高度越界")
	}

	// 截取指定区域的图像
	img, err := screenshot.CaptureRect(image.Rectangle{Min: min, Max: max})
	// 如果截取图像时发生错误，则记录日志
	if err != nil {
		return nil, err
	}
	return img, nil
}

func SaveImgToPng(img *image.RGBA, imgName string) {
	// 定义保存图像的文件名
	fileName := fmt.Sprintf("tmp_img/%s.png", imgName)
	// 创建文件
	file, _ := os.Create(fileName)
	// 关闭文件
	defer file.Close()
	// 将图像保存为PNG格式
	png.Encode(file, img)
}
