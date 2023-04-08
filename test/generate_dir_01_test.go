package test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type DirDesc struct {
	Text     string     `json:"text"`
	Children *[]DirDesc `json:"children"`
}

var currentDir string        //当前目录
var workspaceRootDir string  //项目根目录
var filePathSeparator string //路径分割符
var jsonContent DirDesc      //json文件内容

const dirJsonFileName = "dir.json"

// 读取路径描述
func loadJson() {
	currentDir, _ = os.Getwd()
	filePathSeparator = string(filepath.Separator)
	workspaceRootDir = currentDir[0:strings.LastIndex(currentDir, filePathSeparator)]
	jsonFile, _ := os.ReadFile(strings.Join([]string{currentDir, dirJsonFileName}, filePathSeparator))
	err := json.Unmarshal(jsonFile, &jsonContent)
	if err != nil {
		panic("序列化json文件失败：" + err.Error())
	} else {
		fmt.Println("读取文件成功")
	}
}

// 根据树状数据返回需要创建对目录路径集合
func parseMap(data *[]DirDesc, prePath string, result []string) []string {
	for _, val := range *data {
		currentPath := strings.Join([]string{prePath, val.Text}, filePathSeparator)
		result = append(result, currentPath)
		if val.Children != nil {
			return parseMap(val.Children, currentPath, result)
		}
	}
	return result
}

// 根据路径创建对应的可读写目录
func createDirByPaths(paths []string, baseDir string) {
	for _, val := range paths {
		err := os.MkdirAll(strings.Join([]string{baseDir, val}, filePathSeparator), fs.ModePerm)
		if err != nil {
			panic("创建目录失败：" + err.Error())
		}
	}
}
func TestGenerateDir(t *testing.T) {
	loadJson()
	var paths []string
	paths = parseMap(&[]DirDesc{jsonContent}, "", paths)
	createDirByPaths(paths, workspaceRootDir)
}
