package validate

import "testing"

func TestValidateName_Empty(t *testing.T){
	if err := ValidateName(""); err == nil{
		t.Fatalf("expected error, got nil")
	} else if err != ErrEmptyName{
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidate_NonEmpty(t *testing.T){
	if err := ValidateName("Hello"); err != nil{
		t.Fatalf("unexpected error: %v", err)
	}
}