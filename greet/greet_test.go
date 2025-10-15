package greet

import "testing"

func assertString(t testing.TB, got, want string){
	t.Helper()
	if got != want{
		t.Fatalf("mismatch: got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got error, want bool){
	t.Helper()
	if (got != nil) != want {
        if want {
            t.Fatalf("expected error, but got nil")
        } else {
            t.Fatalf("unexpected error: %v", got)
        }
    }
}

func TestHello(t *testing.T){
	t.Run("basic", func(t *testing.T){
		got, err := Hello("Go")
		assertError(t, err, false)
		assertString(t, got, "Hello, Go")

		got, err = Hello("World")
        assertError(t, err, false)
        assertString(t, got, "Hello, World")
	})

	t.Run("empty", func (t *testing.T){
		got, err := Hello("")
		assertError(t, err, true)
		assertString(t, got, "")
	})

	t.Run("unicode", func (t *testing.T){
		got, err := Hello("Гофер")
		assertError(t, err, false)
		assertString(t, got, "Hello, Гофер")
	})

	t.Run("spaces", func(t *testing.T){
		got, err := Hello(" Go ")
		assertError(t, err, false)
		assertString(t, got, "Hello,  Go ")
	})

}