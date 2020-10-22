package server

import (
	"strings"
	"testing"
	"time"
)

type UnitTestRow struct {
	inputWords     []string // Input words
	expectedResult []bool 	// Predefined expected result
}

// Test server verify result
func TestServer(t *testing.T) {
	// Before test we need start server
	go StartHttpsServer()

	// Prepare test table
	tests := []UnitTestRow{
		{
			[]string{"golang"},
			[]bool{false},
		},
		{
			[]string{"golang"},
			[]bool{true},
		},
		{
			[]string{"golang", "python"},
			[]bool{true, false},
		},
		{
			[]string{"golang", "python", "php"},
			[]bool{true, true, false},
		},
		{
			[]string{"golang", "python", "php"},
			[]bool{true, true, true},
		},
	}

	// Start test
	for _, tt := range tests {
		actual, err := BcjClient(tt.inputWords)
		if err != nil {
			t.Errorf("BcjClient Exception detected: %s", err)
		} else {
			if !boolArrayEquals(actual, tt.expectedResult) {
				t.Errorf("BcjClient(%s); got %s; expected %s", strings.Join(tt.inputWords, ","),
					boolArrayToStr(actual), boolArrayToStr(tt.expectedResult))
			}
		}
		// each sleep 1 seconds convenient review
		time.Sleep(1e9)
	}
}

// Compare two bool array
func boolArrayEquals(a []bool, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Bool array to string
func boolArrayToStr(arr []bool) string {
	str := ""
	for _, v := range arr {
		if v {
			str += "true,"
		} else {
			str += "false,"
		}
	}
	return strings.TrimSuffix(str, ",")
}
