package bin

import (
	"os"
	"path/filepath"
)

type fileUtilImpl struct{}

type fileUtil interface {
	FileExist(path string) (bool, error)                              // FileExist 文件是否存在
	FileIsDir(path string) (bool, error)                              // FileIsDir 文件是否为文件夹
	NotExistMkdir(path string) error                                  // NotExistMkdir 文件夹不存在则创建文件夹
	FindFileByFolderChildren(folder, suffix string) ([]string, error) // FindFileByFolderChildren 查询所有文件夹下级文件
}

// FileUtil 文件工具实现
var FileUtil fileUtil = (*fileUtilImpl)(nil)

// FileExist 文件是否存在
// path: 文件地址
// bool: 是否存在-存在true,不存在false
// error: 错误信息
func (*fileUtilImpl) FileExist(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// FileIsDir 文件是否为文件夹
// path: 文件地址
// bool: 是否为文件夹-是true,不是false
// error: 错误信息
func (*fileUtilImpl) FileIsDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

// NotExistMkdir 文件夹不存在则创建文件夹
// path: 文件地址
// error: 错误信息
func (*fileUtilImpl) NotExistMkdir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(path, os.ModePerm)
		}
		return err
	}
	return nil
}

// FindFileByFolderChildren 查询所有文件夹下级文件
// path: 文件夹地址
// suffix: 文件后缀
// []string: 该文件夹下为该文件后缀的文件列表
// error: 错误信息
func (f *fileUtilImpl) FindFileByFolderChildren(path, suffix string) ([]string, error) {
	isDir, err := f.FileIsDir(path)
	if err != nil {
		return nil, err
	}
	if !isDir {
		fileNameWithSuffix := filepath.Base(path)
		fileSuffix := filepath.Ext(fileNameWithSuffix)
		if suffix == "" || fileSuffix == suffix {
			return []string{path}, nil
		}
		return nil, nil
	}
	dirEntrys, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for _, dirEntry := range dirEntrys {
		if dirEntry.IsDir() {
			files, err := f.FindFileByFolderChildren(path+"/"+dirEntry.Name(), suffix)
			if err != nil {
				continue
			}
			result = append(result, files...)
		} else {
			fileName := path + "/" + dirEntry.Name()
			fileNameWithSuffix := filepath.Base(fileName)
			fileSuffix := filepath.Ext(fileNameWithSuffix)
			if suffix == "" || fileSuffix == suffix {
				result = append(result, fileName)
			}
		}
	}
	return result, nil
}
