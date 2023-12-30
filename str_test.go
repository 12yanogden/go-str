package str

import (
	"testing"
)

func TestIsAlphaNum(t *testing.T) {
	alphaNum := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	expected := true
	actual := IsAlphaNum(alphaNum)

	if expected != actual {
		t.Fatalf("\nExpected:\t%t, the variable is alpha numberic\nActual:\t\t%t, the variable is not alpha numberic\n", expected, actual)
	}
}

func TestIsNotAlphaNum(t *testing.T) {
	alphaNum := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789*"
	expected := false
	actual := IsAlphaNum(alphaNum)

	if expected != actual {
		t.Fatalf("\nExpected:\t%t, the variable is not alpha numberic\nActual:\t\t%t, the variable is alpha numberic\n", expected, actual)
	}
}

func TestTrimSides(t *testing.T) {
	str := ">['value']"
	expected := "value"
	actual := TrimSides(str, 3, 2)

	if expected != actual {
		t.Fatalf("\nExpected:\t%s\nActual:\t\t%s", expected, actual)
	}
}
