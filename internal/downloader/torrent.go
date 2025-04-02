package downloader

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/alucod3/gorrent/internal/cli"
	"github.com/alucod3/gorrent/internal/config"
	"github.com/alucod3/gorrent/pkg/utils"
	"github.com/anacrolix/torrent"
)

// TorrentDownloader gerencia o download de torrents
type TorrentDownloader struct {
	config   *config.Config
	progress *cli.ProgressUI
	client   *torrent.Client
}

// New cria um novo gerenciador de downloads
func New(cfg *config.Config, progress *cli.ProgressUI) *TorrentDownloader {
	return &TorrentDownloader{
		config:   cfg,
		progress: progress,
	}
}

// Download inicia o download de um torrent
func (d *TorrentDownloader) Download(ctx context.Context, link string) error {
	// Garantir que o diretório de download existe
	if err := d.config.EnsureDownloadPath(); err != nil {
		return fmt.Errorf("erro ao criar diretório de download: %w", err)
	}

	// Configurar o cliente torrent
	config := torrent.NewDefaultClientConfig()
	config.DataDir = d.config.DownloadPath
	config.Seed = d.config.Seed

	client, err := torrent.NewClient(config)
	if err != nil {
		return fmt.Errorf("erro ao criar cliente torrent: %w", err)
	}
	defer client.Close()

	d.client = client

	var t *torrent.Torrent

	// Adicionar o torrent baseado no tipo de link
	t, err = d.addTorrent(link)
	if err != nil {
		return err
	}

	// Obter metadados
	if err := d.fetchMetadata(ctx, t); err != nil {
		return err
	}

	// Exibir informações
	d.displayTorrentInfo(t)

	// Iniciar o download
	return d.startDownload(ctx, t)
}

// addTorrent adiciona um torrent baseado no tipo de entrada (arquivo local, magnet, etc)
func (d *TorrentDownloader) addTorrent(link string) (*torrent.Torrent, error) {
	if _, err := os.Stat(link); err == nil {
		// É um arquivo local
		return d.client.AddTorrentFromFile(link)
	} else if strings.HasPrefix(link, "magnet:") {
		// É um magnet link
		return d.client.AddMagnet(link)
	} else {
		// URL não suportada
		return nil, fmt.Errorf("URLs diretas de .torrent não são suportadas nesta versão. Por favor, use um magnet link ou arquivo .torrent local")
	}
}

// fetchMetadata obtém os metadados do torrent
func (d *TorrentDownloader) fetchMetadata(ctx context.Context, t *torrent.Torrent) error {
	// Mostrar barra de progresso para carregamento de metadados
	d.progress.ShowMetadataLoader("Carregando metadados")

	// Contexto para animação de loading
	metadataCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Animação de loading
	go func() {
		for {
			select {
			case <-metadataCtx.Done():
				return
			default:
				d.progress.UpdateMetadataLoader()
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// Aguardar os metadados
	select {
	case <-t.GotInfo():
		cancel()
		d.progress.CompleteMetadataLoader()
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// displayTorrentInfo exibe informações sobre o torrent
func (d *TorrentDownloader) displayTorrentInfo(t *torrent.Torrent) {
	// Criar um serviço de UI aqui e usá-lo para exibir as informações
	ui := cli.NewUI()
	ui.DisplayTorrentInfo(
		t.Name(),
		utils.BytesToString(t.Length()),
		strconv.Itoa(len(t.Files())),
		filepath.Join(d.config.DownloadPath, t.Name()),
	)
}

// startDownload inicia o download do torrent
func (d *TorrentDownloader) startDownload(ctx context.Context, t *torrent.Torrent) error {
	// Iniciar o download
	t.DownloadAll()

	// Criar barra de progresso
	d.progress.CreateDownloadBar(t.Length(), "Baixando")

	// Monitorar o progresso
	ticker := time.NewTicker(d.config.ProgressCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			stats := t.Stats()
			bytesCompleted := t.BytesCompleted()

			// Atualizar a barra de progresso
			d.progress.UpdateDownloadProgress(bytesCompleted)

			// Exibir estatísticas
			d.progress.DisplayDownloadStats(bytesCompleted, stats.ActivePeers, t.Length())

			// Verificar se o download está completo
			if bytesCompleted == t.Length() {
				d.progress.CompleteDownloadBar()
				fmt.Println()
				return nil
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
