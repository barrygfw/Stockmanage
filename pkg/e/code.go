package e

const (
	//基本code
	SUCCESS        = 200 //成功
	ERROR          = 500 //失败
	INVALID_PARAMS = 501 //参数错误

	//token相关报错
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	//用户相关报错
	ERROR_USER_NONE = 30001
	ERROR_USER_PASS = 30002

	//上传图片相关
	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 40001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 40002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 40003

	//出入库相关
	ERROR_UPDATE_STOCK_FAIL  = 50001
	ERROR_ADD_INOUT_ROW_FAIL = 50002
)
