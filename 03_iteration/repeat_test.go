package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestReplace(t *testing.T) {
	replaced := Replace("Luis", "i", "u")
	expected := "Luus"

	if replaced != expected {
		t.Errorf("expected %q but got %q", expected, replaced)
	}
}

func TestJoin(t *testing.T) {
	t.Run("a text contains another", func(t *testing.T) {
		contains := Contains("LuisDev", "Dev")
		expected := true

		if contains != expected {
			t.Errorf("expected %t but received %t", expected, contains)
		}
	})
}

func BenchmarkReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Replace("Dev", "e", "i")
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func ExampleReplace() {
	replaced := Replace("Dev", "e", "i")
	fmt.Println(replaced)
	// Output: Div
}

func ExampleContains() {
	contains := Contains("LuisDev", "Dev")
	fmt.Println(contains)
	// Output: true
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
