package util

import (
	"errors"
	"gochat/global"
	"io"
	"log"
	"strings"
	"syscall"
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func gracefullFail(message string) {
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
		log.Println("ERROR: ", err.Error())
		if err == io.EOF {
			gracefullFail("Disconnected")
		} else if errors.Is(err, syscall.ECONNREFUSED) {
			gracefullFail("Connection refused")
		} else if strings.Contains(err.Error(), "use of closed network connection") {
			gracefullFail("Closed")
		} else {
			gracefullFail("Unknown error")
		}
	}
}
