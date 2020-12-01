package util

import "testing"

func Test_fileExists(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{"empty", "", false},
		{"dir", "testdata", false},
		{"exists", "util.go", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExists(tt.filename); got != tt.want {
				t.Errorf("fileExists() = %v, want %v", got, tt.want)

			}

		})

	}

}
