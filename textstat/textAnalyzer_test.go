package textstat

import "testing"

func TestWordCount(t *testing.T){
	tests := []struct{
		name string
		input string
		want map[string]int
	}{
		{"empty","", map[string]int{}},
		{"only separators - spaces", "    ", map[string]int{}},
		{"only separators - mixed", "  !,.(){}[]!;: \t\n\r", map[string]int{}},
		{"single word", "hello", map[string]int{"hello": 1}},
		{"multiple same words same case", "hello hello hello", map[string]int{"hello": 3}},
		{"multiple same words different case", "Hello HELLO hello heLLO", map[string]int{"hello":4}},
		{"different words", "hello world apple", map[string]int{"hello": 1, "world":1 , "apple":1}},
		{"empty tokens after processing", "word1,,,word2",map[string]int{"word1": 1, "word2": 1}},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			got := WordCount(test.input)
			if len(got) != len(test.want){
				t.Fatalf("len: got %d, want %d", len(got), len(test.want))
			}

			for k, v := range test.want{
				if got[k] != v{
					t.Fatalf("%s: got %d, want %d", k, got[k], v)
				}
			}

			for k := range got {
            if _, exists := test.want[k]; !exists {
                t.Fatalf("unexpected key in result: %q", k)
            }
        }
		})
	}
}