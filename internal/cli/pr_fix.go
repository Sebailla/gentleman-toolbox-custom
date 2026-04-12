package cli

import (
	"fmt"
	"os/exec"

	"github.com/gentleman-programming/gentle-ai/internal/system"
)

// RunPRFix inicia el proceso de corrección automática de un Pull Request basado en comentarios.
func RunPRFix(args []string, detection system.DetectionResult) error {
	if len(args) == 0 {
		fmt.Println("❌ Error: Debés proveer la URL o el número del Pull Request.")
		fmt.Println("👉 Uso: gentle-ai pr-fix <url-o-id>")
		return nil
	}

	prID := args[0]
	fmt.Printf("🛠️  Gentleman PR-Fix - Analizando Pull Request #%s\n", prID)
	fmt.Println("=========================================================")

	// Verificar si gh CLI está disponible
	if _, err := exec.LookPath("gh"); err != nil {
		fmt.Println("⚠️  No se encontró la CLI 'gh' de GitHub.")
		fmt.Println("💡 Tip: Instalá 'gh' y corré 'gh auth login' para habilitar esta función.")
		return nil
	}

	fmt.Println("\n🔍 Extrayendo comentarios del PR...")
	// Señal para que el Orquestador (Agente) tome el mando con la skill pr-fix.
	fmt.Println("\n📢 SEÑAL DE ORQUESTACIÓN: [PR_FIX_TRIGGER]")
	fmt.Printf("👉 Orquestador: Usá 'gh pr view %s --json comments' para leer el feedback y corregirlo.\n", prID)
	
	return nil
}
