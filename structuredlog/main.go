package main

import (
	"context"
	"fmt"
	"golang.org/x/exp/slog"
	"log"
	"os"
)

func main() {
	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	log.Println("Hello from old logger")
	slog.SetDefault(logger)
	slog.Info("Info message")
	log.Println("Hello from old logger")

	logLogger := slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), slog.LevelError)
	fmt.Printf("logLogger: %T\n", logLogger)

	logger.Info(
		"incoming request",
		"method", "GET",
		"time_taken_ms", 158,
		"path", "/hello/world?q=search",
		"status", 200,
		"user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)",
	)

	logger.Info(
		"incoming request",
		slog.String("method", "GET"),
		slog.Int("time_taken_ms", 158),
		slog.String("path", "/hello/world?q=search"),
		slog.Int("status", 200),
		slog.String(
			"user_agent",
			"Googlebot/2.1 (+http://www.google.com/bot.html)",
		),
	)

	logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"incoming request",
		slog.String("method", "GET"),
		slog.Int("time_taken_ms", 158),
		slog.String("path", "/hello/world?q=search"),
		slog.Int("status", 200),
		slog.String(
			"user_agent",
			"Googlebot/2.1 (+http://www.google.com/bot.html)",
		),
	)

	logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"image uploaded",
		slog.Int("id", 23123),
		slog.Group("properties",
			slog.Int("width", 4000),
			slog.Int("height", 3000),
			slog.String("format", "jpeg"),
		),
	)

	child := logger.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", "1"),
		),
	)

	child.Info("image upload successful", slog.String("image_id", "39ud88"))
	child.Warn(
		"storage is 90% full",
		slog.String("available_space", "900.1 mb"),
	)
}
