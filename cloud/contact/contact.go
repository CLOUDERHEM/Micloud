package contactmgr

import "io.github.clouderhem.micloud/cloud/contact/contact"

func ListContacts(limit int) (contact.Contacts, error) {
	return contact.ListContacts(limit)
}
