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
	buff := make([]byte, 1024)
	conn.Read(buff)
	global.MessengerName = string(buff)
	conn.Write([]byte(global.NameEntry.Text))
	return conn, nil
}
