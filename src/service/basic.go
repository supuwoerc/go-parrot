package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-parrot/src/conf"
	"go-parrot/src/global"
	"io"
	"net/http"
)

type BasicService struct {
	RedisClient *conf.RedisClient
}

func NewBasicService() BasicService {
	return BasicService{
		RedisClient: global.RedisClient,
	}
}

func (basic *BasicService) GetRemoteURL(url string) (any, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("请求发生错误,状态码：%d，响应结果：%s", resp.StatusCode, string(body)))
	}
	if err != nil {
		return nil, err
	}
	var bodyJson any
	err = json.Unmarshal(body, &bodyJson)
	if err != nil {
		return nil, err
	}
	return bodyJson, err
}
