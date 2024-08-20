package util

import (
	"gochat/global"
	"io"
	"log"
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CheckErr(err error) {
	if err != nil {
		log.Println("ERROR! :", err.Error())
		if err == io.EOF {
			log.Println("Disconnected")
			global.Window.SetContent(container.NewCenter(widget.NewLabel("Disconnected")))
			time.Sleep(time.Second * 2)
			global.Conn.Close()
			global.ConnActive = false
			global.Messages = []global.Message{}
			global.UpdateChat(global.Messages, global.ChatBox)
			global.Window.SetContent(global.MenuBoxCent)
			return
		}
		if global.Conn != nil {
			global.Conn.Close()
		}
		log.Println("Couldn't handle error quiting...")
		global.Application.Quit()
	}
}
