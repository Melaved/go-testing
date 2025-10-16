package hashutil

import (
	"testing"
)

func TestHashSHA256_Parallel(t *testing.T){
	tests := []struct{
		name string
		in string
		want string
	}{
		{"want", "", GenerateExpectedHash("")},
		{"simple ascii", "hello", GenerateExpectedHash("hello")},
		{"unicode", "Привет", GenerateExpectedHash("Привет")},
	}


	for _, test := range tests{
		test := test
		t.Run(test.name, func(t *testing.T){
			t.Parallel()
			if got := HashSHA256(test.in); got != test.want{
				t.Fatalf("%s: got %s, want %s", test.name, got, test.want)
			}
		})
	}
}