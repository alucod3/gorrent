package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// UI encapsula toda a lógica da interface com o usuário
type UI struct {
	colors     *ColorScheme
	reader     *bufio.Reader
	progressUI *ProgressUI
}

// NewUI cria uma nova instância da interface com usuário
func NewUI() *UI {
	return &UI{
		colors:     NewColorScheme(),
		reader:     bufio.NewReader(os.Stdin),
		progressUI: NewProgressUI(),
	}
}

// ClearScreen limpa a tela do terminal
func (ui *UI) ClearScreen() {
	clearTerminal()
}

// ShowLogo exibe o logo do aplicativo
func (ui *UI) ShowLogo() {
	displayLogo(ui.colors)
}

// ReadTorrentLink solicita e lê um link de torrent do usuário
func (ui *UI) ReadTorrentLink() (string, error) {
	ui.colors.Prompt.Print("🔗 Digite o link ou arraste o arquivo .torrent: ")
	input, err := ui.reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("erro ao ler entrada: %w", err)
	}

	// Remove espaços em branco e possíveis aspas de arquivos arrastados
	input = strings.TrimSpace(input)
	input = strings.Trim(input, "\"'")

	return input, nil
}

// ShowError exibe uma mensagem de erro
func (ui *UI) ShowError(message string, err error) {
	ui.colors.Error.Printf("❌ %s: %v\n", message, err)
}

// ShowSuccess exibe uma mensagem de sucesso
func (ui *UI) ShowSuccess(message string) {
	ui.colors.Success.Printf("✅ %s\n", message)
}

// ShowWarning exibe uma mensagem de aviso
func (ui *UI) ShowWarning(message string) {
	ui.colors.Warning.Printf("⚠️  %s\n", message)
}

// ShowInfo exibe uma mensagem informativa
func (ui *UI) ShowInfo(message string) {
	ui.colors.Info.Printf("ℹ️  %s\n", message)
}

// ProgressTracker retorna o gerenciador de progresso
func (ui *UI) ProgressTracker() *ProgressUI {
	return ui.progressUI
}

// PauseForUserFeedback pausa brevemente para o usuário ler o feedback
func (ui *UI) PauseForUserFeedback() {
	time.Sleep(1 * time.Second)
}

// DisplayTorrentInfo exibe informações detalhadas sobre um torrent
func (ui *UI) DisplayTorrentInfo(name, size, files, path string) {
	fmt.Println()
	ui.colors.Info.Println("📝 Informações do Torrent:")
	ui.colors.Highlight.Printf("   Nome: ")
	fmt.Println(name)
	ui.colors.Highlight.Printf("   Tamanho: ")
	fmt.Println(size)
	ui.colors.Highlight.Printf("   Arquivos: ")
	fmt.Println(files)
	ui.colors.Highlight.Printf("   Salvando em: ")
	fmt.Println(path)
	fmt.Println()
}
