package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-parrot/src/global"
	"go-parrot/src/model"
	"go-parrot/src/service"
	"net"
	"net/http"
)

func GeoIp() gin.HandlerFunc {
	return func(context *gin.Context) {
		ip := net.ParseIP(context.ClientIP())
		record, err := global.GeoIpDB.City(ip)
		if err != nil {
			global.Logger.Error(fmt.Sprintf("查询City发生错误：%s", err.Error()))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		var (
			city      string
			provinces string
			country   string
			continent string
			timeZone  string
			latitude  float64
			longitude float64
		)
		// 城市名称
		if len(record.City.Names) > 0 {
			city = record.City.Names["zh-CN"]
		}
		// 省份
		if len(record.Subdivisions) > 0 {
			if len(record.Subdivisions[0].Names) > 0 {
				provinces = record.Subdivisions[0].Names["zh-CN"]
			}
		}
		// 国家名
		if len(record.Country.Names) > 0 {
			country = record.Country.Names["zh-CN"]
		}
		// 洲名
		if len(record.Continent.Names) > 0 {
			continent = record.Continent.Names["zh-CN"]
		}
		// 时区
		timeZone = record.Location.TimeZone
		// 纬度
		latitude = record.Location.Latitude
		// 经度
		longitude = record.Location.Longitude
		var requestService = service.NewRequestService()
		//TODO:记录请求的用户信息
		err = requestService.RequestRecordAdd(model.Request{
			IP:        context.ClientIP(),
			URI:       context.Request.RequestURI,
			UID:       0,
			UserName:  "",
			City:      city,
			Provinces: provinces,
			Country:   country,
			Continent: continent,
			TimeZone:  timeZone,
			Latitude:  latitude,
			Longitude: longitude,
		})
		if err != nil {
			global.Logger.Error(fmt.Sprintf("记录geo信息发生错误：%s", err.Error()))
		}
	}
}
