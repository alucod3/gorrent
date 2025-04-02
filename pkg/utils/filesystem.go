package utils

import (
	"os"
	"path/filepath"
)

// EnsureDirectoryExists garante que um diretório existe, criando-o se necessário
func EnsureDirectoryExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

// GetAbsolutePath retorna o caminho absoluto para um arquivo ou diretório
func GetAbsolutePath(path string) (string, error) {
	return filepath.Abs(path)
}

// FileExists verifica se um arquivo existe e é acessível
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirectoryExists verifica se um diretório existe e é acessível
func DirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
