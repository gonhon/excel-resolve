package parse

import (
	"log"

	"github.com/xuri/excelize/v2"
)

type ProcessRow interface {
	Process([]string, string, int) error
}
type ProcessRowFunc func([]string, string, int) error

func (f ProcessRowFunc) Process(data []string, sheetName string, index int) error {
	return f(data, sheetName, index)
}

// 解析所有sheet中的所有数据
func ParseExcel(filePath string, processAllRow func([][]string, string), processRows ...ProcessRow) error {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Error closing file: ", err)
		}
	}()

	sheets := file.GetSheetList()
	//sheet
	for _, sheet := range sheets {
		rows, err := file.GetRows(sheet)
		if err != nil {
			return err
		}

		//处理所有行数据
		processAllRow(rows, sheet)
		// 行
		for iRow, row := range rows {
			if row == nil {
				continue
			}
			// 处理指定行数据
			for _, prFunc := range processRows {
				// ProcessIndex(prFunc, iRow)
				prFunc.Process(row, sheet, iRow)
			}
		}
	}
	return nil
}

func ParseSheets(filePath string, processAllRow func([][]string, string, int)) error {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Error closing file: ", err)
		}
	}()

	sheets := file.GetSheetList()
	//sheet
	for i, sheet := range sheets {
		rows, err := file.GetRows(sheet)
		if err != nil {
			return err
		}

		//处理所有行数据
		processAllRow(rows, sheet, i)
	}
	return nil
}
