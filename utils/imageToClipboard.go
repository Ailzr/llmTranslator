package utils

import (
	"bytes"
	"github.com/skanehira/clipboard-image/v2"
	"image"
	"image/png"
)

func ImageToClipboard(img image.Image) error {
	// 创建内存缓冲区
	var buf bytes.Buffer

	// 编码为PNG格式到内存
	if err := png.Encode(&buf, img); err != nil {
		return err
	}

	// 创建io.Reader
	reader := bytes.NewReader(buf.Bytes())

	// 写入剪贴板 ✅ 关键步骤
	if err := clipboard.Write(reader); err != nil {
		return err
	}
	return nil
}
