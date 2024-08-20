package logic

import (
	"gochat/global"
	"gochat/util"
	"log"
	"net"
)

func HandleIn() {
	go func() {
		for global.ConnActive {
			buffer := make([]byte, 128)
			n, err := global.Conn.Read(buffer) // TODO: remove blank message when reconnected
			util.CheckErr(err)
			log.Printf("Recieved: \"%v\",  %v", string(buffer), n)
			global.Messages = append(global.Messages, global.Message{Name: global.Conn.RemoteAddr().String(), Text: string(buffer)})
			global.UpdateChat(global.Messages, global.ChatBox)
			global.ChatBoxScroll.ScrollToBottom()
		}
		log.Println("quitting input")
	}()
}

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	util.CheckErr(err)
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}
