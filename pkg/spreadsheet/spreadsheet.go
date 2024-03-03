package spreadsheet

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ParseSpreadsheet(filePath string) ([]map[string]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening spreadsheet file %w", err)
	}
	defer f.Close()

	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("error readings rows: %w", err)
	}

	// example for now

	var results []map[string]string
	var headers []string
	for i, row := range rows {
		if i == 0 {
			headers = row
			continue
		}
		rowMap := make(map[string]string)
		for j, cell := range row {
			if j < len(headers) {
				rowMap[headers[j]] = cell
			}
		}
		results = append(results, rowMap)
	}

	return results, nil
}
