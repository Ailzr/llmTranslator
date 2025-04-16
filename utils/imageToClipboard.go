package utils

import (
	"bytes"
	"golang.design/x/clipboard"
	"image"
	"image/png"
)

func ImageToClipboard(img image.Image) error {
	//初始化
	clipboard.Init()

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return err
	}
	data := buf.Bytes()

	// 二进制数据写入剪贴板
	// 注意：不同平台剪贴板格式可能需要额外指定 MIME 类型
	clipboard.Write(clipboard.FmtImage, data)
	return nil
}
