package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gentleman-programming/gentle-ai/internal/system"
)

// RunLearn prepara el entorno para que el agente de IA extraiga "instintos"
// basados en el feedback y las correcciones del usuario durante la sesión.
func RunLearn(args []string, detection system.DetectionResult) error {
	fmt.Println("🧠 Gentleman Instinct Learner")
	
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get current directory: %w", err)
	}

	// Asegurar que existe la carpeta de configuración local
	instinctsDir := filepath.Join(cwd, ".gentleman")
	if _, err := os.Stat(instinctsDir); os.IsNotExist(err) {
		err = os.Mkdir(instinctsDir, 0755)
		if err != nil {
			return fmt.Errorf("crear carpeta .gentleman: %w", err)
		}
	}

	instinctsPath := filepath.Join(instinctsDir, "instincts.md")
	
	fmt.Println("\n🔍 Analizando memoria de la sesión...")
	fmt.Println("\n🤖 [INSTRUCTION]")
	fmt.Println("1. Consultá Engram (`mem_context`) para ver las últimas interacciones.")
	fmt.Println("2. Buscá correcciones del usuario, preferencias de estilo o patrones repetidos.")
	fmt.Println("3. Escribí o actualizá los instintos en: " + instinctsPath)
	fmt.Println("4. Estos instintos serán leídos en futuras sesiones para evitar repetir errores.")
	
	if _, err := os.Stat(instinctsPath); err == nil {
		fmt.Printf("\n📝 Archivo de instintos actual detectado en: %s\n", instinctsPath)
	} else {
		fmt.Println("\n✨ No se encontró archivo de instintos previo. Se generará uno nuevo.")
	}

	return nil
}
