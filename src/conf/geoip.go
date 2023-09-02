package conf

import (
	"github.com/oschwald/geoip2-golang"
)

func InitGeoIpDB() (*geoip2.Reader, error) {
	return geoip2.Open("./static/GeoLite2-City.mmdb")
}
