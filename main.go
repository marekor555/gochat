package main

import (
	"gochat/global"
	"gochat/ui"
)

func main() {
	// initialize ui
	ui.InitUi()
	// set content to menu and run
	global.Window.SetContent(global.MenuBoxCent)
	global.Window.ShowAndRun()
}
