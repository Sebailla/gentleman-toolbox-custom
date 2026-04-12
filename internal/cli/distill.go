package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gentleman-programming/gentle-ai/internal/system"
)

func RunDistill(args []string, detection system.DetectionResult) error {
	fmt.Println("⚗️  Gentleman Rules Distiller")
	
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get current directory: %w", err)
	}

	// 1. Detectar archivos de reglas
	rulesFiles := []string{"AGENTS.md", "CLAUDE.md", ".cursorrules", "GEMINI.md"}
	found := []string{}
	for _, f := range rulesFiles {
		if _, err := os.Stat(filepath.Join(cwd, f)); err == nil {
			found = append(found, f)
		}
	}

	if len(found) == 0 {
		fmt.Println("⚠️  No se encontraron archivos de reglas maestros (AGENTS.md, etc.) en este directorio.")
		return nil
	}

	fmt.Printf("🔍 Detectados archivos de reglas: %v\n", found)

	// 2. Escaneo de Skills locales
	skillsDir := filepath.Join(cwd, ".agent", "skills")
	if _, err := os.Stat(skillsDir); err == nil {
		fmt.Println("📚 Detectado directorio de skills locales (.agent/skills)")
	}

	// 3. Lógica de "Distill" (Por ahora reporta el estado)
	// TODO: En el futuro, esto podría llamar a un LLM para consolidar reglas si se provee una API key,
	// o simplemente preparar un 'diff' para que el agente que lo corre lo aplique.
	
	fmt.Println("\n✨ Análisis completado.")
	fmt.Println("👉 Ahora podés pedirle a tu Agente de IA: \"Usá la skill 'gentleman-distiller' para sincronizar mis reglas con los cambios recientes\".")

	return nil
}
