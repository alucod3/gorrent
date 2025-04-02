package validator

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/alucod3/gorrent/internal/config"
)

// Validator verifica a validade de links e arquivos de torrent
type Validator struct {
	config *config.Config
}

// New cria um novo validador com configuração padrão.
// Se houver erro ao carregar a configuração, retorna nil e o erro.
func New() (*Validator, error) {
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar configuração: %w", err)
	}
	return &Validator{
		config: cfg,
	}, nil
}

// WithConfig cria um validador com configurações personalizadas
func WithConfig(cfg *config.Config) *Validator {
	return &Validator{
		config: cfg,
	}
}

// IsValidTorrentLink verifica se o link fornecido é um magnet link ou um link .torrent
func (v *Validator) IsValidTorrentLink(link string) error {
	// Verifica se o link está vazio
	if link == "" {
		return errors.New("link não pode estar vazio")
	}

	// Expressão regular para validar magnet links
	magnetRegex := regexp.MustCompile(v.config.MagnetPattern)

	// Verifica se é um magnet link
	if magnetRegex.MatchString(link) {
		return nil
	}

	// Verifica se é um arquivo local
	if _, err := os.Stat(link); err == nil {
		if strings.HasSuffix(strings.ToLower(link), v.config.TorrentExtension) {
			return nil
		}
		return fmt.Errorf("arquivo deve ter extensão %s", v.config.TorrentExtension)
	}

	// Verifica se é uma URL válida
	if u, err := url.Parse(link); err == nil && u.Scheme != "" && u.Host != "" {
		if strings.HasSuffix(strings.ToLower(u.Path), v.config.TorrentExtension) {
			return nil
		}
		return fmt.Errorf("URL deve apontar para um arquivo %s", v.config.TorrentExtension)
	}

	return errors.New("link inválido: deve ser um magnet link, arquivo local ou URL para torrent")
}
