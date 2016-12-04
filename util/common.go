package util

import (
	"github.com/tealeg/xlsx"
	"github.com/extrame/xls"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
)

func ParseXlsxFile(bytes []byte) ([][]string, error) {
	xlFile, err := xlsx.OpenBinary(bytes)
	if err != nil {
		return [][]string{}, err
	}
	sheets, err := xlFile.ToSlice()
	if err != nil {
		return [][]string{}, err
	}
	if sheets != nil && len(sheets) > 0 {
		return sheets[0], nil
	}
	return [][]string{}, nil
}

func ParseXlsFile(filePath string) [][]string {
	xlFile, err := xls.Open(filePath, "utf-8")
	if err != nil {
		panic(err)
	}
	maxColumn := uint16(0)
	if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
		maxRow := int(sheet1.MaxRow)
		for i := 0; i <= maxRow; i++ {
			row := sheet1.Rows[uint16(i)]
			if row == nil {
				continue
			}
			for n, _ := range row.Cols {
				if n > maxColumn {
					maxColumn = n
				}
			}
		}
		records := [][]string{}
		for i := 0; i <= maxRow; i++ {
			records = append(records, make([]string, maxColumn + 1))
		}
		for i := 0; i <= maxRow; i++ {
			row := sheet1.Rows[uint16(i)]
			if row == nil {
				continue
			}
			for n, col := range row.Cols {
				values := col.String(xlFile)
				if values != nil && len(values) > 0 {
					records[i][n] = values[0]
				}

			}
		}
		return trimArray(records)
	}

	return [][]string{}
}

func SaveFile(bytes []byte) string {
	path := beego.AppConfig.String("tmp.path")
	fileName := path + bson.NewObjectId().Hex() + ".xls"
	err := ioutil.WriteFile(fileName, bytes, 0644)
	if err != nil {
		panic(err)
	}
	return fileName
}

func trimArray(records [][]string) [][]string {
	result := [][]string{}
	for _, list := range records {
		notValid := true
		for _, val := range list {
			if val != "" {
				notValid = false
			}
		}
		if !notValid {
			result = append(result, list)
		}
	}
	return result
}