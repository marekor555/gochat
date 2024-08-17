package logic

import (
	"net"

	"fyne.io/fyne/v2"
)

func Connect(ip string) (net.Conn, error) {
	conn, err := net.Dial("tcp", ip+":8080")
	if err != nil {
		return nil, err
	}
	fyne.CurrentApp().SendNotification(fyne.NewNotification("Gochat", "Connection succesfull to: "+ip))
	return conn, nil
}
