package logic

import (
	"gochat/global"
	"gochat/util"
	"log"
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func UpdateChat(messages []global.Message, chatBox *fyne.Container) {
	chatBox.Objects = []fyne.CanvasObject{}
	for _, message := range messages {
		chatBox.Objects = append(chatBox.Objects, widget.NewLabel(message.Name+": "+message.Text))
	}
}

func HandleIn() {
	go func() {
		for global.ConnActive {
			buffer := make([]byte, 128)
			_, err := global.Conn.Read(buffer) // TODO: remove blank message when reconnected
			util.CheckErr(err)
			global.Messages = append(global.Messages, global.Message{Name: global.Conn.RemoteAddr().String(), Text: string(buffer)})
			UpdateChat(global.Messages, global.ChatBox)
			global.ChatBoxScroll.ScrollToBottom()
		}
		log.Println("quitting input")
	}()
}

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	util.CheckErr(err)
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}
