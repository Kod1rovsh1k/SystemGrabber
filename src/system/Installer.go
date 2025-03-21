package system

import (
	"fmt"
	"os"
)

func Installer(srcPath string) (err error) {
	rootDir := srcPath + "\\Log"
	errDir := os.MkdirAll(rootDir, os.ModePerm)
	if errDir != nil {
		return errDir
	} else {
		fmt.Println("Root dir was created ->", rootDir)
	}
	return err
}
