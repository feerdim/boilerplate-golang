package constant

import "time"

// middleware.
const (
	DefaultMdwHeaderToken      = "Authorization"
	DefaultMdwHeaderBearer     = "Bearer"
	DefaultMdwRateLimiter      = 20
	DefaultMdwSentryDebug      = true
	DefaultMdwSentrySampleRate = 1.0
	DefaultMdwTimeout          = 10 * time.Second
)

// pagination.
const (
	DefaultOrder = "created_at DESC"
	DefaultPage  = 1
	DefaultLimit = 10
)

const DefaultCacheExpireDuration = 24 * time.Hour
