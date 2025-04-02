SYSTEM STRUCTURE:

gorrent/
├── cmd/
│   └── gorrent/
│       └── main.go                   # Ponto de entrada da aplicação
├── internal/
│   ├── cli/
│   │   ├── cli.go                    # Wrapper da interface CLI
│   │   ├── colors.go                 # Definições de cores para o CLI
│   │   ├── progress.go               # Gestão de barras de progresso
│   │   └── ui.go                     # Componentes de UI (logo, formatação)
│   ├── config/
│   │   └── config.go                 # Gerenciamento de configurações
│   ├── downloader/
│   │   ├── torrent.go                # Lógica de download de torrents
│   │   └── stats.go                  # Estatísticas e monitoramento
│   └── validator/
│       └── validator.go              # Validação de links e arquivos
├── pkg/
│   └── utils/
│       ├── conversions.go            # Funções de conversão (bytes, etc)
│       └── filesystem.go             # Operações de sistema de arquivos
├── go.mod
├── go.sum
├── LICENSE
└── README.md