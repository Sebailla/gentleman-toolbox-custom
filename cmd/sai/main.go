package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/gentleman-programming/gentle-ai/internal/cli"
	"github.com/gentleman-programming/gentle-ai/internal/system"
)

var customCommands = map[string]bool{
	"sentinel":  true,
	"doctor":    true,
	"drive":     true,
	"learn":     true,
	"status":    true,
	"briefing":  true,
	"plan":      true,
	"judge":     true,
	"blueprint": true,
	"refactor":  true,
	"pr-fix":    true,
	"console":   true,
	"distill":   true,
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	args := os.Args[1:]
	if len(args) == 0 {
		// If no args, delegate to official TUI
		return delegate()
	}

	cmd := args[0]
	if customCommands[cmd] {
		return runCustom(args)
	}

	return delegate()
}

func runCustom(args []string) error {
	result, err := system.Detect(context.Background())
	if err != nil {
		return fmt.Errorf("detect system: %w", err)
	}

	switch args[0] {
	case "sentinel":
		return cli.RunSentinel(args[1:], result)
	case "doctor":
		return cli.RunDoctor(args[1:], result)
	case "drive":
		return cli.RunDrive(args[1:], result)
	case "learn":
		return cli.RunLearn(args[1:], result)
	case "status":
		return cli.RunStatus(args[1:], result)
	case "briefing":
		return cli.RunBriefing(args[1:], result)
	case "plan":
		return cli.RunPlan(args[1:], result)
	case "judge":
		return cli.RunJudge(args[1:], result)
	case "blueprint":
		return cli.RunBlueprint(args[1:], result)
	case "refactor":
		return cli.RunRefactor(args[1:], result)
	case "pr-fix":
		return cli.RunPRFix(args[1:], result)
	case "console":
		return cli.RunConsole(args[1:], result)
	case "distill":
		return cli.RunDistill(args[1:], result)
	default:
		return delegate()
	}
}

func delegate() error {
	path, err := exec.LookPath("gentle-ai")
	if err != nil {
		return fmt.Errorf("gentle-ai not found in PATH. Please install it via brew first")
	}

	// Prepare arguments for gentle-ai
	// We replace argv[0] with the path to gentle-ai
	newArgs := make([]string, len(os.Args))
	copy(newArgs, os.Args)
	newArgs[0] = "gentle-ai"

	return syscall.Exec(path, newArgs, os.Environ())
}
