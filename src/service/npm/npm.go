package npm

import (
	"go-parrot/src/service"
	"strings"
)

const (
	DownloadsByTimeRange = "https://api.npmjs.org/downloads/range"
	GetPackageInfo       = "https://registry.npmjs.org"
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
	return p.GetRemoteURL(requestUrl)
}

func (p *PackageManagerService) GetPackageInfo(packageName string) (any, error) {
	requestUrl := strings.Join([]string{GetPackageInfo, packageName}, "/")
	return p.GetRemoteURL(requestUrl)
}
