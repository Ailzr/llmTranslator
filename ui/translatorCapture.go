package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/viper"
	"image"
	"image/color"
	"llmTranslator/logHelper"
	"llmTranslator/utils"
)

func (mw *MainWindow) CaptureRectangle() {
	img, err := utils.CaptureAllScreen()
	if err != nil {
		logHelper.Error(err.Error())
		logHelper.WriteLog(err.Error())
		return
	}

	w := mw.App.NewWindow("截屏")
	w.SetFullScreen(true)
	w.SetPadded(false)

	bg := canvas.NewImageFromImage(img)
	bg.FillMode = canvas.ImageFillStretch

	// 准备 overlay
	result := make(chan image.Rectangle)
	overlay := newSelectOverlay(result)

	w.SetContent(container.NewStack(bg, overlay))

	w.Canvas().(desktop.Canvas).SetOnKeyDown(func(e *fyne.KeyEvent) {
		switch e.Name {
		case fyne.KeyEscape:
			w.Close()
		}
	})
	w.Show()
	go func() {
		sel := <-result
		dialog.ShowConfirm("确认选区",
			fmt.Sprintf("是否使用坐标 %v ？", sel),
			func(ok bool) {
				if ok {
					w.Close()
					//将选取到的坐标保存起来
					viper.Set("capture.start_x", sel.Min.X)
					viper.Set("capture.start_y", sel.Min.Y)
					viper.Set("capture.end_x", sel.Max.X)
					viper.Set("capture.end_y", sel.Max.Y)
					err := viper.WriteConfig()
					if err != nil {
						logHelper.Error(err.Error())
						logHelper.WriteLog(err.Error())
					}
					mw.CreateShowWindow()
				} else {
					// 用户取消，只隐藏选区框，允许重新框选
					w.Close()
				}
			},
			w,
		)
	}()
}

type selectOverlay struct {
	widget.BaseWidget
	rect      *canvas.Rectangle
	start     fyne.Position
	selecting bool
	result    chan image.Rectangle
}

func newSelectOverlay(result chan image.Rectangle) *selectOverlay {
	o := &selectOverlay{
		rect:   canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 255, A: 64}),
		result: result,
	}
	o.rect.StrokeColor = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	o.rect.StrokeWidth = 2
	o.ExtendBaseWidget(o) // 注册为可接收事件的 Widget
	o.rect.Hide()
	return o
}

func (o *selectOverlay) CreateRenderer() fyne.WidgetRenderer {
	return &overlayRenderer{objects: []fyne.CanvasObject{o.rect}}
}

type overlayRenderer struct {
	objects []fyne.CanvasObject
}

func (r *overlayRenderer) Layout(size fyne.Size) {}
func (r *overlayRenderer) MinSize() fyne.Size    { return fyne.NewSize(0, 0) }
func (r *overlayRenderer) Refresh() {
	for _, o := range r.objects {
		o.Refresh()
	}
}
func (r *overlayRenderer) Objects() []fyne.CanvasObject { return r.objects }
func (r *overlayRenderer) Destroy()                     {}

func (o *selectOverlay) MouseDown(e *desktop.MouseEvent) {
	if e.Button != desktop.MouseButtonPrimary {
		return
	}
	o.start = e.Position
	o.selecting = true
	o.rect.Move(o.start)
	o.rect.Resize(fyne.NewSize(0, 0))
	o.rect.Show()
	o.Refresh()
}

func (o *selectOverlay) MouseIn(_ *desktop.MouseEvent) {} // 必须实现
func (o *selectOverlay) MouseOut()                     {} // 必须实现

func (o *selectOverlay) MouseMoved(e *desktop.MouseEvent) {
	if !o.selecting {
		return
	}
	x1, y1 := int(o.start.X), int(o.start.Y)
	x2, y2 := int(e.Position.X), int(e.Position.Y)
	minX, minY := float32(min(x1, x2)), float32(min(y1, y2))
	w, h := float32(abs(x2-x1)), float32(abs(y2-y1))

	o.rect.Move(fyne.NewPos(minX, minY))
	o.rect.Resize(fyne.NewSize(w, h))
	o.Refresh()
}

func (o *selectOverlay) MouseUp(e *desktop.MouseEvent) {
	if !o.selecting {
		return
	}
	o.selecting = false
	x1, y1 := int(o.start.X), int(o.start.Y)
	x2, y2 := int(e.Position.X), int(e.Position.Y)
	o.result <- image.Rect(min(x1, x2), min(y1, y2), max(x1, x2), max(y1, y2))
}

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
