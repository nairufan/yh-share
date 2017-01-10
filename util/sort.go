package util

type CountData struct {
	Val   string  `json:"val"`
	Count int     `json:"count"`
}

type SortableCountList []*CountData

func (list SortableCountList) Len() int {
	return len(list)
}

func (list SortableCountList) Less(i, j int) bool {
	if list[i].Count > list[j].Count {
		return true
	} else if list[i].Count < list[j].Count {
		return false
	} else {
		return list[i].Val > list[j].Val
	}
}

func (list SortableCountList) Swap(i, j int) {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}