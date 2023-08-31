package npm

import (
	"go-parrot/src/service"
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
func (p *PackageManagerService) DownloadsByTimeRange(start, end, packageName string) (*http.Response, error) {
	timeRange := strings.Join([]string{start, end}, ":")
	requestUrl := strings.Join([]string{DownloadsByTimeRange, timeRange, packageName}, "/")
	return http.Get(requestUrl)
}
