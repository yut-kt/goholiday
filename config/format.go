package config

import "time"

const (
	// DateFormat is "yyyy-MM-dd"
	DateFormat = "2006-01-02"
)

var (
	// JST is Japan's time zone
	JST = time.FixedZone("Asia/Tokyo", 9*60*60)
)
