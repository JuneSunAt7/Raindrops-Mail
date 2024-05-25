package main

import(
	"fmt"
	"github.com/pterm/pterm"
	"atomicgo.dev/keyboard/keys"
	postclient "github.com"
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
			
		case "Отправить сообщение":
			
		case "Использовать адрес":
			
		case "Конфигурация":
			
		case "Аутенфикация и пароли":
			
		case "Все сообщения":
			
		case "Выход":
			return
		}
	}

}
func main(){
tui()
}