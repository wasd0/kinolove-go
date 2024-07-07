package utils

import (
	"os"
	"strings"
)

const (
	Separator = string(os.PathSeparator)
)

// GetUserPath returns the current user's home directory
func GetUserPath() string {
	usrPath, usrPathErr := os.UserHomeDir()

	if usrPathErr != nil {
		panic(usrPathErr)
	}

	return usrPath
}

// CreateDirectoriesIfNotExists create dirs by env path and file name
func CreateDirectoriesIfNotExists(envPath string) *strings.Builder {
	fullPath := strings.Builder{}
	paths := convertEnvPath(envPath)

	fullPath.WriteString(GetUserPath())

	for _, path := range paths {
		fullPath.WriteString(Separator)
		fullPath.WriteString(path)
		createFileIfNotExists(fullPath.String())
	}

	return &fullPath
}

// ConvertEnvPath converts a path from the environment into a suitable path for the OS.
// Env path separator must be "/"
func convertEnvPath(envPath string) []string {
	return strings.Split(envPath, "/")
}

func createFileIfNotExists(path string) {
	if info, _ := os.Stat(path); info != nil {
		return
	}

	err := os.Mkdir(path, 0755)
	if err != nil {
		panic(err)
	}
}
