package safe

import "testing"

func TestMustAt_Valid(t *testing.T){
	xs := []int{1,3,6}
	got := MustAt(xs,1)
	if got != 3{
		t.Fatalf("got %d, want %d", got, 20)
	}
}

func TestMustAt_NegativeIndex(t *testing.T){
	xs := []int{1}
	defer func(){
		if r:=recover();r ==nil{
			t.Fatalf("expected panic, got none")
		}
	}()

	_ = MustAt(xs, -1)
}

func TestMustAt_OutOfRange(t *testing.T){
	xs := []int{4,2}

	defer func(){
		if r:= recover(); r ==nil{
			 t.Fatalf("expected panic, got none")
		}
	}()

	_ = MustAt(xs, 5)
}