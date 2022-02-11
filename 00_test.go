package main

import (
	"strings"
	"testing"
)

func Test0(t *testing.T) {
	want := string([]rune{72, 101, 108, 108, 111, 32, 128506, 65039, 32, 33})
	check:=PrintS()
	if!strings.EqualFold(want,check) {
		t.Errorf("Result got was incorrect, check: %s, want: %s", want, "Hello :world_map:!")

	}
}
