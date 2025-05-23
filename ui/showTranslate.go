package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image"
	"llmTranslator/configs"
	"llmTranslator/logHelper"
	"llmTranslator/pkg/llm"
	"llmTranslator/pkg/ocr"
	"llmTranslator/utils"
	"time"
)

var contentShow *widget.Label

func (mw *MainWindow) CreateShowWindow() {
	if mw.TranslatorWindow != nil {
		mw.TranslatorWindow.Close()
		mw.TranslatorWindow = nil
	}
	mw.TranslatorWindow = mw.App.NewWindow("Translator")

	mw.TranslatorWindow.SetIcon(resourceIconPng)

	contentShow = widget.NewLabel("")
	contentShow.Wrapping = fyne.TextWrapWord

	mw.TranslatorWindow.SetContent(container.NewScroll(contentShow))

	//TODO 翻译框大小从配置中读取，动态修改写入到配置中保存

	mw.TranslatorWindow.Resize(fyne.NewSize(float32(configs.Setting.Capture.EndX-configs.Setting.Capture.StartX), float32(configs.Setting.Capture.EndY-configs.Setting.Capture.StartY)))
}

// ShowTranslate函数用于显示翻译后的文本
func (mw *MainWindow) ShowTranslate(text string) {
	// 使用go关键字开启一个goroutine，用于异步执行
	go func() {
		// 判断contentShow是否为空
		if contentShow != nil {
			// 使用fyne.Do函数在主线程中执行
			fyne.Do(func() {
				// 设置contentShow的文本为传入的text参数
				contentShow.SetText(text)
				mw.TranslatorWindow.Show()
			})
		} else {
			// 如果contentShow为空，则显示一个信息对话框
			dialog.ShowInformation("错误", "未选择需要翻译的部分", mw.Window)
		}
	}()
}

// Translate函数用于翻译文本
func (mw *MainWindow) Translate() {
	go func() {
		// 获取OCR识别结果
		ocrResult := ocr.GetOCRResult()
		// 如果未识别到文字，显示未识别到文字
		if ocrResult == "" {
			fyne.Do(func() {
				mw.ShowTranslate("未识别到文字")
			})
		}

		// 在fyne主线程中执行
		fyne.Do(func() {
			// 如果不在托盘模式下，显示主窗口
			if !mw.isTray {
				mw.Window.Show()
			}
		})

		// 调用llm翻译接口，将ocr识别结果翻译为简体中文
		result := llm.Translate(ocrResult)
		if configs.Setting.AppSetting.ShowRawText {
			result = ocrResult + "\n--------------分割线----------------\n" + result
		}

		// 在fyne主线程中执行
		fyne.Do(func() {
			// 显示翻译结果
			mw.ShowTranslate(result)
		})
	}()
}

func hideWindowCapture() {
	// 在fyne主线程中执行
	fyne.Do(func() {
		// 隐藏翻译窗口
		mw.TranslatorWindow.Hide()
		// 显示翻译中
		//mw.ShowTranslate("翻译中...")
		// 如果不在托盘模式下，隐藏主窗口
		if !mw.isTray {
			mw.Window.Hide()
		}
	})

	time.Sleep(300 * time.Millisecond)

	// 截取屏幕图像
	img, err := utils.CaptureImg(image.Point{X: configs.Setting.Capture.StartX, Y: configs.Setting.Capture.StartY}, image.Point{X: configs.Setting.Capture.EndX, Y: configs.Setting.Capture.EndY})
	// 将图像保存为png格式
	utils.SaveImgToPng(img, "tmp")
	// 如果截取失败，记录错误并显示截图失败
	if err != nil {
		logHelper.Error(err.Error())
		fyne.Do(func() {
			mw.ShowTranslate("截图失败")
		})
		return
	}
}
