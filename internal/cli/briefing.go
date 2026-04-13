package cli

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/gentleman-programming/gentle-ai/internal/system"
)

// RunBriefing genera un documento de traspaso (handover) para el siguiente agente o sesión.
func RunBriefing(args []string, detection system.DetectionResult) error {
	fmt.Println("🎖️  Gentleman Briefing - Generando Mission Handover")
	fmt.Println("================================================")

	_, err := os.Getwd()
	if err != nil {
		return err
	}

	// 1. Recolectar info de Git
	branch, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	status, _ := exec.Command("git", "status", "--short").Output()

	// 2. Crear el contenido
	briefing := fmt.Sprintf("# 🎖️ Mission Briefing: %s\n", time.Now().Format("2006-01-02 15:04"))
	briefing += fmt.Sprintf("**Rama Actual**: `%s`\n", string(branch))
	briefing += "\n## 📍 Estado de Git\n```text\n" + string(status) + "```\n"
	
	briefing += "\n## 🧠 Instrucciones para el Siguiente Agente\n"
	briefing += "> Este documento es un traspaso de mando. Leé la memoria de Engram (`mem_context`) para entender las decisiones recientes.\n\n"
	briefing += "- **Contexto**: [AGREGE AQUÍ EL RESUMEN DE LA TAREA ACTUAL]\n"
	briefing += "- **Pendientes**: [AGREGUE AQUÍ LOS SIGUIENTES PASOS]\n"
	briefing += "- **Riesgos**: [AGREGUE AQUÍ POSIBLES BLOQUEOS]\n"

	// 3. Escribir archivo
	briefingPath := "MISSION_BRIEFING.md"
	if err := os.WriteFile(briefingPath, []byte(briefing), 0644); err != nil {
		return fmt.Errorf("error al escribir MISSION_BRIEFING.md: %w", err)
	}

	fmt.Printf("✅ Briefing generado en %s\n", briefingPath)
	fmt.Println("👉 Ahora pedile a la IA: 'Completá el briefing usando la skill gentleman-briefing'")
	
	return nil
}
