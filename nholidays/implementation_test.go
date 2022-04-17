package nholidays_test

import (
	"testing"

	"github.com/yut-kt/goholiday"
	"github.com/yut-kt/goholiday/nholidays/jp"
	"github.com/yut-kt/goholiday/nholidays/uk"
)

func TestImplementation(t *testing.T) {
	// jp
	goholiday.New(jp.New())

	// uk
	goholiday.New(uk.NewEngland())
	goholiday.New(uk.NewScotland())
	goholiday.New(uk.NewNorthernIreland())
	goholiday.New(uk.NewWales())

}
