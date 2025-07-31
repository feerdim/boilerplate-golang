package config

import (
	"strconv"
	"time"
)

func parseInt(i int, s string) int {
	if s == "" {
		return i
	}

	o, e := strconv.Atoi(s)
	if e != nil || o <= 0 {
		return i
	}

	return o
}

func parseDuration(d time.Duration, s string) time.Duration {
	if s == "" {
		return d
	}

	o, e := time.ParseDuration(s)
	if e != nil || o <= 0 {
		return d
	}

	return o
}

func parseBool(b bool, s string) bool {
	if s == "" {
		return b
	}

	o, e := strconv.ParseBool(s)
	if e != nil {
		return b
	}

	return o
}
