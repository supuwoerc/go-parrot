package utils

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
)

// 拼接错误
func AppendError(prev, next error) error {
	if prev == nil {
		return next
	}
	return fmt.Errorf("%v,%w", prev, next)
}

// 获取请求的ip
func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}
	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	if net.ParseIP(ip) != nil {
		return ip, nil
	}
	return "", errors.New("未发现IP信息")
}
