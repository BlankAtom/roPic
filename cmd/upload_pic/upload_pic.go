package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/blackswords/roPic/internal/config"
	"os"
)

// configPathName 是配置文件相对于运行程序的路径
// 但是在main方法中，会对这个值进行修改，从相对路径修改到绝对路径
// 这是因为对于typora自定义命令而言，运行时路径是正在编辑的笔记的路径
// 为了避免这一冲突导致的错误，孤儿修改成为绝对路径
var configPathName = "config.yml"

// printHelp 打印帮助信息，但事实上，我很快就会将他删掉
func printHelp(s string) {
	println(s)
	println("Param struct:")
	println("\t*.exe upload|-u <filename>")
}

// CreateURLtoString 将传入的url进行改装，使其变为需要的格式字符串并返回
// 现在使用的是markdown格式，之后也许会对其进行扩充
func CreateURLtoString(url string) (copyString string) {
	s := "![img](" + url + ")"
	return s
}

func main() {
	al := len(os.Args)

	if al < 3 {
		printHelp("Too least param.")
	} else {
		if os.Args[1] == "upload" || os.Args[1] == "-u" {
			if al == 2 {
				printHelp("Need the third param.")
				return
			}

			config.InitPicConfig()

			// 遍历参数中的文件，上传
			for i := 2; i < al; i++ {
				path := os.Args[i]
				//var t TencentYun
				//t.InitConfig()

				link, err := config.UploadFile(path)
				if err != nil {
					println(err.Error())
					continue
				}
				//_, key := UploadFile(path, t)
				err = clipboard.WriteAll(CreateURLtoString(link))
				if err != nil {
					println(err.Error())
				}
				if i == 2 {
					fmt.Printf("[UPLOADER SUCCESS]:\n")
				}
				fmt.Printf("%s\n", link)
			}

		}
	}
}
