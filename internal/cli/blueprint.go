package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gentleman-programming/gentle-ai/internal/system"
)

// RunBlueprint genera un diagrama de arquitectura (Mermaid) basado en el análisis de dependencias entre módulos.
func RunBlueprint(args []string, detection system.DetectionResult) error {
	fmt.Println("🗺️  Gentleman Blueprint - Generando Mapa de Arquitectura")
	fmt.Println("=========================================================")

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	modulesDir := filepath.Join(cwd, "src", "modules")
	if _, err := os.Stat(modulesDir); os.IsNotExist(err) {
		fmt.Println("⚠️  No se encontró 'src/modules'. ¿Es este un búnker modular?")
		return nil
	}

	entries, _ := os.ReadDir(modulesDir)
	var modules []string
	for _, e := range entries {
		if e.IsDir() {
			modules = append(modules, e.Name())
		}
	}

	if len(modules) == 0 {
		fmt.Println("ℹ️  No hay módulos para diagramar.")
		return nil
	}

	// Mapa de dependencias: Modulo -> [Dependencias]
	deps := make(map[string]map[string]bool)
	importRegex := regexp.QuoteMeta("@/modules/")
	re := regexp.MustCompile(fmt.Sprintf(`from\s+'%s([^']+)'`, importRegex))

	for _, mod := range modules {
		deps[mod] = make(map[string]bool)
		modPath := filepath.Join(modulesDir, mod)
		_ = filepath.Walk(modPath, func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() && (strings.HasSuffix(path, ".ts") || strings.HasSuffix(path, ".tsx")) {
				content, _ := os.ReadFile(path)
				matches := re.FindAllStringSubmatch(string(content), -1)
				for _, match := range matches {
					target := match[1]
					if target != mod {
						deps[mod][target] = true
					}
				}
			}
			return nil
		})
	}

	// Generar Mermaid
	mermaid := "graph TD\n"
	for mod, targets := range deps {
		if len(targets) == 0 {
			mermaid += fmt.Sprintf("    %s\n", mod)
		}
		for target := range targets {
			mermaid += fmt.Sprintf("    %s --> %s\n", mod, target)
		}
	}

	fmt.Println("\n📐 Diagrama Mermaid Generado:")
	fmt.Println("```mermaid")
	fmt.Println(mermaid)
	fmt.Println("```")

	blueprintPath := "ARCHITECTURE_BLUEPRINT.md"
	content := "# 🗺️ Architectural Blueprint\n\n```mermaid\n" + mermaid + "```\n"
	_ = os.WriteFile(blueprintPath, []byte(content), 0644)

	fmt.Printf("\n✅ Mapa guardado en %s\n", blueprintPath)
	fmt.Println("💡 Tip: Copiá el código Mermaid en cualquier editor con soporte para verlo visualmente.")

	return nil
}
