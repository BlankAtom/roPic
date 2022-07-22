package picbed

import (
	"context"
	"fmt"
	error2 "github.com/blackswords/roPic/internal/error"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

// CosConfig 配置类
type CosConfig struct {
	SecretKey        string `json:"secretKey"`        // SecretKey
	SecretId         string `json:"secretId"`         // SecretId
	AppId            string `json:"appId"`            // AppId
	Region           string `json:"region"`           // Region
	StorageSpaceName string `json:"storageSpaceName"` // StorageSpaceName 存储空间名
	BucketName       string `json:"bucketName"`       // BucketName 存储桶名字
	SystemVersion    string `json:"version"`          // SystemVersion 系统版本，通常是5
	MemoryPath       string `json:"path"`             // MemoryPath 存储路径
}

var client *cos.Client

func (t *CosConfig) CheckProperty() (propertyName string, isEmpty bool) {
	if t.AppId == "" {
		return "AppId", true
	}
	if t.SecretId == "" {
		return "SecretId", true
	}
	if t.SecretKey == "" {
		return "SecretKey", true
	}
	if t.SystemVersion == "" {
		return "SystemVersion", true
	}
	if t.Region == "" {
		return "Region", true
	}
	if t.BucketName == "" {
		return "BucketName", true
	}
	if t.StorageSpaceName == "" {
		return "StorageSpaceName", true
	}
	if t.MemoryPath == "" {
		return "MemoryPath", true
	}

	return "", false
}

// UploadFile 上传文件
func (t *CosConfig) UploadFile(path string, fileName string) string {
	key := t.MemoryPath + "/" + fileName

	complete, _, err := client.Object.Upload(context.Background(), key, path, nil)
	if err != nil {
		panic(err)
		return ""
	}

	return complete.Location
}

func (t *CosConfig) SearchBucketsList() {

}

func (t *CosConfig) InitClient() error {
	pn, pass := t.CheckProperty()
	if pass {
		return &error2.PropertyEmptyError{PropertyName: pn}
	}

	urlString := fmt.Sprintf("https://%s.cos.%s.myqcloud.com", t.BucketName, t.Region)
	u, _ := url.Parse(urlString)
	b := &cos.BaseURL{BucketURL: u}
	client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  t.SecretId,
			SecretKey: t.SecretKey,
		},
	})

	return nil
}
