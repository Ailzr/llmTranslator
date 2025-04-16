package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"image"
	"llmTranslator/configs"
	"llmTranslator/logHelper"
	"llmTranslator/pkg/llm"
	"llmTranslator/pkg/ocr"
	"llmTranslator/utils"
	"time"
)

var bg *canvas.Image

func (mw *MainWindow) CaptureSelectArea(onSelect func(image.Rectangle)) {
	fyne.Do(func() {
		if mw.CaptureWindow != nil {
			mw.CaptureWindow.Close()
			mw.CaptureWindow = nil
		}

		mw.CaptureWindow = mw.App.NewWindow("截屏")
		mw.CaptureWindow.SetFullScreen(true)
		mw.CaptureWindow.SetPadded(false)

		if !mw.isTray {
			mw.Window.Hide()
		}
		mw.TranslatorWindow.Hide()
	})

	go func() {
		time.Sleep(300 * time.Millisecond)
		img, err := utils.CaptureAllScreen()
		if err != nil {
			logHelper.Error(err.Error())
			logHelper.WriteLog(err.Error())
			return
		}
		go utils.SaveImgToPng(img, "screen")

		bg = canvas.NewImageFromImage(img)
		bg.FillMode = canvas.ImageFillStretch

		result := make(chan image.Rectangle)
		overlay := newSelectOverlay(result)

		fyne.Do(func() {
			mw.CaptureWindow.SetContent(container.NewStack(bg, overlay))
			mw.CaptureWindow.Canvas().(desktop.Canvas).SetOnKeyDown(func(e *fyne.KeyEvent) {
				if e.Name == fyne.KeyEscape {
					mw.CaptureWindow.Close()
				}
			})
			mw.CaptureWindow.Show()
		})

		go func() {
			sel := <-result
			onSelect(sel)
		}()
	}()
}

func (mw *MainWindow) CaptureToClipboard() {
	mw.CaptureSelectArea(func(sel image.Rectangle) {
		fyne.Do(func() {
			mw.CaptureWindow.Close()
		})
		img, err := utils.LoadPngFromTmp("screen")
		if err != nil {
			logHelper.Error(err.Error())
			logHelper.WriteLog(err.Error())
			return
		}
		subImg := img.SubImage(image.Rect(sel.Min.X, sel.Min.Y, sel.Max.X, sel.Max.Y))

		if err != nil {
			logHelper.Error(err.Error())
			logHelper.WriteLog(err.Error())
			return
		}
		go func() {
			if err := utils.ImageToClipboard(subImg); err != nil {
				logHelper.Error(err.Error())
				logHelper.WriteLog(err.Error())
				return
			}
		}()
	})
}

func (mw *MainWindow) CaptureAndTranslate() {
	//TODO 有bug，需要修复，暂不可用
	mw.CaptureSelectArea(func(sel image.Rectangle) {
		go func() {
			img, err := utils.CaptureImg(sel.Min, sel.Max)
			if err != nil {
				logHelper.Error(err.Error())
				logHelper.WriteLog(err.Error())
				mw.CaptureWindow.Close()
				return
			}
			utils.SaveImgToPng(img, "tmp")
			text := ocr.GetOCRResult()

			var loadingWin fyne.Window

			// 创建 loading 窗口（在主线程中）
			fyne.Do(func() {
				loadingWin = ShowLoadingWindow()
			})

			// 模拟流式翻译（逐段返回字符串）
			translated := llm.Translate(text, "简体中文")

			// 翻译完成，绘图 + 弹出图像窗口
			fyne.Do(func() {
				loadingWin.Close()
				drawImg := utils.DrawTextOnImage(img, translated)
				ShowImageInNewWindow(drawImg)
			})
		}()
	})
}

func (mw *MainWindow) CaptureAndSaveSelection() {
	mw.CaptureSelectArea(func(sel image.Rectangle) {
		dialog.ShowConfirm("确认选区",
			fmt.Sprintf("是否使用坐标 %v ？", sel),
			func(ok bool) {
				if ok {
					fyne.Do(func() {
						mw.CaptureWindow.Close()
					})
					//将选取到的坐标保存起来
					configs.Setting.Capture.StartX = sel.Min.X
					configs.Setting.Capture.StartY = sel.Min.Y
					configs.Setting.Capture.EndX = sel.Max.X
					configs.Setting.Capture.EndY = sel.Max.Y
					configs.WriteSettingToFile()
					mw.CreateShowWindow()
				} else {
					// 用户取消，只隐藏选区框，允许重新框选
					fyne.Do(func() {
						mw.CaptureWindow.Close()
					})
				}
			},
			mw.CaptureWindow,
		)
	})
}

func ShowImageInNewWindow(img image.Image) {
	win := mw.App.NewWindow("翻译完成")
	raster := canvas.NewImageFromImage(img)
	raster.FillMode = canvas.ImageFillContain

	win.SetContent(container.NewStack(raster))
	win.Resize(fyne.NewSize(800, 600))
	win.Show()
}
