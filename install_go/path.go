package install_go

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"zenbox/print"
)

const (
	bashConfig = ".bash_profile"
	zshConfig  = ".zshrc"
)

func findGo(ctx context.Context, cmd string) (string, error) {
	out, err := exec.CommandContext(ctx, cmd, "go").CombinedOutput()
	return strings.TrimSpace(string(out)), err
}

func getLocalGoVersion(ctx context.Context) string {
	out, _ := exec.CommandContext(ctx, "go", "version").CombinedOutput()
	fields := strings.Fields(strings.TrimSpace(string(out)))
	if len(fields) > 3 {
		return strings.Title(fields[2])
	}

	return "Unknown"
}

func getGOROOT(ctx context.Context) string {
	out, _ := exec.CommandContext(ctx, "go", "env", "GOROOT").CombinedOutput()
	return strings.TrimSpace(string(out))
}

func getHomeDir() (string, error) {
	home := os.Getenv(homeKey)
	if home != "" {
		return home, nil
	}

	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return u.HomeDir, nil
}

func isShell(name string) bool {
	return strings.Contains(currentShell(), name)
}

func shellConfigFile() (string, error) {
	home, err := getHomeDir()
	if err != nil {
		return "", err
	}

	switch {
	case isShell("bash"):
		return filepath.Join(home, bashConfig), nil
	case isShell("zsh"):
		return filepath.Join(home, zshConfig), nil
	default:
		return "", fmt.Errorf("目前尚不支持该终端: %q", currentShell())
	}
}

func persistEnvVarWindows(name, value string) error {
	out, err := exec.Command(
		"powershell",
		"-command",
		fmt.Sprintf(`[Environment]::SetEnvironmentVariable("%s", "%s", "User")`, name, value),
	).CombinedOutput()
	if out != nil && err == nil && len(out) != 0 {
		print.E(out)
	}
	return err
}

func checkStringExistsFile(filename, value string) (bool, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == value {
			return true, nil
		}
	}

	return false, scanner.Err()
}

func appendToFile(filename, value string) error {
	ok, err := checkStringExistsFile(filename, value)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(lineEnding + value + lineEnding)
	return err
}

func persistEnvVar(name, value string) error {
	if runtime.GOOS == "windows" {
		if err := persistEnvVarWindows(name, value); err != nil {
			return err
		}

		if isShell("cmd.exe") || isShell("powershell.exe") {
			return os.Setenv(strings.ToUpper(name), value)
		}
	}

	rc, err := shellConfigFile()
	if err != nil {
		return err
	}

	line := fmt.Sprintf("export %s=%s", strings.ToUpper(name), value)
	if err := appendToFile(rc, line); err != nil {
		return err
	}

	return os.Setenv(strings.ToUpper(name), value)
}

func appendToPATH(value string) error {
	if isInPATH(value) {
		return nil
	}
	return persistEnvVar("PATH", pathVar+envSeparator+value)
}

func isInPATH(dir string) bool {
	p := os.Getenv("PATH")

	paths := strings.Split(p, envSeparator)
	for _, d := range paths {
		if d == dir {
			return true
		}
	}

	return false
}
