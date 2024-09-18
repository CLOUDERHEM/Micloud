package notesync

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"time"

	notemgr "io.github.clouderhem.micloud/cloud/note"
	"io.github.clouderhem.micloud/cloud/note/note"
	"io.github.clouderhem.micloud/consts"
	mdownload "io.github.clouderhem.micloud/utility/download"
	"io.github.clouderhem.micloud/utility/excel"
	mjson "io.github.clouderhem.micloud/utility/json"
	"io.github.clouderhem.micloud/utility/parallel"
)

func PullAndSave(limit int) error {
	notes, err := notemgr.ListNotes(limit)
	if err != nil {
		return err
	}
	log.LogI("pulled notes, size: ", len(notes.Entries))

	var noteIds []string
	for i := range notes.Entries {
		noteIds = append(noteIds, notes.Entries[i].Id)
	}

	var fullNotes []note.Note
	for i := 0; i < 5 && len(noteIds) > 0; i++ {
		ns, errs := notemgr.ListFullNotes(noteIds)
		log.LogI("list full notes, req size: ", len(noteIds), " err size: ", len(errs))

		fullNotes = append(fullNotes, ns...)

		noteIds = nil
		if len(errs) > 0 {
			for i := range errs {
				noteIds = append(noteIds, errs[i].In)
			}
		}
	}

	if len(fullNotes) == 0 {
		return errors.New("no full notes found, abort saving")
	}
	log.LogI("find full notes, size: ", len(fullNotes))

	var files []note.File
	for i := range fullNotes {
		files = append(files, fullNotes[i].Setting.Data...)
	}

	err = savaFullNotesAsJson(fullNotes)
	if err != nil {
		log.LogE("cannot save full note as json", err)
	} else {
		log.LogI("saved notes as Json")
	}

	err = saveFullNotesAsXlsx(fullNotes)
	if err != nil {
		log.LogE("cannot save full note as xlsx", err)
	} else {
		log.LogI("saved notes as Xlsx")
	}

	err = savaFailedIdsAsJson(noteIds)
	if err != nil {
		log.LogE("cannot save failed note ids to json", err)
	} else {
		log.LogI("saved failed note ids to json, size: ", len(noteIds))
	}

	downloadNoteFiles(files)

	return nil
}

func saveFullNotesAsXlsx(notes []note.Note) error {
	xlsx, err := excel.NewSingleSheetExcel(xlsxFilepath)
	if err != nil {
		return err
	}
	for i := range notes {
		bytes, err := json.Marshal(notes[i])
		if err != nil {
			return err
		}
		keys, values := mjson.ListKeysAndValues(bytes)
		xlsx.SetHead(keys)
		xlsx.AddStrsRow(values)
	}
	return xlsx.Save()
}

func savaFullNotesAsJson(notes []note.Note) error {
	bytes, err := json.Marshal(notes)
	if err != nil {
		return err
	}
	return os.WriteFile(jsonFilepath, bytes, os.ModePerm)
}

func savaFailedIdsAsJson(ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	bytes, _ := json.Marshal(ids)
	return os.WriteFile(noteFailedFilepath, bytes, os.ModePerm)
}

func downloadNoteFiles(files []note.File) {
	if len(files) == 0 {
		log.LogI("no note files to download")
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	outs, errs := parallel.DoParallel[note.File, any](
		files,
		func(file note.File) (any, error) {
			time.Sleep(time.Second *
				time.Duration((r.Intn(len(files)/consts.DefaultReqNumInSec+1))+1))

			fileUrl, err := note.GetNoteFileUrl(note.FileType, file.FileId)
			if err != nil {
				return nil, err
			}
			err = mdownload.Download(fileUrl, filesDirName, file.FileId+getExtFromMimeType(file.MimeType))
			if err != nil {
				return nil, err
			}
			return nil, nil
		})
	log.LogI("download and save files size: ", len(outs), " err size: ", len(errs))

	err := saveFailedFilesErrs(errs)
	if err != nil {
		log.LogE("cannot save failed files", err)
	}
	log.LogI("saved failed files, size: ", len(errs))
}

func saveFailedFilesErrs(errs []parallel.ErrOut[note.File]) error {
	bytes, err := json.Marshal(errs)
	if err != nil {
		return err
	}
	return os.WriteFile(fileFailedFilepath, bytes, os.ModePerm)
}

func getExtFromMimeType(mimeType string) string {
	if mimeType == "image/jpeg" {
		return ".jpg"
	} else if mimeType == "image/png" {
		return ".png"
	} else if mimeType == "audio/mp3" {
		return ".mp3"
	}
	return ""
}
