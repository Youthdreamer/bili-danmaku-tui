package tui

import (
	"reflect"
	"testing"
)

func TestWrapLineASCII(t *testing.T) {
	got := wrapLine("1234567", 6)
	want := []string{"123456", "7"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wrapLine() = %#v, want %#v", got, want)
	}
}

func TestWrapLineWideRunes(t *testing.T) {
	got := wrapLine("你好abc", 4)
	want := []string{"你好", "abc"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wrapLine() = %#v, want %#v", got, want)
	}
}
