package postclient

import(
	"github.com/pterm/pterm"
)
func TuiUseAddr(){
	pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)).
	WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println("Доступные адреса")
	
}
func readAddrFromRegedit(){

}