package noramalize

import "testing"

func TestClean(t *testing.T){
	tests := [] struct{
		name string
		in string
		want string
	}{
		{name: "empty", in: "", want: ""},
		{name: "spaces", in: "   hello   world  ", want: "hello world"},
		{name: "tabs", in: "\thello\t  world\t", want: "hello world"},
		{name: "case", in: "HeLlO", want: "hello"},
		{name: "mixed", in: "  A  b\tC  ", want: "a b c"},
		{name: "punct", in: " Hello,   world! ", want: "hello, world!"},
	}

	for _, test := range tests{
		got := Clean(test.in)
		if got != test.want{
			t.Errorf("%s:got %q, want %q", test.name, got, test.want)
		}
	}
}