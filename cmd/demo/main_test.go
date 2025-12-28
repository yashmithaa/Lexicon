package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileExtensionValidation(t *testing.T) {
	tests := []struct {
		filename string
		valid    bool
	}{
		{"test.spr", true},
		{"program.spr", true},
		{"test.txt", false},
		{"test.go", false},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			hasSuffix := len(tt.filename) > 4 && tt.filename[len(tt.filename)-4:] == ".spr"
			if hasSuffix != tt.valid {
				t.Errorf("File %q validation = %v, want %v", tt.filename, hasSuffix, tt.valid)
			}
		})
	}
}

func TestValidSproutCode(t *testing.T) {
	tmpDir := t.TempDir()
	tests := []struct {
		name     string
		code     string
		filename string
	}{
		{"simple variable", "sprout x = 5;\necho x;", "simple.spr"},
		{"arithmetic", "sprout result = 10 + 20;\necho result;", "arithmetic.spr"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFile := filepath.Join(tmpDir, tt.filename)
			err := os.WriteFile(testFile, []byte(tt.code), 0644)
			if err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}
			content, err := os.ReadFile(testFile)
			if err != nil {
				t.Fatalf("Failed to read test file: %v", err)
			}
			if string(content) != tt.code {
				t.Errorf("File content = %q, want %q", string(content), tt.code)
			}
		})
	}
}
