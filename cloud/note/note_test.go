package note

import (
	"fmt"
	"testing"
)

func TestListNotes(t *testing.T) {
	notes, err := ListNotes(200)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(notes)
}

func TestGetNote(t *testing.T) {
	note, err := GetNote("43465194589668640")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(note)
}

func TestDeleteNote(t *testing.T) {
	err := DeleteNote("43465194589668640", "43465212726624768", false)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateNote(t *testing.T) {
	notes, err := ListDeletedNotes("43465212726624768", 200)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(notes)
}
