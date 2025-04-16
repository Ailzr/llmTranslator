package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func ShowLoadingWindow() fyne.Window {
	win := mw.App.NewWindow("处理中...")
	label := canvas.NewText("正在翻译中...", color.RGBA{0, 0, 255, 255})
	label.Alignment = fyne.TextAlignCenter

	content := container.NewVBox(
		label,
		widget.NewProgressBarInfinite(),
	)

	win.SetContent(content)
	win.Resize(fyne.NewSize(200, 100))
	win.CenterOnScreen()
	win.Show()

	return win
}
