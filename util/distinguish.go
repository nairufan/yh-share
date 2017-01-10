package util

import (
	"strings"
)

func TelRecords(records [][]string) []string {
	minRatio := 0.8
	if records == nil || len(records) == 0 {
		return []string{}
	}
	telsCount := make([]int, len(records[0]))
	for _, line := range records {
		for index, cell := range line {
			if isTel(cell) {
				telsCount[index] += 1;
			}

		}
	}
	telIndex := maxIndex(telsCount)
	maxCount := telsCount[telIndex]
	if float64(maxCount) / float64(len(records)) > minRatio {
		return getColumn(records, telIndex)
	}
	return []string{}
}

func isTel(str string) bool {
	if strings.HasPrefix(str, "1") && len(str) == 11 {
		return true
	}
	return false
}

func maxIndex(numArray []int) int {
	max := numArray[0]
	idx := 0
	for index, v := range numArray {
		if max < v {
			max = v
			idx = index
		}
	}
	return idx
}

func getColumn(records [][]string, column int) []string {
	list := []string{}
	for _, record := range records {
		if (isTel(record[column])) {
			list = append(list, record[column])
		}
	}
	return list
}