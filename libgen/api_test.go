package libgen

import "testing"

func TestGetDetails(t *testing.T) {
	books, err := GetDetails([]string{"2F2DBA2A621B693BB95601C16ED680F", "06E6135019C8F2F43158ABA9ABDC610E"},
		false,
		"")
	if err != nil {
		t.Error(err)
	}
	if books[0].Title != "The Turing Test and the Frame Problem: AI's Mistaken Understanding of Intelligence" {
		t.Error()
	}
	if books[1].Title != "You failed your math test, Comrade Einstein (about Soviet antisemitism)" {
		t.Error()
	}
}
