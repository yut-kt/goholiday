package config

import "time"

const (
	// DateFormat is "yyyy-MM-dd"
	DateFormat = "2006-01-02"
)

var (
	// JST is Japan's time zone
	JST, _ = time.LoadLocation("Asia/Tokyo")
)
