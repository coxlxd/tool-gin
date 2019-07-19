package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

/**
目录下所有的文件夹
*/
func GetDirList(dirPath string) ([]string, error) {
	var dir_list []string
	dir_err := filepath.Walk(dirPath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dir_list = append(dir_list, path)
				return nil
			}

			return nil
		})
	return dir_list, dir_err
}

/**
 * 获取一个目录下所有文件信息，包含子目录
 *
 * @param null
 * @return
 * @Description
 * @author claer www.bajins.com
 * @date 2019/7/19 10:22
 */
func GetDirListAll(files []os.FileInfo, path string) []os.FileInfo {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			files = append(files, f)
		} else {
			currentPath := strings.Replace(path+"\\"+f.Name(), "\\", "/", -1)
			GetDirListAll(files, currentPath)
		}
		return nil
	})
	log.Fatal(err)
	return files
}

/**
 * 获取当前路径下所有文件
 *ioutil中提供了一个非常方便的函数函数ReadDir，他读取目录并返回排好序的文件以及子目录名([]os.FileInfo)
 *
 * @param path string 要查找的目录路径
 * @return
 * @Description
 * @author claer woytu.com
 * @date 2019/6/25 15:09
 */
func GetFileList(path string) []os.FileInfo {
	readerInfos, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if readerInfos == nil {
		return nil
	}
	return readerInfos
}

/**
判断路径
*/
func IsExistDir(dirPath string) bool {
	if IsStringEmpty(dirPath) {
		return false
	}
	_, err := os.Stat(dirPath)
	if err != nil || !os.IsExist(err) {
		return false
	}
	return true
}

/**
 * 判断所给路径文件/文件夹是否存在
 *
 * @param null
 * @return
 * @Description
 * @author claer www.bajins.com
 * @date 2019/6/25 15:34
 */
func Exists(path string) bool {
	if IsStringEmpty(path) {
		return false
	}
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/**
 * 判断文件是否存在：存在，返回true，否则返回false
 * 方法1
 *
 * @author claer www.bajins.com
 * @date 2019/6/28 11:31
 */
func IsFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println(info)
		return false
	}
	fmt.Println("exists", info.Name(), info.Size(), info.ModTime())
	return true
}

/**
 * 判断文件是否存在：存在，返回true，否则返回false
 * 方法2
 *
 * @author claer www.bajins.com
 * @date 2019/6/28 11:31
 */
func IsFileExist1(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

/**
 * 判断所给路径是否为文件夹
 *
 * @param null
 * @return
 * @Description
 * @author claer www.bajins.com
 * @date 2019/6/25 15:34
 */
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

/**
 * 判断所给路径是否为文件
 *
 * @param null
 * @return
 * @Description
 * @author claer www.bajins.com
 * @date 2019/6/25 15:33
 */
func IsFile(path string) bool {
	if !Exists(path) {
		return false
	}
	return !IsDir(path)
}

/**
 * 获取当前程序运行所在路径
 *
 * @param null
 * @return
 * @Description
 * @author claer www.bajins.com
 * @date 2019/6/25 15:34
 */
func OsPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
	}
	return dir
}

/**
 * 获取路径中的文件的后缀
 *
 * @param null
 * @return
 * @Description
 * @author claer www.bajins.com
 * @date 2019/6/25 15:35
 */
func GetSuffix(filePath string) string {
	ext := path.Ext(filePath)
	return ext
}

/**
 * 获取路径中的文件名
 *
 * @param null
 * @return
 * @Description
 * @author claer www.bajins.com
 * @date 2019/6/25 15:36
 */
func GetFileName(filePath string) string {
	ext := filepath.Base(filePath)
	return ext
}

/**
 * 获取路径中的目录及文件名
 *
 * @param null
 * @return
 * @Description
 * @author claer www.bajins.com
 * @date 2019/6/25 15:36
 */
func GetDirFile(filePath string) (dir, file string) {
	paths, fileName := filepath.Split(filePath)
	return paths, fileName
}

/**
 * 获取父级目录
 *
 * @author claer www.bajins.com
 * @date 2019/6/28 15:53
 */
func ParentDirectory(dirctory string) string {
	return path.Join(dirctory, "..")
}

/**
 * 目录分隔符转换
 *
 * @author claer www.bajins.com
 * @date 2019/6/28 15:53
 */
func CurrentDirectory() string {
	return strings.Replace(OsPath(), "\\", "/", -1)
}

/**
 * 获取上下文路径，传入指定目录截取前一部分
 *
 * @author claer woytu.com
 * @date 2019/6/29 3:22
 */
func ContextPath(root string) (path string, err error) {
	// 获取当前绝对路径
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	index := strings.LastIndex(dir, root)
	if len(dir) < len(root) || index <= 0 {
		err = errors.New("错误：路径不正确")
		return dir, err
	}
	return dir[0 : index+len(root)], nil
}

/**
 * 路径标准化拼接
 *
 * @param paths 可变路径参数
 * @return
 * @author claer woytu.com
 * @date 2019/6/29 3:46
 */
func PathStitching(paths ...string) string {
	sep := string(os.PathSeparator)
	path := ""
	for _, value := range paths {
		path = path + sep + value
	}
	return path[1:]
}
