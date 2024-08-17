package main

import (
	"net"

	"gochat/layouts"
	"gochat/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	chatBox, mainBox, menuBoxCent *fyne.Container // messages rendered here
	textInput                     *widget.Entry
	chatBoxScroll                 *container.Scroll
	application                   fyne.App
	window                        fyne.Window
	messages                      []logic.Message // raw message data
	conn                          net.Conn
	err                           error
)

func checkErr(err error) {
	if err != nil {
		if conn != nil {
			conn.Close()
		}
		application.SendNotification(fyne.NewNotification("ERROR!", err.Error()))
		application.Quit()
	}
}

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	checkErr(err)
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}

func handleIn() {
	go func() {
		for {
			buffer := make([]byte, 128)
			_, err := conn.Read(buffer)
			checkErr(err)
			messages = append(messages, logic.Message{Name: conn.RemoteAddr().String(), Text: string(buffer)})
			logic.UpdateChat(messages, chatBox)
			chatBoxScroll.ScrollToBottom()
		}
	}()
}

func main() {
	application = app.New()
	window = application.NewWindow("Gochat")
	window.Resize(fyne.NewSize(512, 512))

	// chat gui
	textInput = widget.NewEntry()
	sendBtn := widget.NewButton("Send", func() {
		conn.Write([]byte(textInput.Text))
		messages = append(messages, logic.Message{Name: "You", Text: textInput.Text})
		logic.UpdateChat(messages, chatBox)
		textInput.SetText("")
	})
	inputBox := container.New(&layouts.InputLayout{}, textInput, sendBtn)

	chatBox = container.NewVBox()
	chatBoxScroll = container.NewScroll(chatBox)

	mainBox = container.New(&layouts.MainLayout{}, chatBoxScroll, inputBox)

	// main menu
	menuLabel := widget.NewLabel("Gochat, safe chat with p2p")
	ipInput := widget.NewEntry()

	ipConfirmBtn := widget.NewButton("Connect", func() {
		conn, err = logic.Connect(ipInput.Text)
		checkErr(err)
		handleIn()
		window.SetContent(mainBox)
	})

	listenBtn := widget.NewButton("Listen", func() {
		window.SetContent(container.NewCenter(container.NewVBox(widget.NewLabel("Waiting..."), widget.NewLabel(GetLocalIP().String()))))
		go func() {
			conn, err = logic.Listen()
			checkErr(err)
			handleIn()
			window.SetContent(mainBox)
		}()
	})

	menuBtnGrid := container.NewGridWithColumns(2, ipConfirmBtn, listenBtn)

	menuBox := container.NewVBox(menuLabel, ipInput, menuBtnGrid)
	menuBoxCent = container.NewCenter(menuBox)

	logic.UpdateChat(messages, chatBox)

	// set content to menu and run
	window.SetContent(menuBoxCent)
	window.ShowAndRun()
}
