package util

import (
	"mime/multipart"
	"github.com/tealeg/xlsx"
	"github.com/astaxie/beego"
)

type Size interface {
	Size() int64
}

func ParseFile(file multipart.File) [][]string {
	size := file.(Size).Size()
	bytes := make([]byte, size)
	_, err := file.Read(bytes)
	if err != nil {
		panic(err)
	}
	xlFile, err := xlsx.OpenBinary(bytes)
	if err != nil {
		beego.Error(err)
		panic(err)
	}
	sheets, err := xlFile.ToSlice()
	if err != nil {
		panic(err)
	}
	if sheets != nil && len(sheets) > 0 {
		return sheets[0]
	}
	return nil
}