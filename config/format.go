package config

import "time"

const (
	DateFormat = "2006-01-02"
)

var (
	JST = time.FixedZone("Asia/Tokyo", 9*60*60)
)
