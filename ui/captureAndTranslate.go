package ui

import (
	"fyne.io/fyne/v2"
	"image"
	"llmTranslator/logHelper"
	"llmTranslator/utils"
)

func (mw *MainWindow) CaptureAndTranslate() {
	mw.CaptureSelectArea(func(sel image.Rectangle) {
		fyne.Do(func() {
			mw.CaptureWindow.Close()
		})

		go func() {
			screen, err := utils.LoadPngFromTmp("screen")
			if err != nil {
				logHelper.Error(err.Error())
				return
			}
			subImg := screen.SubImage(sel)
			utils.SaveImgToPng(subImg.(*image.RGBA), "tmp")
			mw.Translate()
		}()
	})
}
