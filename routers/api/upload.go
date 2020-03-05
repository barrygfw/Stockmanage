package api

import (
	"graduationProjectPeng/pkg/e"
	"graduationProjectPeng/pkg/logging"
	"graduationProjectPeng/pkg/upload"
	"graduationProjectPeng/service/common"

	"github.com/gin-gonic/gin"
)

/**
上传图片
api : /api/upload/image
params : form-data
image -> file
*/
func UploadImage(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		common.Json_return(c, code, "")
		return
	}

	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		savePath := upload.GetImagePath()

		src := savePath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(savePath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["imageUrl"] = upload.GetImageFullUrl(imageName)
				logging.Info(data["imageUrl"] + "upload success!")
			}
		}
	}

	common.Json_return(c, code, data)
}
