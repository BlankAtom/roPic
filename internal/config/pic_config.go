package config

import (
	"context"
	"encoding/json"
	"github.com/blackswords/roPic/internal/log"
	"github.com/blackswords/roPic/internal/randr"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const PicConfigPath string = "../config/"

const PicConfigFilename string = "pic-config.json"

var picConfig *PicBedConfig

type TencentCOS struct {
	SecretKey        string `json:"secretKey"`
	SecretId         string `json:"secretId"`
	AppId            string `json:"appId"`
	Region           string `json:"region"`
	StorageSpaceName string `json:"storageSpaceName"`
	BucketName       string `json:"bucketName"`
	SystemVersion    string `json:"version"`
	MemoryPath       string `json:"path"`
}

type Github struct {
	RepoName     string `json:"repo-name"`
	BranchName   string `json:"branch-name"`
	Token        string `json:"token"`
	StoragePath  string `json:"storage-path"`
	DefineDomain string `json:"define-domain"`
}

type PicBedConfig struct {
	TencentCOS   *TencentCOS `json:"tencent-cos"`
	Github       *Github     `json:"github"`
	PicBedChoose int         `json:"pic-bed-choose"`
}

type PropertyEmptyError struct {
	name string
}

type UploadFailureError struct {
	msg string
}

// HasConfigFile 判断是否存在pic-config.json文件
func HasConfigFile() bool {
	fn, err := picConfigFullPath()
	if err != nil {
		return false
	}
	if _, err := os.Stat(fn); err != nil {
		return false
	}
	return true
}

// NewPicConfig Create new config file of pic-upload.
func NewPicConfig() error {
	return newConfig()
}

func InitPicConfig() {
	hasFile := HasConfigFile()
	if hasFile {
		println("File is existed.")
		err := LoadPicConfig()
		if err != nil {
			log.Warning(err.Error())
		}
	} else {
		println("File is not exist.")
		err := newConfig()
		if err != nil {
			log.Log(err.Error())
		}
		err = initPicConfig()
		if err != nil {
			log.Log(err.Error())
		}
	}
}

func LoadPicConfig() error {
	fn, err := picConfigFullPath()
	if err != nil {
		return err
	}
	f, err := os.ReadFile(fn)
	if err != nil {
		return err
	}

	var pc PicBedConfig
	err = json.Unmarshal(f, &pc)
	picConfig = &pc
	if err != nil {
		return err
	}

	return nil
}

func UploadFile(path string) (string, error) {
	var res = ""
	var err error = nil
	switch picConfig.PicBedChoose {
	case 0:
		res, err = tencentUpload(path)
	case 1:
		res, err = githubUpload(path)
	}

	return res, err
}

func (t *TencentCOS) CheckProperty() (propertyName string, isEmpty bool) {
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

func (g *Github) CheckProperty() (propertyName string, isEmpty bool) {
	if g.Token == "" {
		return "Token", true
	}
	if g.RepoName == "" {
		return "Token", true
	}
	if g.BranchName == "" {
		return "Token", true
	}
	if g.StoragePath == "" {
		return "Token", true
	}
	//if g.DefineDomain == "" {
	//	return "Token", true
	//}
	return "", false
}

// UploadFile 上传文件
func (t *TencentCOS) UploadFile(path string) string {
	u, _ := url.Parse("https://" + t.BucketName + ".cos." + t.Region + ".myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  t.SecretId,
			SecretKey: t.SecretKey,
		},
	})

	str := randr.RenamePicture()
	ts := strings.Split(path, ".")
	format := ts[len(ts)-1]

	key := t.MemoryPath + "/" + str + "." + format

	_, _, err := client.Object.Upload(context.Background(), key, path, nil)
	if err != nil {
		panic(err)
		return ""
	}

	return u.String() + "/" + key
}
func (g Github) UploadFile(path string) string {
	return ""
}

func (e *PropertyEmptyError) Error() string {
	return "属性为空"
}

// ================ private =========================

func newConfig() error {
	fullName, err := picConfigFullPath()
	if err != nil {
		return err
	}
	if _, err := os.Stat(fullName); err != nil {
		if os.IsNotExist(err) {
			// 创建文件
			f := createFile(fullName)
			// 关闭文件
			closeFile(f)
		} else {
			return err
		}
	}
	return nil
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Log(err.Error())
	}
}

func createFile(fullName string) *os.File {
	err := os.Mkdir(filepath.Dir(fullName), 0644)
	if err != nil {
		log.Warning(err.Error())
	}
	f, err := os.Create(fullName)
	if err != nil {
		log.Log(err.Error())
	}
	return f
}

func initPicConfig() error {
	picBed := PicBedConfig{
		Github: &Github{
			RepoName:     "",
			BranchName:   "",
			Token:        "",
			StoragePath:  "",
			DefineDomain: "",
		},
		TencentCOS: &TencentCOS{
			SecretKey:        "",
			SecretId:         "",
			AppId:            "",
			Region:           "",
			StorageSpaceName: "",
			BucketName:       "",
			SystemVersion:    "",
			MemoryPath:       "",
		},
		PicBedChoose: 0,
	}

	ouput, err := json.MarshalIndent(picBed, "", "\t\t")
	if err != nil {
		return err
	}

	fn, err := picConfigFullPath()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fn, ouput, 0644)
	if err != nil {
		return err
	}

	picConfig = &picBed
	return nil
}

func picConfigFullPath() (string, error) {
	return filepath.Abs(PicConfigPath + PicConfigFilename)
}

func tencentUpload(path string) (string, error) {
	if picConfig != nil && picConfig.TencentCOS != nil {
		if p, empty := picConfig.TencentCOS.CheckProperty(); empty {
			return "", &PropertyEmptyError{name: p}
		}
		res := picConfig.TencentCOS.UploadFile(path)
		return res, nil
	}
	return "", &PropertyEmptyError{name: "picConfig, TencentCOS"}
}

func githubUpload(path string) (string, error) {
	if picConfig != nil && picConfig.Github != nil {
		if p, empty := picConfig.Github.CheckProperty(); empty {
			return "", &PropertyEmptyError{name: p}
		}
		res := picConfig.Github.UploadFile(path)
		return res, nil
	}
	return "", &PropertyEmptyError{name: "picConfig, Github"}
}
