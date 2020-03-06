package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已过期",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_USER_NONE:                 "用户不存在",
	ERROR_USER_PASS:                 "密码错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验失败，请检查图片格式或大小",
	ERROR_UPDATE_STOCK_FAIL:         "更新商品库存失败",
	ERROR_ADD_INOUT_ROW_FAIL:        "添加出入库记录失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
