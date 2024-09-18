package excel

import (
	"os"
	"testing"
)

func TestExcel(t *testing.T) {
	excel, err := NewSingleSheetExcel("/tmp/micloud_test.xlsx")
	if err != nil {
		t.Error(err)
	}
	excel.SetHead([]string{"a", "b", "c"})

	excel.AddStrsRow([]string{"1", "2", "3"})
	excel.AddStrsRow([]string{"4", "5", "6"})

	err = excel.Save()
	if err != nil {
		t.Error(err)
	}
	os.Remove(excel.Filepath)
}
