package main

import "testing"

func TestReverse(t *testing.T) {
	test1 := reverse("the sky is blue")
	test2 := reverse("   hello world")
	test3 := reverse("a good   example")
	if test1 != "blue is sky the" || test2 != "world hello" || test3 != "example good a" {
		t.Error("Wrong Answer")
	}
}
