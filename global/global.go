package global

import (
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	Conn net.Conn

	Application fyne.App
	Window      fyne.Window

	ChatBox, MainBox, MenuBoxCent *fyne.Container // messages rendered here
	TextInput                     *widget.Entry
	ChatBoxScroll                 *container.Scroll
)
