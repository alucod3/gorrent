// Package main implements a command-line torrent downloader application
// that provides a user-friendly interface for downloading torrents.
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alucod3/gorrent/internal/cli"
	"github.com/alucod3/gorrent/internal/config"
	"github.com/alucod3/gorrent/internal/downloader"
	"github.com/alucod3/gorrent/internal/validator"
)

func main() {
	// Initialize the UI
	ui := cli.NewUI()
	ui.ClearScreen()
	ui.ShowLogo()

	// Load settings
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		ui.ShowError("Error loading configuration", err)
		os.Exit(1)
	}

	// Create cancelable context to manage lifecycle
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Configure signal capture for interrupt
	setupSignalHandler(cancel, ui)

	// Read and validate the torrent link
	link, err := ui.ReadTorrentLink()
	if err != nil {
		ui.ShowError("Error reading input", err)
		os.Exit(1)
	}

	// Create validator and validate the link
	v, err := validator.New()
	if err != nil {
		ui.ShowError("Error initializing validator", err)
		os.Exit(1)
	}

	if err := v.IsValidTorrentLink(link); err != nil {
		ui.ShowError("Invalid link", err)
		os.Exit(1)
	}

	ui.ShowSuccess("Valid link! Preparing download...")

	// Start download
	dl := downloader.New(cfg, ui.ProgressTracker())
	if err := dl.Download(ctx, link); err != nil {
		if err == context.Canceled {
			os.Exit(0)
		}
		ui.ShowError("Error during download", err)
		os.Exit(1)
	}

	// Short pause for user to see completion message
	ui.PauseForUserFeedback()
}

// setupSignalHandler configures signal handling for interrupt
func setupSignalHandler(cancel context.CancelFunc, ui *cli.UI) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		ui.ShowWarning("Stopping download...")
		cancel()
	}()
}
