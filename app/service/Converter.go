package service

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"os"
	"path/filepath"
)

func OleConvert(fullFileName string) (string, error) {
	path, err := os.Getwd()

	fileExt := filepath.Ext(fullFileName)
	fileName := fullFileName[:len(fullFileName)-len(fileExt)]

	pathToFile := path + "/FilesToConvert/" + fullFileName
	outputPath := path + "/ConvertedFiles/" + fileName + ".pdf"

	err = ole.CoInitialize(0)
	unknown, err := oleutil.CreateObject("Word.Application")
	word, err := unknown.QueryInterface(ole.IID_IDispatch)

	documents := oleutil.MustGetProperty(word, "Documents").ToIDispatch()
	document := oleutil.MustCallMethod(documents, "Open", pathToFile).ToIDispatch()

	oleutil.MustCallMethod(document, "SaveAs", outputPath, 17).ToIDispatch()
	oleutil.CallMethod(documents, "Close")
	oleutil.CallMethod(word, "Quit")
	word.Release()

	return outputPath, err
}
