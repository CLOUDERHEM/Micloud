package note

import (
	"io.github.clouderhem.micloud/cloud/note/note"
	"io.github.clouderhem.micloud/cloud/note/recyclebin"
)

func ListNotes(limit int) (note.Notes, error) {
	return note.ListNotes(limit)
}

func GetNote(id string) (note.Note, error) {
	return note.GetNote(id)
}

func DeleteNote(id, tag string, purge bool) error {
	return note.DeleteNote(id, tag, purge)
}

func ListDeletedNotes(syncTag string, limit int) (note.Notes, error) {
	return recyclebin.ListDeletedNotes(syncTag, limit)
}
