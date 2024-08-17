package logic

import (
	"fyne.io/fyne/v2"
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
