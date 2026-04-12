package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gentleman-programming/gentle-ai/internal/system"
)

// RunRefactor ejecuta tareas de refactorización automatizada, incluyendo el uso de instintos aprendidos.
func RunRefactor(args []string, detection system.DetectionResult) error {
	instinctMode := false
	for _, arg := range args {
		if arg == "--instincts" {
			instinctMode = true
		}
	}

	if !instinctMode {
		fmt.Println("❌ Error: Debés especificar un modo de refactorización.")
		fmt.Println("👉 Uso: gentle-ai refactor --instincts")
		return nil
	}

	fmt.Println("⚡ Gentleman Refactor - Evolucionando el Código por Instinto")
	fmt.Println("==========================================================")

	cwd, _ := os.Getwd()
	instinctsPath := filepath.Join(cwd, ".gentleman", "instincts.md")
	
	if _, err := os.Stat(instinctsPath); os.IsNotExist(err) {
		fmt.Println("⚠️  No se encontró el archivo '.gentleman/instincts.md'.")
		fmt.Println("💡 Tip: Corré 'gentle-ai learn' primero para grabar tus instintos.")
		return nil
	}

	fmt.Println("\n🧠 Cargando instintos del búnker...")
	// Señal para que el Orquestador (Agente) tome el mando con la skill instinct-refactor.
	fmt.Println("\n📢 SEÑAL DE ORQUESTACIÓN: [INSTINCT_REFACTOR_TRIGGER]")
	fmt.Println("👉 Orquestador: Aplicá los instintos guardados en .gentleman/instincts.md")
	
	return nil
}
