# GoRrent - Gerenciador de Torrents

GoRrent é um gerenciador de torrents simples, rápido e eficiente, desenvolvido em Go. Ele permite baixar arquivos de torrents a partir de magnet links ou arquivos .torrent locais.

## Características

- Interface de linha de comando amigável e intuitiva
- Suporte para magnet links e arquivos .torrent locais
- Exibição em tempo real do progresso de download
- Formatação amigável de tamanhos (KB, MB, GB)
- Feedback visual com cores e emojis
- Cancelamento seguro com Ctrl+C

## Instalação

### Pré-requisitos

- Go 1.16 ou superior

### Compilando a partir do código fonte

```bash
# Clone o repositório
git clone https://github.com/alucod3/gorrent.git
cd gorrent

# Instale as dependências
go mod tidy

# Compile o projeto
go build -o gorrent cmd/gorrent/main.go

# Ou instale diretamente
go install github.com/alucod3/gorrent/cmd/gorrent@latest
```

## Uso

```bash
# Executar o programa
./gorrent

# Quando solicitado, você pode:
# - Colar um magnet link
# - Arrastar um arquivo .torrent para o terminal
# - Digitar o caminho para um arquivo .torrent local
```

## Estrutura do Projeto

O projeto segue uma estrutura modular segundo as melhores práticas do Go:

- `cmd/gorrent`: Ponto de entrada da aplicação
- `internal/`: Pacotes privados específicos da aplicação
  - `cli/`: Interface de linha de comando
  - `config/`: Configurações da aplicação
  - `downloader/`: Lógica de download de torrents
  - `validator/`: Validação de links e arquivos
- `pkg/`: Pacotes públicos que podem ser reutilizados
  - `utils/`: Utilitários diversos

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo LICENSE para mais detalhes.