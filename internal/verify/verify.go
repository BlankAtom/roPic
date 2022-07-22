package verify

import (
	"os"
)

func IsFileInputCheck(param string) bool {
	return IsFile(param)
}

func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}
