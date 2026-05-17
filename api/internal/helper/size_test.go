package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrettyByteSize(t *testing.T) {

	t.Parallel()

	tests := []struct {
		name     string
		bytes    int64
		expected string
	}{
		{"zero", 0, "0 B"},
		{"negative", -1, "0 B"},
		{"1 byte", 1, "1 B"},
		{"999 bytes", 999, "999 B"},
		{"1023 bytes", 1023, "1023 B"},
		{"1 KB", 1024, "1 KB"},
		{"1.5 KB", 1536, "1.50 KB"},
		{"just below 1 MB", 1048575, "1024.00 KB"},
		{"exactly 1 MB", 1048576, "1 MB"},
		{"2.3 MB", 2411725, "2.30 MB"},
		{"just below 1 GB", 1073741823, "1024.00 MB"},
		{"exactly 1 GB", 1073741824, "1 GB"},
		{"just below 1 TB", 1099511627775, "1024.00 GB"},
		{"exactly 1 TB", 1099511627776, "1 TB"},
		{"15.6 TB", 17152381392962, "15.60 TB"},
		{"2.09 KB", 2140, "2.09 KB"},
		{"1.09 MB", 1142947, "1.09 MB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PrettyByteSize(tt.bytes)
			assert.Equal(t, tt.expected, result, "PrettyByteSize(%d) = %q, want %q", tt.bytes, result, tt.expected)
		})
	}
}
