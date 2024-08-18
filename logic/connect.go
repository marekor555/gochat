package logic

import (
	"gochat/global"
	"net"
	"strings"

	"fyne.io/fyne/v2"
)

func Connect(ip string) (net.Conn, error) {
	if ip == "" {
		ip = "localhost"
	}
	if !strings.Contains(ip, global.Port) {
		ip = ip + global.Port
	}
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		return nil, err
	}
	fyne.CurrentApp().SendNotification(fyne.NewNotification("Gochat", "Connection succesfull to: "+conn.RemoteAddr().String()))
	return conn, nil
}
