package logic

import (
	"gochat/global"
	"log"
	"net"
	"strings"
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
	// fyne.CurrentApp().SendNotification(fyne.NewNotification("Gochat", "Connection succesfull to: "+conn.RemoteAddr().String()))
	log.Println("Connected succesfully")
	global.Messages = []global.Message{}
	global.ConnActive = true
	return conn, nil
}
