package main

import "testing"

func TestGetRobot(t *testing.T) {
	if got := getRobot("unknown"); got != nil {
		t.Errorf("getRobot(\"unknown\") = %v, want %v", got, nil)
	}
}
