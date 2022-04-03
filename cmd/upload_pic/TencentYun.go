package main

import (
	"context"
	"github.com/blackswords/roPic/internal/randr"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gopkg.in/yaml.v2"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TencentYun struct {
	BucketFromTencent BucketConfig `yaml:"txyun"`
}
type BucketConfig struct {
	SecretKey        string `yaml:"secretKey"`
	SecretId         string `yaml:"secretId"`
	AppId            string `yaml:"appId"`
	Region           string `yaml:"region"`
	StorageSpaceName string `yaml:"storageSpaceName"`
	BucketName       string `yaml:"bucketName"`
	SystemVersion    string `yaml:"version"`
	MemoryPath       string `yaml:"path"`
}

// InitConfig 读取配置文件，并对当前对象进行赋值
// 返回当前对象
func (tencentYun *TencentYun) InitConfig() (self *TencentYun) {
	file, err := os.ReadFile(configPathName)
	if err != nil {
		println(err.Error())
		return tencentYun
	}
	err = yaml.Unmarshal(file, tencentYun)
	if err != nil {
		println(err.Error())
		return tencentYun
	}
	return tencentYun
}

// UploadFile 上传文件，参数p是文件路径，参数t是配置的实例化对象
// 返回文件名和完整连接
func UploadFile(p string, t TencentYun) (filename string, uri string) {

	u, _ := url.Parse("https://" + t.BucketFromTencent.BucketName + ".cos." + t.BucketFromTencent.Region + ".myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  t.BucketFromTencent.SecretId,
			SecretKey: t.BucketFromTencent.SecretKey,
		},
	})

	str := randr.RenamePicture()
	ts := strings.Split(p, ".")
	format := ts[len(ts)-1]

	key := t.BucketFromTencent.MemoryPath + "/" + str + "." + format

	_, _, err := client.Object.Upload(context.Background(), key, p, nil)
	if err != nil {
		panic(err)
	}

	return str + "." + format, u.String() + "/" + key
}
