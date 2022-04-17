package uk_test

import (
	"testing"

	"github.com/yut-kt/goholiday"
	"github.com/yut-kt/goholiday/nholidays/uk"
)

func TestNewEngland(t *testing.T) {
	goholiday.New(uk.NewEngland())
}
