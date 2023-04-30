package gotools_test

import (
	"testing"

	gotools "github.com/bastianrob/go-tools"
)

func TestIf(t *testing.T) {
	tests := []struct {
		name      string
		condition bool
		truthy    any
		falsy     any
		want      any
	}{
		{
			name:      "",
			condition: 1 == 1,
			truthy:    "one",
			falsy:     "not one",
			want:      "one",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gotools.If(tt.condition, tt.truthy, tt.falsy)
			if got == tt.want {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}
