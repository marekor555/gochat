package util

import (
	"errors"
	"gochat/global"
	"io"
	"log"
	"syscall"
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func gracefullFail(message string) {
	log.Println("ERROR: ", message)
	global.Window.SetContent(container.NewCenter(widget.NewLabel(message)))
	time.Sleep(time.Second * 2)
	if global.ConnActive {
		global.Conn.Close()
		global.ConnActive = false
		global.Messages = []global.Message{}
		global.UpdateChat(global.Messages, global.ChatBox)
	}
	global.Window.SetContent(global.MenuBoxCent)
}

func CheckErr(err error) {
	if err != nil {
		if err == io.EOF {
			gracefullFail("Disconnected")
		} else if errors.Is(err, syscall.ECONNREFUSED) {
			gracefullFail("Connection refused")
		} else {
			gracefullFail(err.Error())
		}
	}
}
