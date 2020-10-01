package arguments

import (
	"log"
	"os/exec"
)

func ValidateArgsRunning(args []string) bool {
	if len(args) < 2 {
		return false
	}
	return true
}

func ValidateSearchArguments(args []string) bool {
	if len(args) < 1 {
		return false
	}
	return true
}

func ValidateEditorArguments(args []string) bool {
	if len(args) < 1 {
		return false
	}

	// type コマンドだと何故か上手く行かない
	cmd := exec.Command("which", args...)
	if err := cmd.Run(); err != nil {
		log.Println(err)
		return false
	}
	exitCode := cmd.ProcessState.ExitCode()

	if exitCode != 0 {
		return false
	}
	return true
}
