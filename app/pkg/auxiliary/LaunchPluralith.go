package auxiliary

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"pluralith/pkg/ux"
)

// Function to run OS specific launch command
func runOsCommand(command []string) error {
	functionName := "runOsCommand"

	// Instantiating new custom spinner
	spinner := ux.NewSpinner("Launching Pluralith...", "Pluralith Running\n", "Failed to launch Pluralith\n")
	spinner.Start()

	// Creating command to launch Pluralith on given OS
	cmd := exec.Command(command[0], command[1:]...)

	// Handling success and failure cases for terminal command
	// Adding slight delay to debounce for UI to get there
	if err := cmd.Run(); err != nil {
		time.Sleep(200 * time.Millisecond)
		spinner.Fail()
		return fmt.Errorf("%v: %w", functionName, err)
	} else {
		time.Sleep(200 * time.Millisecond)
		spinner.Success()
	}
}

// Function to launch Pluralith UI
func LaunchPluralith() error {
	functionName := "LaunchPluralith"

	// Running terminal command to launch application based on current OS
	switch os := runtime.GOOS; os {
	case "windows":
		if runErr := runOsCommand([]string{filepath.Join(PathInstance.HomePath, "AppData", "Local", "Programs", "pluralith", "Pluralith.exe")}); runErr != nil {
			return fmt.Errorf("could not run launch command -> %v: %w", functionName, runErr)
		}
	case "darwin":
		if runErr := runOsCommand([]string{"open", "-a", "Pluralith"}); runErr != nil {
			return fmt.Errorf("could not run launch command -> %v: %w", functionName, runErr)
		}
	default:
		if runErr := runOsCommand([]string{"command", "and", "arguments"}); runErr != nil {
			return fmt.Errorf("could not run launch command -> %v: %w", functionName, runErr)
		}
	}

	return nil
}
