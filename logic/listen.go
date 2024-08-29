package logic

import (
	"gochat/global"
	"log"
	"net"
)

func Listen() (net.Conn, error) {
	listener, err := net.Listen("tcp", global.Port)
	if err != nil {
		return nil, err
	}
	defer listener.Close()
	conn, err := listener.Accept()
	if err != nil {
		return nil, err
	}
	// fyne.CurrentApp().SendNotification(fyne.NewNotification("Gochat", "Connection succesfull to: "+conn.RemoteAddr().String()))
	log.Println("Connection found succesfully")
	global.Messages = []global.Message{}
	global.ConnActive = true
	buff := make([]byte, 1024)
	conn.Read(buff)
	conn.Write([]byte(global.NameEntry.Text))
	global.MessengerName = string(buff)
	return conn, nil
}
