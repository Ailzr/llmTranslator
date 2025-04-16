package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/color"
	"llmTranslator/configs"
	"llmTranslator/logHelper"
	"llmTranslator/utils"
	"time"
)

//TODO 拆分截图功能

func (mw *MainWindow) CaptureRectangle() {

	fyne.Do(func() {
		if mw.CaptureWindow != nil {
			mw.CaptureWindow.Close()
			mw.CaptureWindow = nil
		}
		//创建截屏窗口
		mw.CaptureWindow = mw.App.NewWindow("截屏")
		mw.CaptureWindow.SetFullScreen(true)
		mw.CaptureWindow.SetPadded(false)

		//截图前将主窗口和翻译窗口隐藏
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

		bg := canvas.NewImageFromImage(img)
		bg.FillMode = canvas.ImageFillStretch

		// 准备 overlay
		result := make(chan image.Rectangle)
		overlay := newSelectOverlay(result)

		fyne.Do(func() {
			mw.CaptureWindow.SetContent(container.NewStack(bg, overlay))
		})

		mw.CaptureWindow.Canvas().(desktop.Canvas).SetOnKeyDown(func(e *fyne.KeyEvent) {
			switch e.Name {
			case fyne.KeyEscape:
				mw.CaptureWindow.Close()
			}
		})

		fyne.Do(func() {
			mw.CaptureWindow.Show()
			//重新显示主窗口和翻译窗口
			if !mw.isTray {
				mw.Window.Show()
			}
			mw.TranslatorWindow.Show()
		})

		go func() {
			sel := <-result
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
		}()
	}()
}

type selectOverlay struct {
	widget.BaseWidget
	mask      *canvas.Raster
	rect      *canvas.Rectangle
	start     fyne.Position
	curr      fyne.Position
	selecting bool
	result    chan image.Rectangle
}

func newSelectOverlay(result chan image.Rectangle) *selectOverlay {
	o := &selectOverlay{
		rect:   canvas.NewRectangle(color.NRGBA{R: 0, G: 255, B: 0, A: 0}), // 先只画边框
		result: result,
	}
	o.rect.StrokeColor = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	o.rect.StrokeWidth = 2

	// 蒙版：除了选区内部，全屏都用半透明灰
	o.mask = canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color {
		sx, sy := int(o.start.X), int(o.start.Y)
		cx, cy := int(o.curr.X), int(o.curr.Y)
		minX, maxX := min(sx, cx), max(sx, cx)
		minY, maxY := min(sy, cy), max(sy, cy)
		if x >= minX && x <= maxX && y >= minY && y <= maxY {
			// 选区内完全透明，露出原始截图
			return color.NRGBA{0, 0, 0, 0}
		}
		// 选区外半透明灰
		return color.NRGBA{R: 0, G: 0, B: 0, A: 128}
	})

	o.ExtendBaseWidget(o)
	return o
}

type overlayRenderer struct {
	mask *canvas.Raster
	rect *canvas.Rectangle
}

func (o *selectOverlay) CreateRenderer() fyne.WidgetRenderer {
	//return &overlayRenderer{objects: []fyne.CanvasObject{o.rect}}
	return &overlayRenderer{mask: o.mask, rect: o.rect}
}

func (r *overlayRenderer) Layout(size fyne.Size) {
	// 蒙版铺满
	r.mask.Resize(size)
	r.mask.Move(fyne.NewPos(0, 0))
	// rect 的位置和大小由 overlay 自己在 MouseMoved 里设置
}
func (r *overlayRenderer) MinSize() fyne.Size { return fyne.NewSize(0, 0) }
func (r *overlayRenderer) Refresh() {
	r.mask.Refresh()
	r.rect.Refresh()
}
func (r *overlayRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.mask, r.rect}
}
func (r *overlayRenderer) Destroy() {}

func (o *selectOverlay) MouseDown(e *desktop.MouseEvent) {
	if e.Button != desktop.MouseButtonPrimary {
		return
	}
	o.start = e.Position
	o.curr = e.Position
	o.selecting = true

	o.rect.Move(o.start)
	o.rect.Resize(fyne.NewSize(0, 0))
	o.rect.Show()

	o.mask.Refresh()
	o.rect.Refresh()
}

func (o *selectOverlay) MouseMoved(e *desktop.MouseEvent) {
	if !o.selecting {
		return
	}
	o.curr = e.Position

	x1, y1 := int(o.start.X), int(o.start.Y)
	x2, y2 := int(o.curr.X), int(o.curr.Y)
	minX, minY := float32(min(x1, x2)), float32(min(y1, y2))
	w, h := float32(abs(x2-x1)), float32(abs(y2-y1))

	o.rect.Move(fyne.NewPos(minX, minY))
	o.rect.Resize(fyne.NewSize(w, h))

	o.mask.Refresh()
	o.rect.Refresh()
}

func (o *selectOverlay) MouseUp(e *desktop.MouseEvent) {
	if !o.selecting {
		return
	}
	o.selecting = false
	x1, y1 := int(o.start.X), int(o.start.Y)
	x2, y2 := int(o.curr.X), int(o.curr.Y)
	o.result <- image.Rect(min(x1, x2), min(y1, y2), max(x1, x2), max(y1, y2))
}

func (o *selectOverlay) MouseIn(_ *desktop.MouseEvent) {} // 必须实现
func (o *selectOverlay) MouseOut()                     {} // 必须实现

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
