package excelTable

import "fmt"

type Rows struct {
	rows [][]string
}

const TitleIndex = 2

func (s *Rows) GetTitleCellIndex(str string) int {
	return findIndex(s.rows[TitleIndex], str)
}

// 获取包含指定标题内容的所有行
func (s *Rows) GetFlagRows(flag, value string) ([][]string, error) {
	index := s.GetTitleCellIndex(flag)
	if index < 0 {
		return nil, fmt.Errorf("Not Find Flag <%s>", flag)
	}
	var result [][]string
	for _, v := range s.rows {
		if len(v) <= index {
			continue
		}
		if v[index] == value {
			result = append(result, v)
		}
	}
	return result, nil
}

func (s *Rows) CompareAndGetRows(rows [][]string, count int, newflag string) ([][]string, error) {
	var result [][]string
	index := s.GetTitleCellIndex(newflag)
	if index < 0 {
		return result, fmt.Errorf("Not find flag %s", newflag)
	}
	for _, _vv := range rows {
		// 遍历旧表内容
		for _, v := range s.rows {
			if compareStrings(v, _vv, count) {
				_vv = append(_vv, v[index])
				result = append(result, _vv)
				break
			}
		}
	}
	return result, nil
}

func (s *Rows) CompareAndWriteGetRows(rows [][]string, count int) ([][]string, error) {
	for i := 0; i < len(s.rows); i++ {
		for _, _vv := range rows {
			// 遍历旧表内容
			if compareStrings(s.rows[i], _vv, count) {
				s.rows[i] = _vv
				break
			}
		}
	}
	return s.rows, nil
}

func (s *Rows) SetValue(rows [][]string) error {
	s.rows = rows
	return nil
}
