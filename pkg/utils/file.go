package utils

import (
	"bufio"
	"fmt"
	"github.com/woaijssss/tros/pkg/utils/encrypt"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return fmt.Errorf("file is same ")
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

// ReadFile 获取本地文件内容
func ReadFile(file string) string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("read file fail", err)
		return ""
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}

	return string(fd)
}

// SaveFile 将文本内容保存到本地
func SaveFile(filepath, content string) error {
	// 创建一个文件
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将字符串写入文件
	return ioutil.WriteFile(file.Name(), []byte(content), 0644)
}

// GetSize get the file size
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

func GetFileSize(filename string) (int, error) {
	fileObj, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	size, err := GetSize(fileObj)
	if err != nil {
		return 0, err
	}

	return size, nil
}

// GetExt get the file ext
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckNotExist check if the file exists
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckPermission check if the file has permission
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir create a directory if it does not exist
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir create a directory
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open a file according to a specific mode
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// MustOpen maximize trying to open the file
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}

func ReadJsonFile(filePath string) (result string) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		//fmt.Println("ERROR:", err)
		return ""
	}
	buf := bufio.NewReader(file)
	for {
		s, err := buf.ReadString('\n')
		result += s
		if err != nil {
			if err == io.EOF {
				//fmt.Println("Read is ok")
				break
			} else {
				//fmt.Println("ERROR:", err)
				return ""
			}
		}
	}
	return result
}

func GetFileFullName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = fileName + strconv.Itoa(rand.Int())
	return fileName + ext
}

func GetFilename(filePath string) string {
	strArr := strings.Split(filePath, "/")
	return strArr[len(strArr)-1]
}

// CheckFileExt check image file ext
func CheckFileExt(fileName string, allowExts []string) bool {
	ext := GetExt(fileName)
	for _, allowExt := range allowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

// CheckFileSize check image size
func CheckFileSize(f multipart.File, fileSize int) bool {
	size, err := GetSize(f)
	if err != nil {
		return false
	}

	return size <= fileSize
}

// CheckFileExist check if the file exists
func CheckFileExist(src string) bool {
	_, err := os.Stat(src)
	return !os.IsNotExist(err)
}

//func CheckFileExist(src string) (bool, error) {
//	dir, err := os.Getwd()
//	if err != nil {
//		return false, fmt.Errorf("os.Getwd err: %v", err)
//	}
//
//	err = IsNotExistMkDir(dir + "/" + src)
//	if err != nil {
//		return false, fmt.Errorf("file.IsNotExistMkDir err: %v", err)
//	}
//
//	perm := CheckPermission(src)
//	if perm == true {
//		return false, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
//	}
//
//	return true, nil
//}

// CheckPathExist 检测文件夹路径时候存在
func CheckPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileName(fileFullPath string) string {
	//fileFullPath := "/mnt/data/ds-pipeline/caplog/2020/06/28/01/827a0b7edf92808152305cffb1d559ff.jpg"
	fileName := filepath.Base(fileFullPath)
	fileNameList := strings.Split(fileName, ".")
	if len(fileNameList) == 0 {
		return encrypt.EncodeMD5(fileFullPath)
	}
	return fileNameList[0]
}

func GetFileBaseName(fileFullPath string) string {
	return filepath.Base(fileFullPath)
}

func GetFileModifyTime(fileFullPath string) int64 {
	fileInfo, _ := os.Stat(fileFullPath)
	return fileInfo.ModTime().Unix()
}

func GetCommonFileSize(fileFullPath string) (int64, error) {
	fileInfo, err := os.Stat(fileFullPath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

func DeleteFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	err = os.Remove(filepath)
	if err != nil {
		return err
	}
	return nil
}

func GetFullHourString(timeIn int64) (string, string) {
	nowTime := int64(0)
	if timeIn == 0 {
		nowTime = time.Now().Unix()
	} else {
		nowTime = timeIn
	}
	t := time.Unix(nowTime, 0)
	year := strconv.Itoa(t.Year())
	month := getFormat(int(t.Month()))
	day := getFormat(int(t.Day()))
	hour := getFormat(int(t.Hour()))
	return year + "/" + month + "/" + day + "/" + hour, hour
}

func getFormat(input int) string {
	if input < 10 {
		return "0" + strconv.Itoa(input)
	} else {
		return strconv.Itoa(input)
	}
}

func GetHourString(timeIn int64) string {
	nowTime := int64(0)
	if timeIn == 0 {
		nowTime = time.Now().Unix()
	} else {
		nowTime = timeIn
	}
	t := time.Unix(nowTime, 0)
	hour := getFormat(int(t.Hour()))
	return hour
}

func GetLastDayString(timeIn int64) (string, string) {
	nowTime := int64(0)
	if timeIn == 0 {
		nowTime = time.Now().Unix()
	} else {
		nowTime = timeIn
	}
	t := time.Unix(nowTime-86400, 0)
	year := strconv.Itoa(t.Year())
	month := getFormat(int(t.Month()))
	day := getFormat(int(t.Day()))
	return year + "/" + month + "/" + day, day
}
