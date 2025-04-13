package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var loadingDialog dialog.Dialog

func startLoading(title, message string, parent fyne.Window) {

	if parent == nil {
		return
	}

	// 创建无限进度条
	content := widget.NewProgressBarInfinite()

	// 创建模态对话框
	loadingDialog = dialog.NewCustom(
		title,
		message,
		content,
		parent,
	)

	loadingDialog.Show()
}

func closeLoading() {
	if loadingDialog != nil {
		loadingDialog.Hide()
		loadingDialog = nil
	}
}
