package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gentleman-programming/gentle-ai/internal/system"
)

// RunDoctor diagnostica el estado del búnker y, opcionalmente, lo repara.
func RunDoctor(args []string, detection system.DetectionResult) error {
	fmt.Println("🏥 Gentleman Doctor - Diagnosis del Búnker")
	fmt.Println("==========================================")

	fixMode := false
	for _, arg := range args {
		if arg == "--fix" {
			fixMode = true
		}
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get current directory: %w", err)
	}

	// 1. Chequeo de Dependencias
	fmt.Println("\n📦 Dependencias del Sistema:")
	deps := []string{"bun", "git", "go", "pip3"}
	for _, d := range deps {
		if _, err := exec.LookPath(d); err == nil {
			fmt.Printf("  ✅ %-10s: Instalado\n", d)
		} else {
			fmt.Printf("  ❌ %-10s: NO ENCONTRADO\n", d)
		}
	}

	// 2. Chequeo de Engram
	fmt.Println("\n🧠 Memoria (Engram):")
	if _, err := exec.LookPath("engram"); err == nil {
		fmt.Println("  ✅ Engram: Binario encontrado en el PATH")
	} else {
		fmt.Println("  ❌ Engram: Binario no encontrado. Los agentes no tendrán memoria persistente.")
	}

	// 3. Chequeo de Graphify
	fmt.Println("\n🗺️  Knowledge Graph (Graphify):")
	if _, err := os.Stat(filepath.Join(cwd, "graphify-out")); err == nil {
		fmt.Println("  ✅ Reporte: Encontrado en graphify-out/")
	} else {
		fmt.Println("  ⚠️  Reporte: No encontrado. Ejecutá 'graphify .' para generar el mapa del proyecto.")
	}

	// 4. Chequeo de Sincronización de Reglas
	fmt.Println("\n⚖️  Sincronización de Reglas:")
	masterRule := filepath.Join(cwd, "AGENTS.md")
	if _, err := os.Stat(masterRule); err == nil {
		fmt.Println("  ✅ Maestro: AGENTS.md presente")
		
		// Verificar réplicas
		replicas := []string{"CLAUDE.md", ".cursorrules", "GEMINI.md"}
		for _, r := range replicas {
			replicaPath := filepath.Join(cwd, r)
			if _, err := os.Stat(replicaPath); err == nil {
				fmt.Printf("  ✅ Réplica: %s presente\n", r)
			} else {
				if fixMode {
					fmt.Printf("  🔧 Reparando: Creando réplica %s desde el maestro...\n", r)
					if err := copyFile(masterRule, replicaPath); err != nil {
						fmt.Printf("  ❌ Error al reparar %s: %v\n", r, err)
					} else {
						fmt.Printf("  ✅ Réplica: %s REPARADA\n", r)
					}
				} else {
					fmt.Printf("  ❌ Réplica: %s FALTANTE (Corré con --fix para reparar)\n", r)
				}
			}
		}
	} else {
		fmt.Println("  ❌ Maestro: AGENTS.md NO ENCONTRADO. El proyecto no sigue el estándar Gentleman.")
	}

	if fixMode {
		fmt.Println("\n✨ Reparación completada.")
	} else {
		fmt.Println("\n✨ Diagnóstico terminado. ¡Mantené tu búnker en orden!")
	}
	return nil
}

// Función auxiliar para copiar archivos
func copyFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, input, 0644)
}
