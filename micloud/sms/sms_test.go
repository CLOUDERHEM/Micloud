package smsmgr

import (
	"fmt"
	"testing"
)

func TestListMessage(t *testing.T) {
	m1, err := ListMessages("0", "0", 20)
	if err != nil {
		t.Error(err)
	}
	if len(m1.Entries) == 0 {
		t.Error("no messages found")
	}
	fmt.Println(m1)

	if len(m1.Entries) == 20 {
		m2, err := ListMessages(m1.Watermark.SyncTag, m1.Watermark.SyncThreadTag, 20)
		if err != nil {
			t.Error(err)
		}
		if len(m2.Entries) == 0 {
			t.Error("no messages found")
		}
		fmt.Println(m2)
	}

}

func TestListDeletedMessages(t *testing.T) {
	m1, err := ListDeletedMessages("0", "0", 20)
	if err != nil {
		t.Error(err)
	}
	if len(m1.Entries) == 0 {
		t.Error("no deleted messages found")
	}
	fmt.Println(m1)

	if len(m1.Entries) == 20 {
		m2, err := ListDeletedMessages(m1.Watermark.SyncTag, m1.Watermark.SyncThreadTag, 20)
		if err != nil {
			t.Error(err)
		}
		if len(m2.Entries) == 0 {
			t.Error("no deleted messages found")
		}
		fmt.Println(m2)
	}
}
