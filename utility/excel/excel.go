package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"time"
)

const SheetName = "Sheet1"

type Excel struct {
	Filepath string
	ColNames []string
	Rows     [][]any

	file *excelize.File
}

func NewSingleSheetExcel(filePath string) (*Excel, error) {
	f := excelize.NewFile()
	index, err := f.NewSheet(SheetName)
	if err != nil {
		return nil, err
	}
	f.SetActiveSheet(index)

	return &Excel{
		ColNames: nil,
		file:     f,
		Filepath: filePath,
	}, nil
}

// SetHead do set when first call
func (e *Excel) SetHead(colNames []string) {
	if len(e.ColNames) == 0 && len(colNames) != 0 {
		e.ColNames = colNames
	}
}

func (e *Excel) AddStrsRow(row []string) {
	e.AddCellsRow(convStrsToCells(row))
}

func (e *Excel) HeadInsertCellsRow(row []excelize.Cell) {
	rows := [][]any{convCellsToAnys(row)}
	e.Rows = append(rows, e.Rows...)
}

func (e *Excel) AddCellsRow(row []excelize.Cell) {
	e.Rows = append(e.Rows, convCellsToAnys(row))
}

func (e *Excel) setRows() error {
	e.HeadInsertCellsRow(convStrsToCells(e.ColNames))

	sw, err := e.file.NewStreamWriter(SheetName)
	if err != nil {
		return err
	}

	for i := 0; i < len(e.Rows); i++ {
		line := e.Rows[i]
		cell, err := excelize.CoordinatesToCellName(1, i+1)
		if err != nil {
			return err
		}
		if err := sw.SetRow(cell, line); err != nil {
			return err
		}
	}

	err = sw.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (e *Excel) Save() error {
	defer func(file *excelize.File) {
		_ = file.Close()
	}(e.file)
	err := e.setRows()
	if err != nil {
		return err
	}
	err = e.file.SaveAs(e.Filepath)
	if err != nil {
		log.Print("cannot save to Excel, try another Filepath: err: ", err)
		err := e.file.SaveAs(fmt.Sprintf("%v.tmp.%v", e.Filepath, time.Now().UnixMilli()))
		if err != nil {
			return err
		}
	}
	return nil
}

func convStrsToCells(row []string) []excelize.Cell {
	newRow := make([]excelize.Cell, len(row))
	for i := range row {
		newRow[i] = excelize.Cell{
			Value: row[i],
		}
	}
	return newRow
}

func convCellsToAnys(row []excelize.Cell) []any {
	newRow := make([]any, len(row))
	for i := range row {
		newRow[i] = row[i]
	}
	return newRow
}
