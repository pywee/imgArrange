package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type FileInfo struct {
	// DirName 新创建的目录名称
	DirName string `json:"dirName"`
	// PreFix 文件移动后是否在文件名称前面加上年月日
	PreFix int `json:"ymd,omitempty"`
	// Suffix 当前目录包含的后缀
	Suffix []string `json:"suffix"`
	// Count 计数器
	Count int
}

var cfg []*FileInfo

func main() {
	ret, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err)
	}

	// 获取配置文件
	cfg = getConfig("./config.json")
	var (
		fix  string
		conf *FileInfo
	)
	for _, v := range ret {
		if v.IsDir() {
			continue
		}
		if fix = getFix(v.Name()); fix == "" {
			continue
		}
		if conf = getPathNameFromFileFix(fix); conf == nil {
			continue
		}

		dirName := fmt.Sprintf("./%d", v.ModTime().Year()) + "/" + conf.DirName
		if err = os.MkdirAll(dirName, os.ModePerm); err != nil {
			panic("创建目录失败, 原因: " + err.Error())
		}

		ymd := ""
		modTime := v.ModTime()
		if conf.PreFix == 1 {
			ymd = fmt.Sprintf("%d%02d%02d_", modTime.Year(), modTime.Month(), modTime.Day())
		}
		newFileName := dirName + "/" + ymd + v.Name()
		if err = os.Rename(v.Name(), newFileName); err != nil {
			fmt.Println("移动文件失败, 原因: " + err.Error())
			continue
		}

		// 保留文件原来的 atime mtime 时间
		// _ = os.Chtimes(newFileName, modTime, modTime)

		fmt.Println("正在整理文件:", newFileName)
		conf.Count++
	}

	var msg string
	for _, v := range cfg {
		if v.Count > 0 {
			msg += fmt.Sprintf("移动了%d个文件到【%s】文件夹\n", v.Count, v.DirName)
		}
	}
	fmt.Println("\n" + msg)
	fmt.Println("全部已完成, 5秒后退出程序..")
	time.Sleep(time.Second * 5)
}

// getConfig 读取 JSON 配置文件
func getConfig(path string) []*FileInfo {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		panic("读取文件失败, 原因: " + err.Error())
	}

	var c []*FileInfo
	if err = json.Unmarshal(bs, &c); err != nil {
		panic("配置文件格式有误")
	}
	return c
}

// getFix 获取文件的后缀名
func getFix(name string) string {
	if idx := strings.LastIndex(name, "."); idx != -1 {
		return name[idx+1:]
	}
	return ""
}

// getPathNameFromFileFix
// 判断当前文件是否在配置中
// 是否需要移动
func getPathNameFromFileFix(sep string) *FileInfo {
	for _, arr := range cfg {
		for _, v2 := range arr.Suffix {
			if strings.EqualFold(strings.Trim(v2, "."), sep) {
				return arr
			}
		}
	}
	return nil
}

// copyFile 复制文件
// func copyFile(src, dst string) error {
// 	sourceFileStat, err := os.Stat(src)
// 	if err != nil {
// 		return err
// 	}

// 	if !sourceFileStat.Mode().IsRegular() {
// 		return err
// 	}

// 	source, err := os.Open(src)
// 	if err != nil {
// 		return err
// 	}
// 	defer source.Close()

// 	destination, err := os.Create(dst)
// 	if err != nil {
// 		return err
// 	}
// 	defer destination.Close()

// 	_, err = io.Copy(destination, source)
// 	return err
// }
