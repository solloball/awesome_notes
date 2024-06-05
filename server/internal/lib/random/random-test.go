package random

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "size = 1",
			size: 1,
		},
		{
			name: "size = 5",
			size: 5,
		},
		{
			name: "size = 10",
			size: 10,
		},
		{
			name: "size = 20",
			size: 20,
		},
		{
			name: "size = 30",
			size: 30,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str1 := NewRandomString(tt.size, int64(i + 1))
			str2 := NewRandomString(tt.size, int64(i))

			assert.Len(t, str1, tt.size)
			assert.Len(t, str2, tt.size)

            // It is not a guarantee
			assert.NotEqual(t, str1, str2)
		})
	}
}
