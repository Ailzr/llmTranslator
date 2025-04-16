package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"llmTranslator/configs"
	"strings"
)

func createAppSettingForm() *widget.Form {
	form := &widget.Form{}

	remindLabel := widget.NewLabel("支持的按键：Ctrl、Shift、Alt、F1-F12、A-Z | 输入格式：Ctrl+Shift+T")
	remindLabel.Wrapping = fyne.TextWrapWord

	// 从配置读取初始值
	trCombo := configs.Setting.HotKey.Translate
	cpCombo := configs.Setting.HotKey.Capture
	tcCombo := configs.Setting.HotKey.CaptureTranslate
	ctcCombo := configs.Setting.HotKey.CaptureToClipboard
	defaultTraySet := configs.Setting.DefaultTray

	// 构造表单控件
	translateEntry := widget.NewEntry()
	translateEntry.SetText(trCombo)
	translateEntry.SetPlaceHolder("输入快捷键组合")

	captureEntry := widget.NewEntry()
	captureEntry.SetText(cpCombo)
	captureEntry.SetPlaceHolder("输入快捷键组合")

	tcEntry := widget.NewEntry()
	tcEntry.SetText(tcCombo)
	tcEntry.SetPlaceHolder("输入快捷键组合")

	ctcEntry := widget.NewEntry()
	ctcEntry.SetText(ctcCombo)
	ctcEntry.SetPlaceHolder("输入快捷键组合")

	showRawText := widget.NewCheck("显示原文", nil)
	showRawText.Checked = configs.Setting.ShowRawText

	defaultTray := widget.NewCheck("默认托盘", nil)
	defaultTray.Checked = defaultTraySet

	form.AppendItem(widget.NewFormItem("快捷键设置", remindLabel))
	form.AppendItem(widget.NewFormItem("框选区翻译热键", translateEntry))
	form.AppendItem(widget.NewFormItem("选区热键", captureEntry))
	form.AppendItem(widget.NewFormItem("截图翻译热键", tcEntry))
	form.AppendItem(widget.NewFormItem("截图热键", ctcEntry))
	form.AppendItem(widget.NewFormItem("显示原文", showRawText))
	form.AppendItem(widget.NewFormItem("启动时默认托盘", defaultTray))

	form.SubmitText = "保存"
	form.OnSubmit = func() {
		tText := strings.TrimSpace(translateEntry.Text)
		cText := strings.TrimSpace(captureEntry.Text)
		tcText := strings.TrimSpace(tcEntry.Text)
		ctcText := strings.TrimSpace(ctcEntry.Text)

		// 基本非空检查
		if tText == "" || cText == "" || tcText == "" {
			dialog.ShowError(fmt.Errorf("请填写完整的组合键"), mw.Window)
			return
		}
		// 冲突检测：字符串一致（忽略大小写、空格）
		if strings.EqualFold(strings.ReplaceAll(tText, " ", ""), strings.ReplaceAll(cText, " ", "")) {
			dialog.ShowError(fmt.Errorf("框选区翻译热键和截图热键不能相同"), mw.Window)
			return
		}
		// 校验并拆分
		_, _, err := ParseHotKey(tText)
		if err != nil {
			dialog.ShowError(fmt.Errorf("框选区翻译热键格式错误：%v", err), mw.Window)
			return
		}
		_, _, err = ParseHotKey(cText)
		if err != nil {
			dialog.ShowError(fmt.Errorf("截图热键格式错误：%v", err), mw.Window)
			return
		}
		_, _, err = ParseHotKey(tcText)
		if err != nil {
			dialog.ShowError(fmt.Errorf("截图翻译热键格式错误：%v", err), mw.Window)
			return
		}
		_, _, err = ParseHotKey(ctcText)
		if err != nil {
			dialog.ShowError(fmt.Errorf("截图翻译热键格式错误：%v", err), mw.Window)
			return
		}

		// 保存到配置
		configs.Setting.HotKey.Translate = tText
		configs.Setting.HotKey.Capture = cText
		configs.Setting.HotKey.CaptureTranslate = tcText
		configs.Setting.HotKey.CaptureToClipboard = ctcText
		configs.Setting.ShowRawText = showRawText.Checked
		configs.Setting.DefaultTray = defaultTray.Checked
		configs.WriteSettingToFile()

		// 立即重新注册热键
		UnregisterAllHotKey()
		RegisterAllHotKey()

		dialog.ShowInformation("保存成功", "设置已保存", mw.Window)
	}

	form.CancelText = "取消"
	form.OnCancel = func() {
		translateEntry.SetText(trCombo)
		captureEntry.SetText(cpCombo)
		tcEntry.SetText(tcCombo)
	}

	return form
}
