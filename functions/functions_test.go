package functions

import "testing"

func TestApostrophe(t *testing.T) {
	words := "' I am a optimist ,but a optimist who carries a raincoat . ' "
	want :=  "'I am a optimist ,but a optimist who carries a raincoat .'"
	got := Apostrophe(words)
	if want != got {
		t.Fatalf("wanted: %v but got: %v ",want,got)
	}
}
func TestConvert_Binary(t *testing.T) {
	words := "Simply add and 10 (bin)"
	got := Convert_Binary(words)
	want := "Simply add and 2"
	if  got != want {
		t.Fatalf("wanted:%v but got: %v",want,got)
	}
}
func TestHexiDecimal(t *testing.T) {
	words := "1E (hex) files were added"
	 want := "30 files were added"
	 got  := HexiDecimal(words)
	 if got != want {
		t.Fatalf("wanted:%v but got: %v",want ,got)
	 }
}
func TestIsLower(t *testing.T) {
    words := "There is no greater AGONY (low) than bearing an UNTOLD STORY (low, 2) inside you."
	got := IsLower(words)
	want := "There is no greater agony than bearing an untold story inside you."
	 if got != want {
		t.Fatalf("wanted:%v got :%v",want,got)
	 }
}
func TestPunctuation(t *testing.T) {
	words := "' I am a optimist ,but a optimist who carries a raincoat .'"
	got := Punctuation(words)
	want := "' I am a optimist, but a optimist who carries a raincoat. '"
	if got != want {
		t.Fatalf("wanted: %v but got %v", want, got)
	}
}
func TestIsUpper(t *testing.T){
	words := "it (up) was the best of times, it was the worst of times.This is so exciting (up, 2)"
	want := "IT was the best of times, it was the worst of times.This is SO EXCITING"
	got := IsUpper(words)
	if got != want {
		t.Fatalf("wanted: %v but got: %v",want,got)
	}
}
func TestChangeVowel(t *testing.T) {
	words := "There it was. A amazing rock!"
	got := ChangeVowel(words)
	want := "There it was. An amazing rock!"
	if got != want {
		t.Fatalf("wanted: %v but got %v", want, got)
	}
}
func TestCapitalize(t *testing.T) {
	word := "it (cap) it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity"
	got := Capitalize(word)
	want := "It it was the age of wisdom, It Was The Age Of Foolishness , it was the epoch of belief, it was the epoch of incredulity"
	if got != want {
		t.Fatalf("wanted: %v got: %v",want,got)
	}
}