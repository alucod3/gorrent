package cli

import (
	"fmt"
	"time"

	"github.com/alucod3/gorrent/pkg/utils"
	"github.com/schollz/progressbar/v3"
)

// ProgressUI manages the display of progress bars
type ProgressUI struct {
	metadataBar   *progressbar.ProgressBar
	downloadBar   *progressbar.ProgressBar
	lastBytes     int64
	lastTime      time.Time
	currentPeers  int
	currentSpeed  float64
	description   string
	totalSize     int64
	bytesComplete int64
}

// NewProgressUI creates a new progress interface
func NewProgressUI() *ProgressUI {
	return &ProgressUI{}
}

// ShowMetadataLoader displays a progress bar for loading metadata
func (p *ProgressUI) ShowMetadataLoader(description string) {
	p.metadataBar = progressbar.NewOptions(-1,
		progressbar.OptionSetDescription(description),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
		progressbar.OptionShowCount(),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionClearOnFinish(),
	)
}

// UpdateMetadataLoader updates the metadata loader
func (p *ProgressUI) UpdateMetadataLoader() {
	if p.metadataBar != nil {
		p.metadataBar.Add(1)
	}
}

// CompleteMetadataLoader completes the metadata loader
func (p *ProgressUI) CompleteMetadataLoader() {
	if p.metadataBar != nil {
		p.metadataBar.Finish()
		fmt.Println() // Add a line after completion
	}
}

// CreateDownloadBar creates a progress bar for the download
func (p *ProgressUI) CreateDownloadBar(total int64, description string) {
	p.currentPeers = 0
	p.currentSpeed = 0
	p.lastBytes = 0
	p.lastTime = time.Now()
	p.description = description
	p.totalSize = total
	p.bytesComplete = 0

	// Initial description
	initialDesc := fmt.Sprintf("%s | Iniciando...", description)

	p.downloadBar = progressbar.NewOptions(int(total),
		progressbar.OptionSetDescription(initialDesc),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(50),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "█",
			SaucerHead:    "█",
			SaucerPadding: "░",
			BarStart:      "[",
			BarEnd:        "]",
		}),
		// Important: Do not insert new lines when rendering
		progressbar.OptionUseANSICodes(true),
		// Make sure the bar stays on the same line
		progressbar.OptionFullWidth(),
		// Does not show automatic ETA statistics (we will show our own statistics)
		progressbar.OptionSetPredictTime(false),
		// Ensures the bar is cleared correctly when completed
		progressbar.OptionClearOnFinish(),
	)
}

// UpdateDownloadProgress updates download progress
func (p *ProgressUI) UpdateDownloadProgress(bytesCompleted int64) {
	if p.downloadBar != nil {
		p.bytesComplete = bytesCompleted
		p.downloadBar.Set(int(bytesCompleted))
	}
}

// DisplayDownloadStats updates statistics about the current download
func (p *ProgressUI) DisplayDownloadStats(bytesCompleted int64, peers int, totalSize int64) {
	// Calcula a velocidade
	currentTime := time.Now()
	elapsedTime := currentTime.Sub(p.lastTime).Seconds()

	// Avoid division by zero or very small intervals
	if elapsedTime > 0.1 {
		p.currentSpeed = float64(bytesCompleted-p.lastBytes) / elapsedTime
		p.lastBytes = bytesCompleted
		p.lastTime = currentTime
	}

	// Update the number of peers
	p.currentPeers = peers
	p.bytesComplete = bytesCompleted

	// Updates the bar description with the new information
	var description string

	if p.currentPeers == 0 && p.bytesComplete < p.totalSize {
		description = fmt.Sprintf("%s | ⚠️  Waiting for peers...", p.description)
	} else if p.bytesComplete < p.totalSize {
		description = fmt.Sprintf("%s | 📶 Peers: %d | 🚀 %s/s",
			p.description,
			p.currentPeers,
			utils.BytesToString(int64(p.currentSpeed)))
	} else {
		description = p.description + " | Completed!"
	}

	// Set the new description in the bar
	if p.downloadBar != nil {
		p.downloadBar.Describe(description)
	}
}

// CompleteDownloadBar finish the download progress bar
func (p *ProgressUI) CompleteDownloadBar() {
	if p.downloadBar != nil {
		p.downloadBar.Finish()
		fmt.Println() // Add a line after finish
	}
}
