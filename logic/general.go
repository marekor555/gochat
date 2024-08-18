package logic

import (
	"gochat/global"
	"gochat/util"
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var (
	Messages []Message // raw message data
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

func HandleIn() {
	go func() {
		for {
			buffer := make([]byte, 128)
			_, err := global.Conn.Read(buffer)
			util.CheckErr(err)
			Messages = append(Messages, Message{Name: global.Conn.RemoteAddr().String(), Text: string(buffer)})
			UpdateChat(Messages, global.ChatBox)
			global.ChatBoxScroll.ScrollToBottom()
		}
	}()
}

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	util.CheckErr(err)
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}
