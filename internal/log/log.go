package log

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/kyoh86/xdg"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *log.Logger

func init() {
	logDir := filepath.Join(xdg.DataHome(), "render-alt-delete", "logs")

	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Error("failed to create log directory", "error", err)
		// Fall back to current directory
		logDir = "."
	}

	logPath := filepath.Join(logDir, "app.log")

	// Set up lumberjack for log rotation
	logWriter := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1, // megabytes
		MaxBackups: 1,
	}

	// Configure charmbracelet/log to use lumberjack
	log.SetOutput(logWriter)

	Logger = log.Default()

}
