package global

import (
	"image/color"
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Message struct {
	Name, Text string
}

func UpdateChat(messages []Message, chatBox *fyne.Container) {
	chatBox.Objects = []fyne.CanvasObject{}
	for _, message := range messages {
		if message.Name == "You" {
			chatBox.Objects = append(chatBox.Objects, container.NewHBox(canvas.NewText(message.Name+": ", color.RGBA{146, 235, 52, 255}), canvas.NewText(message.Text, color.White)))
		} else {
			chatBox.Objects = append(chatBox.Objects, container.NewHBox(canvas.NewText(message.Name+": ", color.RGBA{52, 192, 235, 255}), canvas.NewText(message.Text, color.White)))
		}
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
