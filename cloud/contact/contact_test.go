package contact

import (
	"fmt"
	"testing"
)

func TestListContacts(t *testing.T) {
	contacts, err := ListContacts(200)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(contacts)
}
