package npm

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-parrot/src/service"
	"io"
	"net/http"
	"strings"
)

const (
	DownloadsByTimeRange = "https://api.npmjs.org/downloads/range"
)

type PackageManagerService struct {
	service.BasicService
}

var packageService *PackageManagerService

func NewPackageService() *PackageManagerService {
	if packageService == nil {
		packageService = &PackageManagerService{
			BasicService: service.NewBasicService(),
		}
	}
	return packageService
}

// 根据包名和时间范围获取下载数据 https://api.npmjs.org/downloads/range/2023-01-01:2023-01-31/express
func (p *PackageManagerService) DownloadsByTimeRange(start, end, packageName string) (any, error) {
	timeRange := strings.Join([]string{start, end}, ":")
	requestUrl := strings.Join([]string{DownloadsByTimeRange, timeRange, packageName}, "/")
	resp, err := http.Get(requestUrl)
	if err != nil {
		return "", err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("请求发生错误,状态码：%d，响应结果：%s", resp.StatusCode, string(body)))
	}
	if err != nil {
		return "", err
	}
	var bodyJson any
	err = json.Unmarshal(body, &bodyJson)
	if err != nil {
		return "", err
	}
	return bodyJson, err
}
