package utils

import (
	"os"
	"path"

	"go.uber.org/zap"
)

var (
	workingDirectory string
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		zap.S().Fatalw("os.Getwd error", "error", err)
	}

	workingDirectory = dir
}

// AbsPath returns the absolute path relative to the working directory.
func AbsPath(relativePath string) string {
	if path.IsAbs(relativePath) {
		return relativePath
	}
	return path.Join(workingDirectory, relativePath)
}
