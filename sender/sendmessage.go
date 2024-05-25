package sender

import(
	"github.com/pterm/pterm"
	"net"
	post_client "github.com/JuneSunAt7/Raindrops-Mail/post_client"
)

func TuiTextArea(){
	result, _ := pterm.DefaultInteractiveTextInput.Show()
	pterm.Print()
	sendMessage(result)

}
func sendMessage(msg string){
	fulladdr := post_client.ReadAddrFromRegedit() + ":" + "25"

	conn, err := net.Dial("tcp", fulladdr)
    if err != nil {
        pterm.Error.Println("Ошибка при подключении:", err)
        return
    }
    defer conn.Close()

    message := []byte(msg)
    _, err = conn.Write(message)
    if err != nil {
        pterm.Error.Println("Ошибка при отправке сообщения:", err)
        return
    }

	
    pterm.Success.Println("Сообщение успешно отправлено на адрес и порт")
}