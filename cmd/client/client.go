package main

import (
	"fmt"

	"atomicgo.dev/keyboard/keys"
	"github.com/pterm/pterm"

	postclient "github.com/JuneSunAt7/Raindrops-Mail/post_client"
	sender "github.com/JuneSunAt7/Raindrops-Mail/sender"
	server "github.com/JuneSunAt7/Raindrops-Mail/server"
	mgmtdata "github.com/JuneSunAt7/Raindrops-Mail/mgmtdata"
	
)
func tui(){
	var options []string

	options = append(options, fmt.Sprintf("Создать адрес"))
	options = append(options, fmt.Sprintf("Использовать адрес"))
	options = append(options, fmt.Sprintf("Отправить сообщение"))
	options = append(options, fmt.Sprintf("Конфигурация"))
	options = append(options, fmt.Sprintf("Аутенфикация и пароли"))
	options = append(options, fmt.Sprintf("Все сообщения"))
	options = append(options, fmt.Sprintf("Выход"))

	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.TextStyle.Add(*pterm.NewStyle(pterm.FgBlue))
	printer.KeyConfirm = keys.Enter

	for {
		selectedOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

		switch selectedOptions {
		case "Создать адрес":
			postclient.TuiCreateAddr()
		case "Отправить сообщение":
			sender.TuiTextArea()
		case "Использовать адрес":
			postclient.TuiUseAddr()
		case "Конфигурация":
			
		case "Аутенфикация и пароли":
			
		case "Все сообщения":
			mgmtdata.TuiReadAllMessages()
		case "Выход":
			return
		}
	}

}
func main(){
	go server.StartServer()
	tui()
}