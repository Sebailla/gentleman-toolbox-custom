package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gentleman-programming/gentle-ai/internal/system"
)

// RunPlan genera el andamiaje (scaffolding) para un nuevo módulo siguiendo el patrón Modular Vertical Slicing.
func RunPlan(args []string, detection system.DetectionResult) error {
	var featureName string
	for i, arg := range args {
		if strings.HasPrefix(arg, "--feature=") {
			featureName = strings.TrimPrefix(arg, "--feature=")
		} else if arg == "--feature" && i+1 < len(args) {
			featureName = args[i+1]
		}
	}

	if featureName == "" {
		fmt.Println("❌ Error: Falta el nombre de la feature.")
		fmt.Println("👉 Uso: gentle-ai plan --feature=<nombre>")
		return nil
	}

	featureName = strings.ToLower(featureName)
	fmt.Printf("🏗️  Gentleman Plan - Scaffolding de Módulo: %s\n", featureName)
	fmt.Println("====================================================")
	
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	moduleDir := filepath.Join(cwd, "src", "modules", featureName)
	if _, err := os.Stat(moduleDir); err == nil {
		fmt.Printf("⚠️  El módulo '%s' ya existe. Omitiendo generación.\n", featureName)
		return nil
	}

	// 1. Crear directorios
	dirs := []string{
		filepath.Join(moduleDir, "services"),
		filepath.Join(moduleDir, "components"),
	}

	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return fmt.Errorf("error al crear directorio %s: %w", d, err)
		}
	}

	// 2. Crear archivos base
	files := map[string]string{
		"actions.ts": "'use server';\n\n/**\n * " + featureName + " actions\n * Validar con Zod y delegar lógica a servicios.\n */\n",
		"types.ts":   "/**\n * " + featureName + " domain types\n */\n\nexport interface " + strings.Title(featureName) + " {\n  id: string;\n}\n",
		"index.ts":   "/**\n * Public API for " + featureName + " module\n */\n\nexport * from './types';\n",
	}

	for name, content := range files {
		path := filepath.Join(moduleDir, name)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("error al crear archivo %s: %w", path, err)
		}
	}

	fmt.Printf("✅ Módulo '%s' creado exitosamente.\n", featureName)
	fmt.Println("📂 Estructura: actions.ts, types.ts, index.ts, services/, components/")
	
	return nil
}
