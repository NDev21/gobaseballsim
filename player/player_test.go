package player

import (
	"regexp"
	"strconv"
	"testing"
)

func TestNameString(t *testing.T) {
	name := "TestPlayer"
	testplayer := NewPlayer(name, 25, 30)
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg := NameString(testplayer)
	if !want.MatchString(msg) {
		t.Fatalf(`NameString() = %q want match for %#q`, msg, want)
	}
}

func TestNumberString(t *testing.T) {
	number := 25
	testplayer := NewPlayer("TestPlayer", number, 25)
	want := regexp.MustCompile(`\b` + strconv.Itoa(number) + `\b`)
	msg := NumberString(testplayer)
	if !want.MatchString(msg) {
		t.Fatalf(`NumberString() = %q want match %#q`, msg, want)
	}
}
func TestAgeString(t *testing.T) {
	age := 30
	testplayer := NewPlayer("TestPlayer", 25, age)
	want := regexp.MustCompile(`\b` + strconv.Itoa(age) + `\b`)
	msg := AgeString(testplayer)
	if !want.MatchString(msg) {
		t.Fatalf(`AgeString() = %q want match for %#q`, msg, want)
	}
}
