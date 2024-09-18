package mdownload

import (
	"os"
)

func Download(url string, dir, filename string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	// todo

	return nil
}
