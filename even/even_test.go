package even

import "testing"

func TestIsEven(t *testing.T) {
	if !IsEven(2){
		t.Error("Expected true for 2, but receive false")
	}

	if IsEven(3){
		t.Error("Expected false for 3, but receive true")
	}

	if !IsEven(0){
		t.Error("Expected true for 0, but receive false")
	}

	if !IsEven(-4){
		t.Error("Expected true for -4, but receive false")
	}

	if IsEven(-7){
		t.Error("Expected false for -7, but receive true")
	}
}