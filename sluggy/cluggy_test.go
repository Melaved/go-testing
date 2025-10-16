package sluggy

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestSLug(t *testing.T){
	tests := []struct {
		name string
		in string
		want string
	}{
		{"basic", "Hello World", "hello-world"},
		{"punct", "Go, Dev!", "go-dev"},
		{"dupes", "a---b_c","a-b-c"},
		{"unicode", "Привет, Мир", "привет-мир"},
		{"empty","",""},
	}

	for _, tc := range tests{
		t.Run(tc.name, func(t *testing.T){
			got := Slug(tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}