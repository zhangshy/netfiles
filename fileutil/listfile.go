package fileutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//函数首字母大写才能被其他包访问
func GetCurrentDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

//返回当前目录下的文件，不往下遍历
func GetCurrentDirFiles(dirname string) ([]string, error) {
	filenames := make([]string, 0)
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return filenames, err
	}
	for _, fi := range files {
		if !fi.IsDir() {
			filenames = append(filenames, fi.Name())
		}
	}
	return filenames, nil
}

//遍历当前目录，获取指定后缀的文件
func GetCurrentDirAllFiles(dirname string, suffix string) ([]string, error) {
	filenames := make([]string, 0)
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), suffix) {
			filenames = append(filenames, path)
		}
		return nil
	})
	return filenames, err
}
