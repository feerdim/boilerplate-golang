package config

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iancoleman/strcase"
)

func NewRuntimeContext() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	return
}

type Runtime struct {
	Name                    string        `json:"name"`
	Host                    string        `json:"host"`
	Port                    int           `json:"port"`
	Prometheus              bool          `json:"prometheus"`
	ShutdownTimeoutDuration time.Duration `json:"shutdown_timeout_duration"`
	ShutdownWaitDuration    time.Duration `json:"shutdown_wait_duration"`
}

func NewRuntime() *Runtime {
	return &Runtime{
		Name:                    strcase.ToSnake(os.Getenv("APP_NAME")),
		Host:                    os.Getenv("APP_HOST"),
		Port:                    parseInt(defaultAppPort, os.Getenv("APP_PORT")),
		Prometheus:              parseBool(defaultAppPrometheus, os.Getenv("APP_PROMETHEUS")),
		ShutdownTimeoutDuration: defaultAppShutdownTimeout,
		ShutdownWaitDuration:    defaultAppShutdownWait,
	}
}
