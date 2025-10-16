package tempfiledemo

import (
	"errors"
	"os"
	"testing"
)

func TestWriteLinesToTemp(t *testing.T){
	t.Run("single line", func (t *testing.T){
		path, err := WriteLinesToTemp("log-", []string{"Hello"})
		if err != nil{
			t.Fatalf("unexpected error: %v", err)
		}

		defer os.Remove(path)

		data, err := os.ReadFile(path)
		if err != nil{
			t.Fatalf("read: %v", err)
		}

		if string(data) != "Hello"{
			t.Fatalf("got %q, want %q", data, "Hello")
		}
	})

	t.Run("multi line and empty", func(t *testing.T){
	cases := []struct{
		lines []string
		want string
	}{
		{lines: []string{"a", "b", "c"}, want: "a\nb\nc"},
		{lines: nil, want: ""},
	}

	for _, c := range cases{
		path, err := WriteLinesToTemp("log-", c.lines)
		if err != nil{
			t.Fatalf("unexpected error: %v", err)
		}

		data, err := os.ReadFile(path)
		if err != nil{
			t.Fatalf("read: %v", err)
		}

		if string(data) != c.want{
			t.Fatalf("got %q, want %q", data, c.want)
		}

		if err := os.Remove(path); err != nil{
			t.Fatalf("remove: %v", err)
		}

		if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist){
			t.Fatalf("expected not exists, got: %v", err)
		}
	}
	})
}