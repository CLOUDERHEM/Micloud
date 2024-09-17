package notemgr

import (
	"io.github.clouderhem.micloud/cloud/note/note"
	"io.github.clouderhem.micloud/cloud/note/recyclebin"
	"io.github.clouderhem.micloud/consts"
	"io.github.clouderhem.micloud/utility/parallel"
	"math/rand"
	"time"
)

func ListNotes(limit int) (note.Notes, error) {
	return note.ListNotes(limit)
}

func GetNote(id string) (note.Note, error) {
	return note.GetNote(id)
}

func ListFullNotes(noteIds []string) ([]note.Note, []parallel.ErrOut[string]) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res, errs := parallel.DoParallel[string, note.Note](noteIds,
		func(id string) (note.Note, error) {
			time.Sleep(time.Second *
				time.Duration((r.Intn(len(noteIds)/consts.DefaultReqNumInSec))+1))
			fullNote, err := note.GetNote(id)
			return fullNote, err
		})
	return res, errs
}

func DeleteNote(id, tag string, purge bool) error {
	return note.DeleteNote(id, tag, purge)
}

func ListDeletedNotes(syncTag string, limit int) (note.Notes, error) {
	return recyclebin.ListDeletedNotes(syncTag, limit)
}

// GetNoteFileUrl get file url in note, pic or record
func GetNoteFileUrl(fileId string) (string, error) {
	return note.GetNoteFileUrl(note.FileType, fileId)
}
