package upload

import (
	"contrplatform/configs"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeContract = iota + 1

func GetSavePath() string {
	return configs.UploadSavePath
}

func CheckSavePathNotExist(dst string) bool {
	_,err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CreatSavePath(dst string, perm os.FileMode) error {
	if err:=os.MkdirAll(dst,perm); err != nil {
		return err
	}
	return nil
}


func CheckContainExt(t FileType,filename string) bool {
	ext := path.Ext(filename)
	switch t {
	case TypeContract:
		return strings.ToLower(ext)==strings.ToLower(configs.UploadContrExt)
	}
	return false
}

func CheckOutMaxSize(t FileType, f multipart.File) bool {
	content,_ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeContract:
		return size >= configs.UploadContrMaxSize*1024*1024
	}
	return false
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out,err:=os.Create(dst)
	defer out.Close()
	_,err = io.Copy(out,src)
	return err
}