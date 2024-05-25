package mgmtdata
import(
	"github.com/pterm/pterm"
	"bufio"
    "os"

)
func TuiReadAllMessages(){
 // Открываем файл для чтения
 file, err := os.Open("log.log")
 if err != nil {
	 pterm.Error.Println("Error opening file:", err)
	 return
 }
 defer file.Close()

 // Создаем новый сканер для файла
 scanner := bufio.NewScanner(file)

 // Читаем строки из файла и выводим их на экран
 for scanner.Scan() {
	 pterm.Success.Println(scanner.Text())
 }

 // Проверяем наличие ошибок во время сканирования
 if err := scanner.Err(); err != nil {
	 pterm.Error.Println("Error scanning file:", err)
 }
}