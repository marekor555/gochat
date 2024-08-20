package global

import (
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Message struct {
	Name, Text string
}

func UpdateChat(messages []Message, chatBox *fyne.Container) {
	chatBox.Objects = []fyne.CanvasObject{}
	for _, message := range messages {
		chatBox.Objects = append(chatBox.Objects, widget.NewLabel(message.Name+": "+message.Text))
	}
}

var (
	Conn     net.Conn
	Messages []Message // raw message data

	Application fyne.App
	Window      fyne.Window

	ChatBox, MainBox, MenuBoxCent *fyne.Container // messages rendered here
	TextInput                     *widget.Entry
	ChatBoxScroll                 *container.Scroll

	Port       string = ":8080"
	ConnActive        = false
)
