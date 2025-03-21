package system

import (
	"fmt"
	"os/exec"
)

func GetInfo(srcPath string) (err error) {

	errCommand := exec.Command("cmd", "/c", "systeminfo /FO LIST >> ", srcPath+"\\Log\\systeminfo.txt").Run()

	if errCommand != nil {
		println(errCommand.Error())
	} else {
		fmt.Println("Info was successfully get it")
	}

	return err
}
