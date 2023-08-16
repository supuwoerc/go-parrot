package test

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type DirInfo struct {
	Dir      string     `json:"dir"`
	Children *[]DirInfo `json:"children"`
}

// 当前文件对路径
var currentDir string

// 项目所在路径
var projectDir string

// 路径分隔符
var fileSeparator string

const json_file_name = "project_dir.json"

// 初始化变量
func init() {
	fileSeparator = string(filepath.Separator)
	currentDir, _ = os.Getwd()
	projectDir = currentDir[0:strings.LastIndex(currentDir, fileSeparator)]
}

// 读取json文件
func loadProjectDirFile() (*DirInfo, error) {
	fileContent, err := os.ReadFile(strings.Join([]string{currentDir, json_file_name}, fileSeparator))
	if err != nil {
		return nil, err
	}
	dirInfo := &DirInfo{}
	err = json.Unmarshal(fileContent, dirInfo)
	if err != nil {
		return nil, err
	}
	return dirInfo, nil
}

// 根据json文件获取要创建目录名数组
func parseDir(dirInfos *[]DirInfo, basicDir string, result *[]string) {
	for _, val := range *dirInfos {
		if val.Dir != "" {
			currentDir := strings.Join([]string{basicDir, val.Dir}, fileSeparator)
			*result = append(*result, currentDir)
			if val.Children != nil {
				parseDir(val.Children, currentDir, result)
			}
		}
	}
}

// 创建项目目录
func TestGenerateDirByJson(t *testing.T) {
	jsonContent, err := loadProjectDirFile()
	if err != nil {
		panic("读取json文件发生错误，请检查json文件" + err.Error())
	}
	var dirs []string
	parseDir(&[]DirInfo{*jsonContent}, projectDir, &dirs)
	for _, val := range dirs {
		err := os.MkdirAll(val, fs.ModePerm)
		if err != nil {
			t.Errorf("创建目录失败%s，失败原因%s", val, err.Error())
		}
	}
	t.Logf("当前项目目录%s", projectDir)
}
