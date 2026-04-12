package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gentleman-programming/gentle-ai/internal/system"
)

// RunDrive ejecuta un comando de test y formatea el resultado para que el agente de IA
// pueda procesar las fallas y aplicar correcciones automáticamente.
func RunDrive(args []string, detection system.DetectionResult) error {
	fmt.Println("🏎️  Gentleman Driver (Autonomous Mode)")

	if len(args) == 0 {
		fmt.Println("Uso: gentle-ai drive --test=\"go test ./...\"")
		return nil
	}

	var testCmd string
	for _, arg := range args {
		if strings.HasPrefix(arg, "--test=") {
			testCmd = strings.TrimPrefix(arg, "--test=")
		}
	}

	if testCmd == "" {
		// Si no se usa --test, intentamos reconstruir desde los argumentos planos
		testCmd = strings.Join(args, " ")
	}

	fmt.Printf("🚀 Ejecutando: %s\n\n", testCmd)
	
	// Dividimos el comando correctamente para exec.Command
	// Esto es simplificado; para comandos complejos con quotes se requeriría un parser de shell
	cmdParts := strings.Fields(testCmd)
	if len(cmdParts) == 0 {
		return fmt.Errorf("comando de test vacío")
	}

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("\n❌ Falla detectada: %v\n", err)
		fmt.Println("\n🤖 [GENTLEMAN ADVICE]")
		fmt.Println("1. Analizá el stack trace arriba.")
		fmt.Println("2. Identificá si es un regression bug o una falla de lógica nueva.")
		fmt.Println("3. Usá la skill 'gentleman-driver' para iterar el fix.")
		fmt.Println("4. Repetí este comando hasta que veas el verde.")
		
		// Retornamos nil para que el wrapper no aborte el proceso del agente
		// y éste pueda leer la salida estándar.
		return nil
	}

	fmt.Println("\n✅ ¡Todo verde! Los tests pasaron. Podés proceder con el siguiente paso.")
	return nil
}
