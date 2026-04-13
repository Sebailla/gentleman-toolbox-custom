package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gentleman-programming/gentle-ai/internal/system"
)

// RunStatus muestra el estado del proyecto y calcula el Architecture Health Index (AHI).
func RunStatus(args []string, detection system.DetectionResult) error {
	fmt.Println("📊 Gentleman Status - Architecture Health Index")
	fmt.Println("===============================================")

	scoreMode := false
	for _, arg := range args {
		if arg == "--score" {
			scoreMode = true
		}
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	modulesDir := filepath.Join(cwd, "src", "modules")
	if _, err := os.Stat(modulesDir); os.IsNotExist(err) {
		fmt.Println("⚠️  No se encontró la carpeta 'src/modules'. ¿Es este un búnker Gentleman?")
		return nil
	}

	entries, err := os.ReadDir(modulesDir)
	if err != nil {
		return err
	}

	var totalModules int
	var perfectModules int
	var issues []string

	requiredFiles := []string{"actions.ts", "types.ts", "index.ts", "services", "components"}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		totalModules++
		moduleName := entry.Name()
		missing := []string{}

		for _, req := range requiredFiles {
			path := filepath.Join(modulesDir, moduleName, req)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				missing = append(missing, req)
			}
		}

		if len(missing) == 0 {
			perfectModules++
		} else if scoreMode {
			issues = append(issues, fmt.Sprintf("  ❌ Módulo [%s]: Faltan %v\n", moduleName, missing))
		}
	}

	if totalModules == 0 {
		fmt.Println("ℹ️  No hay módulos definidos en src/modules.")
		return nil
	}

	ahi := (float64(perfectModules) / float64(totalModules)) * 100

	fmt.Printf("\n📈 Architecture Health Index (AHI): %.1f%%\n", ahi)
	fmt.Printf("   Módulos perfectos: %d/%d\n", perfectModules, totalModules)

	if scoreMode && len(issues) > 0 {
		fmt.Println("\n🔍 Hallazgos de Deuda Técnica:")
		for _, issue := range issues {
			fmt.Print(issue)
		}
		fmt.Println("\n💡 Consejo: Cada módulo debe tener su propio silo de componentes y servicios.")
	} else if ahi < 100 {
		fmt.Println("\n💡 Usá 'gentle-ai status --score' para ver el detalle de la deuda técnica.")
	}

	return nil
}
