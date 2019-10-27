package hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	result := SayHello("test")
	want := "Hello: test"
	if result != want {
		t.Fatalf("SayHello(%q) = %s, Expected %s", "test", result, want)
	}
}
