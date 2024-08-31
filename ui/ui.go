package ui

import (
	"gochat/global"
	"gochat/logic"
	"gochat/ui/layouts"
	"gochat/util"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func InitUi() {
	global.Application = app.New()
	global.Window = global.Application.NewWindow("Gochat")
	global.Window.Resize(fyne.NewSize(0, 512))

	// chat gui
	global.TextInput = widget.NewEntry()
	sendBtn := widget.NewButton("Send", func() {
		if strings.TrimSpace(global.TextInput.Text) == "" {
			return
		}
		if len(global.TextInput.Text) >= 2048 {
			global.TextInput.Text = global.TextInput.Text[0:2048]
		}
		global.Conn.Write([]byte(global.TextInput.Text))
		global.Messages = append(global.Messages, global.Message{Name: "You", Text: global.TextInput.Text})
		global.UpdateChat(global.Messages, global.ChatBox)
		global.TextInput.SetText("")
		global.ChatBoxScroll.ScrollToBottom()
	})
	inputBox := container.New(&layouts.InputLayout{}, global.TextInput, sendBtn)

	global.ChatBox = container.NewVBox()
	global.ChatBoxScroll = container.NewScroll(global.ChatBox)

	global.MainBox = container.New(&layouts.MainLayout{}, container.NewHBox(widget.NewButton("Quit", func() {
		global.Window.SetContent(global.MenuBoxCent)
		global.Conn.Close()
		global.ConnActive = false
		global.Messages = []global.Message{}
		global.UpdateChat(global.Messages, global.ChatBox)
	})), global.ChatBoxScroll, inputBox)

	// main menu
	menuLabel := widget.NewLabel("Gochat, safe chat with p2p")
	menuLabelWrap := container.NewHBox(layout.NewSpacer(), menuLabel, layout.NewSpacer())

	ipInputLabel := widget.NewLabel("IP:")
	ipInputEntry := widget.NewEntry()
	ipInputWrap := container.New(&layouts.EntryLayout{}, ipInputLabel, ipInputEntry)

	nameLabel := widget.NewLabel("Name:")
	global.NameEntry = widget.NewEntry()
	global.NameEntry.Text = "Unknown"
	nameWrap := container.New(&layouts.EntryLayout{}, nameLabel, global.NameEntry)

	ipConfirmBtn := widget.NewButton("Connect", func() {
		global.Window.SetContent(container.NewCenter(container.NewVBox(widget.NewLabel("Waiting..."))))
		go func() {
			newConnection, err := logic.Connect(ipInputEntry.Text)
			util.CheckErr(err)
			if newConnection == nil {
				return
			}
			global.Conn = newConnection
			logic.HandleIn()
			global.Window.SetContent(global.MainBox)
		}()
	})

	listenBtn := widget.NewButton("Listen", func() {
		global.Window.SetContent(container.NewCenter(container.NewVBox(widget.NewLabel("Waiting..."), widget.NewLabel("IP:"+logic.GetLocalIP().String()+global.Port))))
		go func() {
			newConnection, err := logic.Listen()
			util.CheckErr(err)
			if newConnection == nil {
				return
			}
			global.Conn = newConnection
			logic.HandleIn()
			global.Window.SetContent(global.MainBox)
		}()
	})

	menuBtnGrid := container.NewGridWithColumns(2, ipConfirmBtn, listenBtn)

	menuBox := container.NewVBox(menuLabelWrap, nameWrap, ipInputWrap, menuBtnGrid)
	global.MenuBoxCent = container.NewCenter(menuBox)

	log.Println("initialized ui, running")
	// set content to menu and run
	global.Window.SetContent(global.MenuBoxCent)
	global.Window.ShowAndRun()
}
