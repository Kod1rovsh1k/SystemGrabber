package system

import (
	"fmt"
	"os"
)

func SysManager() {
	rootDir, _ := os.Getwd()
	errInstall := Installer(rootDir)
	errInfo := GetInfo(rootDir)

	if errInstall != nil {
		println(errInstall.Error())
	}

	if errInfo != nil {
		println(errInfo.Error())
	}

	fmt.Println("To continue, press any key. . .")
	fmt.Scanf(" ")
}
