package picbed

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

const PicConfigPath string = "../config/"

const PicConfigFilename string = "pic-config.json"

var picConfig *Config

type Config struct {
	TencentCOS   *CosConfig `json:"tencent-rocos"`
	Github       *Github    `json:"github"`
	PicBedChoose int        `json:"pic-bed-choose"`
}

func LoadConfiguration() {
	fn, err := picConfigFullPath()
	if err != nil {
		panic(err)
	}

	hasFile := hasConfigFile()
	if !hasFile {
		return
	}

	f, err := os.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	var pc Config
	err = json.Unmarshal(f, &pc)
	picConfig = &pc
	if err != nil {
		panic(err)
	}
}

// InitConfiguration 当文件存在时跳过，当文件不存在时创建文件 并初始化
func InitConfiguration() {
	hasFile := hasConfigFile()
	if hasFile {
		return
	}

	// println("File is not exist.")
	fullName, err := picConfigFullPath()
	if err != nil {
		panic(err)
	}
	// 创建文件
	if _, err := os.Stat(fullName); err != nil {
		if os.IsNotExist(err) {
			// 创建文件
			err := os.Mkdir(filepath.Dir(fullName), 0644)
			if err != nil {
				panic(err)
			}
			f, err := os.Create(fullName)
			if err != nil {
				panic(err)
			}
			// 关闭文件
			err = f.Close()
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	picBed := Config{
		Github: &Github{
			RepoName:     "",
			BranchName:   "",
			Token:        "",
			StoragePath:  "",
			DefineDomain: "",
		},
		TencentCOS: &CosConfig{
			SecretKey:        "",
			SecretId:         "",
			AppId:            "",
			Region:           "",
			StorageSpaceName: "",
			BucketName:       "",
			SystemVersion:    "5",
			MemoryPath:       "",
		},
		PicBedChoose: 0,
	}

	output, err := json.MarshalIndent(picBed, "", "\t\t")
	if err != nil {
		panic(err)
	}

	fn, err := picConfigFullPath()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fn, output, 0644)
	if err != nil {
		panic(err)
	}

}

// hasConfigFile 判断是否存在pic-config.json文件
func hasConfigFile() bool {
	fn, err := picConfigFullPath()
	if err != nil {
		return false
	}
	if _, err := os.Stat(fn); err != nil {
		return false
	}
	return true
}

func picConfigFullPath() (string, error) {
	return filepath.Abs(PicConfigPath + PicConfigFilename)
}
