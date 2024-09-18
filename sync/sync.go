package sync

import "io.github.clouderhem.micloud/sync/pull/note"

func PullAndSave() {
	err := notesync.PullAndSave(999)
	if err != nil {
		panic(err)
	}
}
