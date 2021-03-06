package util

import (
	"io/ioutil"
	"os"
)

//func OverWriteFile(pathFile string, content *string) (err error) {
//	file, err := os.OpenFile(pathFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
//	if err != nil {
//		return
//	}
//	defer file.Close()
//	_, err = io.WriteString(file, *content)
//	if err != nil {
//		return
//	}
//	return
//}

func OverWriteFile(pathFile string, data []byte) (err error) {
	err = ioutil.WriteFile(pathFile, data, 0644) //写入文件(字节数组)
	if err != nil {
		return
	}
	return
}

func PathFileExists(pathFile string) bool {
	_, err := os.Stat(pathFile)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func MkdirAll(path string) (err error) {
	_, err = os.Stat(path)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return
		}
		return
	}
	return
}
