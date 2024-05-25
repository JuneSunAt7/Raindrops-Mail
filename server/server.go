package server

import(
	"github.com/pterm/pterm"
	"net"
	"io/ioutil"
	"os"
	"fmt"
)

func StartServer(){
    ln, err := net.Listen("tcp", ":25")
    if err != nil {
        pterm.Error.Println("Ошибка при запуске сервера:", err)
        return
    }
    defer ln.Close()

    pterm.Success.Println("Сервер успешно запущен и ожидает подключения")

    for {
        conn, err := ln.Accept()
        if err != nil {
            pterm.Error.Println("Ошибка при принятии подключения:", err)
            continue
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn){
    defer conn.Close()

    message, err := ioutil.ReadAll(conn)
    if err != nil {
        pterm.Error.Println("Ошибка при чтении сообщения:", err)
        return
    }

    pterm.Println("Получено сообщение:", string(message))
	writeMessageToLog(string(message))

}
func writeMessageToLog(message string){
	file, err := os.OpenFile("log.log", os.O_APPEND|os.O_WRONLY, 0644) 
	if err != nil { 
		fmt.Println("Error opening file:", err) 
		return
	 }
		
		defer file.Close()

_, err = file.WriteString(message + "\n")
if err != nil {
    fmt.Println("Error writing to file:", err)
    return
}
}