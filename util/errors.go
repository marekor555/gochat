package util

import (
	"gochat/global"

	"fyne.io/fyne/v2"
)

func CheckErr(err error) {
	if err != nil {
		if global.Conn != nil {
			global.Conn.Close()
		}
		global.Application.SendNotification(fyne.NewNotification("ERROR!", err.Error()))
		global.Application.Quit()
	}
}
