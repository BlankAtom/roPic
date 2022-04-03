package main

import (
	"gopkg.in/yaml.v2"
	"os"
)

// CreateFile 创建文件
func CreateFile(pathName string) (err error) {
	f, err := os.Create(pathName)
	if err != nil {
		return err
	}
	err = f.Close()
	return err
}

// InitConfigFileBeforeCreate 创建完文件后调用，使目标配置文件初始化
func InitConfigFileBeforeCreate(path string) (err error) {
	var t TencentYun
	t.BucketFromTencent.SystemVersion = "v5"
	marshal, err := yaml.Marshal(t)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, marshal, 0777)
	return err
}

// CheckConfigFile 检查配置文件是否存在
func CheckConfigFile() {
	if _, err := os.Stat(configPathName); err != nil {
		if os.IsNotExist(err) {
			err := CreateFile(configPathName)
			if err != nil {
				println(err.Error())
			}
			err = InitConfigFileBeforeCreate(configPathName)
			if err != nil {
				println(err.Error())
			}
		}
	}
}
