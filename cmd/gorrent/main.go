// Package main implements a command-line torrent downloader application
// that provides a user-friendly interface for downloading torrents.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alucod3/gorrent/internal/cli"
	"github.com/alucod3/gorrent/internal/config"
	"github.com/alucod3/gorrent/internal/downloader"
	"github.com/alucod3/gorrent/internal/validator"
)

const usage = `Usage:
  gorrent                                   # Start interactive mode
  gorrent <magnet-link>                     # Start download with magnet link
  gorrent <path/to/file.torrent>           # Start download with torrent file`

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

	// Create validator
	v, err := validator.New()
	if err != nil {
		ui.ShowError("Error initializing validator", err)
		os.Exit(1)
	}

	// Get torrent link from args or prompt
	link, err := getTorrentLink(ui)
	if err != nil {
		if err == errShowUsage {
			fmt.Println(usage)
			os.Exit(0)
		}
		ui.ShowError("Error getting torrent link", err)
		os.Exit(1)
	}

	// Validate the link
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

var errShowUsage = fmt.Errorf("show usage")

// getTorrentLink returns a torrent link from command line args or prompts the user
func getTorrentLink(ui *cli.UI) (string, error) {
	args := os.Args[1:]

	// If no args provided, prompt for link
	if len(args) == 0 {
		return ui.ReadTorrentLink()
	}

	// If help flag provided, show usage
	if args[0] == "-h" || args[0] == "--help" {
		return "", errShowUsage
	}

	// If more than one arg provided, show usage
	if len(args) > 1 {
		return "", errShowUsage
	}

	// Return the provided link/path
	return args[0], nil
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
