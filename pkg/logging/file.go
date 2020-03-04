package logging

import (
	"fmt"
	"graduationProjectPeng/pkg/file"
	"graduationProjectPeng/pkg/setting"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", setting.AppSetting.LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", setting.AppSetting.LogSaveName, time.Now().Format(setting.AppSetting.LogTimeFormat), setting.AppSetting.LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) (*os.File, error) {
	perm := file.CheckPermission(filePath)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: $s", filePath)
	}

	err := file.IsNotExistMkdir(filePath)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", filePath, err)
	}

	f, err := file.Open(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile: %v", err)
	}

	return f, nil
}
