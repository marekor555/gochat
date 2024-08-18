package ui

import (
	"gochat/global"
	"gochat/logic"
	"gochat/ui/layouts"
	"gochat/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InitUi() {
	global.Application = app.New()
	global.Window = global.Application.NewWindow("Gochat")
	global.Window.Resize(fyne.NewSize(512, 512))

	// chat gui
	global.TextInput = widget.NewEntry()
	sendBtn := widget.NewButton("Send", func() {
		global.Conn.Write([]byte(global.TextInput.Text))
		logic.Messages = append(logic.Messages, logic.Message{Name: "You", Text: global.TextInput.Text})
		logic.UpdateChat(logic.Messages, global.ChatBox)
		global.TextInput.SetText("")
	})
	inputBox := container.New(&layouts.InputLayout{}, global.TextInput, sendBtn)

	global.ChatBox = container.NewVBox()
	global.ChatBoxScroll = container.NewScroll(global.ChatBox)

	global.MainBox = container.New(&layouts.MainLayout{}, global.ChatBoxScroll, inputBox)

	// main menu
	menuLabel := widget.NewLabel("Gochat, safe chat with p2p")
	ipInput := widget.NewEntry()

	ipConfirmBtn := widget.NewButton("Connect", func() {
		newConnection, err := logic.Connect(ipInput.Text)
		util.CheckErr(err)
		global.Conn = newConnection
		logic.HandleIn()
		global.Window.SetContent(global.MainBox)
	})

	listenBtn := widget.NewButton("Listen", func() {
		global.Window.SetContent(container.NewCenter(container.NewVBox(widget.NewLabel("Waiting..."), widget.NewLabel("IP:"+logic.GetLocalIP().String()+global.Port))))
		go func() {
			newConn, err := logic.Listen()
			util.CheckErr(err)
			global.Conn = newConn
			logic.HandleIn()
			global.Window.SetContent(global.MainBox)
		}()
	})

	menuBtnGrid := container.NewGridWithColumns(2, ipConfirmBtn, listenBtn)

	menuBox := container.NewVBox(menuLabel, ipInput, menuBtnGrid)
	global.MenuBoxCent = container.NewCenter(menuBox)

	logic.UpdateChat(logic.Messages, global.ChatBox)
}
