package logic

import (
	"net"

	"fyne.io/fyne/v2"
)

func Listen() (net.Conn, error) {
	listener, err := net.Listen("tcp", ":8080")
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
