package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type customTheme struct {
}

func (c customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	//setTheme := viper.GetString("ui.theme")

	//switch name {
	////背景色
	//case theme.ColorNameBackground:
	//	if setTheme == "dark" {
	//		return color.Black
	//	}
	//	if setTheme == "light" {
	//		return color.White
	//	}
	////字体颜色
	//case theme.ColorNameForeground:
	//	if setTheme == "dark" {
	//		return color.White
	//	}
	//	if setTheme == "light" {
	//		return color.Black
	//	}
	////其余使用默认颜色
	//default:
	//	if setTheme == "dark" {
	//		return color.RGBA{R: 0, G: 0, B: 0, A: 255}
	//	}
	//	if setTheme == "light" {
	//		return color.RGBA{R: 255, G: 255, B: 255, A: 255}
	//	}
	//}
	//TODO 亮暗主题
	return theme.DefaultTheme().Color(name, variant)
}

func (c customTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (c customTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (c customTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
