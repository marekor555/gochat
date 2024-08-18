package logic

import (
	"gochat/global"
	"net"

	"fyne.io/fyne/v2"
)

func Listen() (net.Conn, error) {
	listener, err := net.Listen("tcp", global.Port)
	if err != nil {
		return nil, err
	}
	conn, err := listener.Accept()
	if err != nil {
		return nil, err
	}
	fyne.CurrentApp().SendNotification(fyne.NewNotification("Gochat", "Connection succesfull to: "+conn.RemoteAddr().String()))
	return conn, nil
}
