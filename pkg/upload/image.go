package upload

import (
	"fmt"
	"graduationProjectPeng/pkg/file"
	"graduationProjectPeng/pkg/logging"
	"graduationProjectPeng/pkg/setting"
	"graduationProjectPeng/pkg/util"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

/**
获取图片访问路径
*/
func GetImageFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/cdn/images/" + name
}

/**
获取文件名
*/
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

/**
获取图片保存路径
*/
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

/**
检查图片格式
*/
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

/**
检查图片大小是否合法
*/
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f) //获取到的大小单位为字节
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= (setting.AppSetting.ImageMaxSize * 1048576)
}

/**
检查图片
*/
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwn err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
