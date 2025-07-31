package config

import (
	"time"

	"gorm.io/gorm/logger"
)

// app runtime.
const (
	defaultAppPort            = 8000
	defaultAppPrometheus      = false
	defaultAppShutdownTimeout = 200 * time.Millisecond
	defaultAppShutdownWait    = 200 * time.Millisecond
)

// database connection.
const (
	defaultDBMaxIdle           = 20
	defaultDBMaxOpen           = 100
	defaultDBMaxLifetime       = 10 * time.Second
	defaultDBKeepAliveInterval = 3 * time.Minute

	// gorm option.
	defaultGormLog              = true
	defaultGormLogSlowThreshold = 200 * time.Millisecond
	defaultGormLogLevel         = logger.Warn
	defaultGormLogIgnoreErr     = false
	defaultGormLogColorful      = true
)
