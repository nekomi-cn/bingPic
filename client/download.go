package client

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ZMuSiShui/bingImg/conf"
	"github.com/ZMuSiShui/bingImg/utils"
	"github.com/matryer/try"
)

func loadFromBing() (image conf.BingImagesDoc, err error) {
	var request *http.Request
	request, err = http.NewRequest(http.MethodGet, conf.BingURL, nil)
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 3,
			DisableKeepAlives:   false,
		},
		Timeout: time.Duration(6) * time.Second,
	}
	var resp *http.Response

	err = try.Do(func(attempt int) (bool, error) {
		var rErr error
		resp, rErr = client.Do(request)
		return attempt < 3, rErr
	})

	if err != nil {
		log.Errorf("%v", err)
		return
	}
	defer resp.Body.Close()

	var syncRespBodyBytes []byte
	syncRespBodyBytes, err = getResponseBody(resp)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	err = json.Unmarshal(syncRespBodyBytes, &image)
	log.Infof("%v", image)
	return
}

// 获取响应结构体
func getResponseBody(resp *http.Response) (body []byte, err error) {
	var output io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		output, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Errorf("%v", err)
			return
		}
		if err != nil {
			log.Errorf("%v", err)
			return
		}
	default:
		output = resp.Body
		if err != nil {
			log.Errorf("%v", err)
			return
		}
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(output)
	if err != nil {
		return
	}
	body = buf.Bytes()
	return
}

func DownImage() (err error) {
	loaded, err := loadFromBing()
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	imgPath := conf.WriteToFile + time.Now().Format("2006-01-01") + ".jpg"
	imgUrl := "https://www.bing.com" + loaded.Images[0].Url
	log.Infof("开始下载，目标地址: %s", imgUrl)
	res, err := http.Get(imgUrl)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	Write(imgPath, b)

	return
}

func Write(path string, src []byte) bool {
	var file *os.File
	if utils.FileExists(path) {
		var err error
		file, err = os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Errorf("文件打开失败.")
		}
	} else {
		var err error
		file, err = utils.CreatNestedFile(path)
		if err != nil {
			return false
		}
	}
	defer func() {
		_ = file.Close()
	}()
	_, err := file.Write(src)

	return err == nil
}
