package cli

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// clearTerminal limpa a tela do terminal
func clearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// displayLogo exibe o logo do aplicativo
func displayLogo(colors *ColorScheme) {
	const (
		logo = `
   ██████╗  ██████╗ ██████╗ ██████╗ ███████╗███╗   ██╗████████╗
  ██╔════╝ ██╔═══██╗██╔══██╗██╔══██╗██╔════╝████╗  ██║╚══██╔══╝
  ██║  ███╗██║   ██║██████╔╝██████╔╝█████╗  ██╔██╗ ██║   ██║   
  ██║   ██║██║   ██║██╔══██╗██╔══██╗██╔══╝  ██║╚██╗██║   ██║   
  ╚██████╔╝╚██████╔╝██║  ██║██║  ██║███████╗██║ ╚████║   ██║   
   ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝   
`
		appVersion = "0.1"
	)

	colors.Title.Println(logo)
	colors.Subtitle.Printf("                       Versão %s\n\n", appVersion)
	fmt.Println("  A simple, fast and efficient CLI torrent downloader.")
	fmt.Println("  -------------------------------------------------------")
	fmt.Println()
}
