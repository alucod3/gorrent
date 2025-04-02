# 🚀 GoRrent - Torrent Manager

**GoRrent** is a simple, fast, and efficient torrent manager developed in Go. Download torrent files from magnet links or local .torrent files with a minimalist and powerful experience.

![GoRrent Demo](https://raw.githubusercontent.com/alucod3/gorrent/main/docs/demo.gif)

## ✨ Features

- **Friendly CLI interface** - Simplified command-line experience
- **Complete support** - Works with magnet links and local .torrent files
- **Real-time progress** - Track your downloads with live updates
- **Smart formatting** - Clear display of sizes in KB, MB, GB
- **Modern visual** - Visual feedback with colors and emojis for a better experience
- **Safe operation** - Cancel with Ctrl+C without corrupting your files

## 🔧 Installation

### ⚡ Direct installation (recommended)

```bash
go install github.com/alucod3/gorrent/cmd/gorrent@latest
```

### Prerequisites

- Go 1.16 or higher

### Installation from source code

```bash
# Clone the repository
git clone https://github.com/alucod3/gorrent.git
cd gorrent

# Install dependencies
go mod tidy

# Build the project
go build -o gorrent cmd/gorrent/main.go
```

## 🚀 Usage

```bash
# Run the program
gorrent

# When prompted, you can:
# - Paste a magnet link
# - Drag a .torrent file to the terminal
# - Type the path to a local .torrent file
```

### Usage examples

```bash
# Start the program and follow the interactive instructions
gorrent

# Start download directly with a magnet link
gorrent "magnet:?xt=urn:btih:..."

# Start download with a local .torrent file
gorrent ~/Downloads/ubuntu-22.04.torrent
```

## 🏗️ Project Structure

The project follows a modular structure according to Go best practices:

```
gorrent/
├── cmd/
│   └── gorrent/      # Application entry point
├── internal/         # Private application-specific packages
│   ├── cli/          # Command-line interface
│   ├── config/       # Application configurations
│   ├── downloader/   # Torrent download logic
│   └── validator/    # Link and file validation
└── pkg/              # Public reusable packages
    └── utils/        # Various utilities
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork this repository
2. Create your feature branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -m 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Open a Pull Request

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Contact

For questions, suggestions, or feedback, please open an [issue](https://github.com/alucod3/gorrent/issues).

---

Made with ❤️ in Go