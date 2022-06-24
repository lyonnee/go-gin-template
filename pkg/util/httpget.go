package util

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/LyonNee/app-layout/pkg/log"

	"go.uber.org/zap"
)

func HttpGet(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Timeout: 30 * time.Second, Transport: tr}

	response, err := client.Do(req)
	if err != nil {
		log.ZapLogger().Warn("发起http get请求失败", zap.String("url", url), zap.Error(err))
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.ZapLogger().Warn("读取response body失败", zap.String("url", url), zap.Error(err))
		return nil, err
	}

	return body, nil
}
