package main

import (
	"os"
	"testing"
)

func TestFilterOut(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		ext      string
		minSize  int64
		expected bool
	}{
		{"FilterNoExtension", "testdata/dir.log", "", 0, false},
		{"FilterExtensionMatch", "testdata/dir.log", ".log", 0, false},
		{"FilterExtensionNoMatch", "testdata/dir.log", ".sh", 0, true},
		{"FilterExtensionSizeMatch", "testdata/dir.log", ".log", 10, false},
		{"FilterExtensionSizeNoMatch", "testdata/dir.log", ".log", 20, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			info, err := os.Stat(tc.file)
			if err != nil {
				t.Error(err)
			}

			f := filterOut(tc.file, tc.ext, tc.minSize, info)

			assertBool(t, f, tc.expected)
		})
	}
}

// assertBool is a helper function that checks that the res matches exp
func assertBool(t *testing.T, res, exp bool) {
	t.Helper()

	if res != exp {
		t.Errorf("Expected '%t', got '%t' instead\n", res, exp)
	}
}
