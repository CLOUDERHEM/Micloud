package contactmgr

import "github.com/clouderhem/micloud/micloud/contact/contact"

func ListContacts(limit int) (contact.Contacts, error) {
	return contact.ListContacts(limit)
}
