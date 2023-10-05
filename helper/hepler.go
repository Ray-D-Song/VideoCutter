package helper

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"
)

func TryCatch(task func(), errCallback func(R any)) {
	defer func() {
		if r := recover(); r != nil {
			errCallback(r)
			debug.PrintStack()
		}
	}()
	task()
}

func Exp(expression bool, trueRes any, falseRes any) any {
	if expression {
		return trueRes
	} else {
		return falseRes
	}
}

func CopyToClipboard(content string) error {
	command := exec.Command("pbcopy")
	iw, err := command.StdinPipe()
	if err != nil {
		return err
	}

	if err := command.Start(); err != nil {
		return err
	}
	if _, err := iw.Write([]byte(content)); err != nil {
		return err
	}
	if err := iw.Close(); err != nil {
		return err
	}
	return command.Wait()
}

func TimeConverter(inputTime string) float64 {
	parsedTime, err := time.Parse("15:04:05", inputTime)

	if err != nil {
		panic("Time parse error")
	}
	return float64(parsedTime.Hour()*3600 + parsedTime.Minute()*60 + parsedTime.Second())
}

func ColorLog(info any) {
	fmt.Printf("\033[34m%v\033[0m\n", info)
}

func GetFileNameAndExtension(pathOrName string) (string, string) {
	fileName := filepath.Base(pathOrName)

	dotIndex := strings.LastIndex(fileName, ".")

	if dotIndex == -1 {
		return fileName, ""
	}

	fileExtension := fileName[dotIndex+1:]

	return fileName[:dotIndex], fileExtension
}
