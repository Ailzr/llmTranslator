package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"llmTranslator/configs"
	"net"
	"os"
)

const ipcPort = "127.0.0.1:23258"

type MainWindow struct {
	App              fyne.App     //应用程序
	Listener         net.Listener //防止重复启动应用的监听器
	Window           fyne.Window  //主窗口
	TranslatorWindow fyne.Window  //翻译展示窗口
	CaptureWindow    fyne.Window  //截图窗口
	isTray           bool         //是否启用系统托盘
}

var mw = &MainWindow{}

func init() {

	//新建app和main window
	mw.App = app.New()
	//mw.App.Settings().SetTheme(&customTheme{})
	mw.Window = mw.App.NewWindow("llmTranslator")
	if configs.Setting.DefaultTray {
		mw.Window.Hide()
		mw.isTray = true
	} else {
		mw.Window.Show()
		mw.isTray = false
	}

	//设置为主窗口
	mw.Window.SetMaster()

	//设置窗口大小和图标
	mw.Window.Resize(fyne.NewSize(800, 600))
	mw.Window.SetIcon(resourceLlmTranslatorPng)

	//监听ipc端口
	var err error
	mw.Listener, err = net.Listen("tcp", ipcPort)
	if err != nil {
		// 已有实例运行，发送激活信号并退出
		sendActivateSignal()
		os.Exit(0)
	}

	//监听激活信号
	go listenForActivateSignal(mw)

	//设置内容
	tabs := container.NewAppTabs(
		container.NewTabItem(
			"主页",
			container.NewVBox(
				widget.NewLabel("主页"),
				createHomeForm(mw),
			),
		),
		container.NewTabItem(
			"OCR",
			container.NewVBox(
				widget.NewLabel("OCR设置"),
				createOCRForm(),
			),
		),
		container.NewTabItem(
			"LLM",
			container.NewVBox(
				widget.NewLabel("LLM设置"),
				createLLMForm(),
			),
		),
		container.NewTabItem(
			"应用设置",
			container.NewVBox(
				widget.NewLabel("应用设置"),
				createAppSettingForm(),
			),
		),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	mw.Window.SetContent(tabs)

	mw.CreateShowWindow()

	//允许系统托盘
	startTray(mw)

	//添加热键
	RegisterAllHotKey()
}

func GetMainWindow() *MainWindow {
	return mw
}

func ShowAndRun() {
	mw.App.Run()
}
