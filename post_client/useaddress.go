package postclient

import(
	"github.com/pterm/pterm"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"encoding/base64"
)
func TuiUseAddr(){
	pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)).
	WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println("Доступные адреса")

	pterm.FgLightBlue.Println(ReadAddrFromRegedit())

}
func ReadAddrFromRegedit() string{
	key := registry.CURRENT_USER
	subKey := `Software\RaindropsMail`
	valueName := "addr"
	val, err := ReadRegistryValue(key, subKey, valueName)
	if err != nil {
		fmt.Printf("Error reading registry value: %v", err)
	} 
	
	data, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		fmt.Println("Ошибка декодирования base64:", err)
		return ""
	}

	decodedString := string(data)
	return decodedString
}