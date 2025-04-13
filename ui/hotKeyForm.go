package ui

import (
	"fyne.io/fyne/v2/widget"
)

func createHotKeyForm() *widget.Form {
	var form *widget.Form

	form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "测试", Widget: widget.NewLabel("测试")},
		},
	}

	return form
}
