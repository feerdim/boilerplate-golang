package config

import "time"

func SetDefaultTimezone() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		loc = time.Now().Location()
	}

	time.Local = loc
}
