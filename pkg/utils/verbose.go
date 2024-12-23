package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/aeswibon/dep/pkg/types"
)

// GetSystemInfo collects system information.
func GetSystemInfo() types.SystemInfo {
	return types.SystemInfo{
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
		NumCPU:       runtime.NumCPU(),
		GoVersion:    runtime.Version(),
	}
}

// PrintSystemInfo prints system information if verbose is enabled.
func PrintSystemInfo(info types.SystemInfo, verbose bool) {
	if verbose {
		fmt.Printf("\nSystem Info:\n")
		fmt.Printf("  OS: %s\n", info.OS)
		fmt.Printf("  Architecture: %s\n", info.Architecture)
		fmt.Printf("  Number of CPUs: %d\n", info.NumCPU)
		fmt.Printf("  Go Version: %s\n", info.GoVersion)
	}
}

// ExecuteCommand executes the given command with the provided arguments and optionally prints execution time and system info.
func ExecuteCommand(command string, args []string, config *types.Config, flags *types.GlobalFlags) error {
	start := time.Now()
	cmd := exec.Command(command, args...)

	// Create pipes to capture stdout and stderr live
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %v", err)
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to get stderr pipe: %v", err)
	}

	// Start the command execution
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %v", err)
	}

	// Create channels to handle the output in real-time
	go io.Copy(os.Stdout, stdoutPipe)
	go io.Copy(os.Stderr, stderrPipe)

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("command execution failed: %v", err)
	}

	elapsed := time.Since(start)
	if flags.Verbose {
		fmt.Printf("\nExecution Time: %s\n", elapsed)
		systemInfo := GetSystemInfo()
		PrintSystemInfo(systemInfo, flags.Verbose)
	}
	return nil
}
