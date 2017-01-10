package util

import (
	"github.com/tealeg/xlsx"
	"github.com/extrame/xls"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"math/rand"
	"sort"
)

func parseXlsFile(file *xlsx.File) ([][]string, error) {
	sheets, err := file.ToSlice()
	if err != nil {
		return [][]string{}, err
	}
	if sheets != nil && len(sheets) > 0 {
		return sheets[0], nil
	}
	return [][]string{}, nil
}

func ParseXlsxFile(bytes []byte) ([][]string, error) {
	xlFile, err := xlsx.OpenBinary(bytes)
	if err != nil {
		return [][]string{}, err
	}
	return parseXlsFile(xlFile)
}

func ParseXlsxFileWithPath(path string) ([][]string, error) {
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		return [][]string{}, err
	}
	return parseXlsFile(xlFile)
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

func GetRandomString(size int) string {
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z"}
	charsLen := len(chars)
	result := ""
	for i := 0; i < size; i++ {
		index := rand.Intn(charsLen)
		result += chars[index]
	}
	return result
}

func GetOrderedList(list []string) SortableCountList {
	result := SortableCountList{}
	if len(list) == 0 {
		return result
	}
	countMap := map[string]int{}
	for _, str := range list {
		countMap[str] += 1
	}
	for k, v := range countMap {
		result = append(result, &CountData{
			Val: k,
			Count: v,
		})
	}
	sort.Sort(result)
	return result
}