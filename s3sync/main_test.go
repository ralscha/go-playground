package main

import (
	"strings"
	"testing"
)

func TestRunRejectsInvalidArguments(t *testing.T) {
	for _, tt := range []struct {
		name string
		args []string
		want string
	}{
		{"no arguments", nil, "usage:"},
		{"missing job count", []string{"source", "target"}, "usage:"},
		{"extra argument", []string{"source", "target", "2", "extra"}, "usage:"},
		{"non-integer jobs", []string{"source", "target", "many"}, "positive integer"},
		{"zero jobs", []string{"source", "target", "0"}, "positive integer"},
		{"negative jobs", []string{"source", "target", "-1"}, "positive integer"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			err := run(t.Context(), tt.args)
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("run(%q) = %v, want error containing %q", tt.args, err, tt.want)
			}
		})
	}
}
