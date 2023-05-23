package excelTable

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

type ExcelTable struct {
	FileName string
	file     *excelize.File
	rows     map[string]*Rows
}

// NewExcelTable 获取Excel为一个对象
func NewExcelTable(fileName string) *ExcelTable {
	result := &ExcelTable{
		FileName: fileName,
		file:     openExcel(fileName),
		rows:     make(map[string]*Rows),
	}
	return result
}

// GetRows 传入sheet，获取它的每一行
func (s *ExcelTable) GetRows(sheet string) *Rows {
	if rows, ok := s.rows[sheet]; ok {
		return rows
	} else {
		_rows, _ := s.file.GetRows(sheet)
		rows = &Rows{rows: _rows}
		s.rows[sheet] = rows
		return rows
	}
}

// GetSheets 获取Excel表中的 sheets名称到数组
func (s *ExcelTable) GetSheets() []string {
	return s.file.GetSheetList()
}

// SaveToFile 保存Excel表到文件
func (s *ExcelTable) SaveToFile(fileName string) {

	for sheet, row := range s.rows {
		metrics, err := s.file.GetMergeCells(sheet)
		if err != nil {
			fmt.Println(err)
		}
		for i, v := range row.rows {
			params := strSliceToInterfaseSlice(v)
			if err := s.file.SetSheetRow(sheet, fmt.Sprintf("a%d", i+1), &params); err != nil {
				fmt.Println(err)
			}
		}

		{
			cols, err := s.file.GetCols(sheet)
			fmt.Println(cols, err)
		}
		fmt.Println(metrics)
	}
	s.file.SaveAs(fileName)
}

func strSliceToInterfaseSlice(strs []string) []interface{} {
	var result []interface{}
	for _, v := range strs {
		result = append(result, v)
	}
	return result
}
