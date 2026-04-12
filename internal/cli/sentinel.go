package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gentleman-programming/gentle-ai/internal/system"
)

// RunSentinel instala los ganchos de Git para activar el modo "Centinela".
// Soporta tanto Git nativo como integraciones con Husky.
func RunSentinel(args []string, detection system.DetectionResult) error {
	fmt.Println("🛡️  Gentleman Sentinel - Instalación de Ganchos")
	fmt.Println("============================================")
	
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get current directory: %w", err)
	}

	gitDir := filepath.Join(cwd, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return fmt.Errorf("no se detectó un repositorio Git en este directorio. Ejecutá 'git init' primero")
	}

	// 1. Determinar ruta de instalación
	huskyDir := filepath.Join(cwd, ".husky")
	var hookPath string
	useHusky := false

	if _, err := os.Stat(huskyDir); err == nil {
		fmt.Println("🐶 Husky detectado. Inyectando en .husky/pre-commit...")
		hookPath = filepath.Join(huskyDir, "pre-commit")
		useHusky = true
	} else {
		fmt.Println("⚓ Git nativo detectado. Inyectando en .git/hooks/pre-commit...")
		hookPath = filepath.Join(gitDir, "hooks", "pre-commit")
		// Asegurar que el directorio hooks existe (podría estar vacío)
		_ = os.MkdirAll(filepath.Join(gitDir, "hooks"), 0755)
	}

	// 2. Preparar contenido del hook
	sentinelHook := "\n# --- Gentle AI Sentinel ---\n"
	sentinelHook += "gentle-ai doctor || {\n"
	sentinelHook += "  echo \"❌ El Sentinel detectó problemas en el búnker.\"\n"
	sentinelHook += "  echo \"👉 Corré 'gentle-ai doctor --fix' para intentar repararlo automáticamente.\"\n"
	sentinelHook += "  exit 1\n"
	sentinelHook += "}\n"

	// 3. Leer contenido actual para evitar duplicados
	currentContent, _ := os.ReadFile(hookPath)
	if strings.Contains(string(currentContent), "Gentle AI Sentinel") {
		fmt.Println("ℹ️  El Sentinel ya está instalado en este hook. Omitiendo.")
		return nil
	}

	// 4. Escribir o Append
	f, err := os.OpenFile(hookPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo de hook (%s): %w", hookPath, err)
	}
	defer f.Close()

	// Si el archivo es nuevo y no es Husky, añadir shebang
	if len(currentContent) == 0 && !useHusky {
		if _, err := f.WriteString("#!/bin/sh\n"); err != nil {
			return err
		}
	}

	if _, err := f.WriteString(sentinelHook); err != nil {
		return fmt.Errorf("error al escribir el hook: %w", err)
	}

	// Asegurar permisos de ejecución
	_ = os.Chmod(hookPath, 0755)

	fmt.Println("✅ Sentinel instalado exitosamente.")
	return nil
}
