package postclient
import(

	"github.com/pterm/pterm"
	"golang.org/x/sys/windows/registry"

	"encoding/base64"
)
type Addr struct { 
	Address string
	Port int 
}
func TuiCreateAddr(){
	pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)).
	WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println("Создание адреса")
	addr, _ := pterm.DefaultInteractiveTextInput.Show("IP: ")
	createAddr(addr)

}
func createAddr(address string) {
	// create default port
	addr := Addr{ Address: address, Port: 2048 }
	writeToRegedit(addr)
}
func writeToRegedit(address Addr ){
	CreateSettingsToRegedit("addr", base64.StdEncoding.EncodeToString([]byte(address.Address)))
}

func CreateSettingsToRegedit(settingsName string, settingValue string) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\RaindropsMail`, registry.ALL_ACCESS)
	if err != nil {
		k, _, err = registry.CreateKey(registry.CURRENT_USER, `Software\RaindropsMail`, registry.ALL_ACCESS)
		if err != nil {
			return
		}
	}

	defer k.Close()

	err = k.SetStringValue(settingsName, settingValue)
	if err != nil {
		return
	}
}
func ReadRegistryValue(key registry.Key, subKey string, valueName string) (string, error) {
	k, err := registry.OpenKey(key, subKey, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	val, _, err := k.GetStringValue(valueName)
	if err != nil {
		return "", err
	}

	return val, nil
}