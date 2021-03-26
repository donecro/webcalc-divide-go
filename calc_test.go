package main

import (
	"testing"
)


func TestCalc(t *testing.T) {
	answer := DoCalc("6", "2")
	if answer != "3" {
		t.Error("Calc failed")
	}
}