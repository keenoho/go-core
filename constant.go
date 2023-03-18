package core

// header
var HEADER_APP = "x-app"
var HEADER_SIGNATURE = "x-signature"

// 签名特征
var CTX_SIGNATURE_DATA_KEY = "signatureData"

// 响应方式
var RESPONSE_MODE_HTTP_STATUS = "httpStatus"
var RESPONSE_MODE_ALWAYS200 = "always200"

// 响应Code
var (
	CODE_SYSTEM_ERROR    = 10001
	CODE_PARAMS_MISSING  = 10101
	CODE_PARAMS_UNEXPECT = 10102
)

// 响应Code对应文案
var CodeMsgMap map[int]string = map[int]string{
	CODE_SYSTEM_ERROR:    "系统异常",
	CODE_PARAMS_MISSING:  "参数缺失",
	CODE_PARAMS_UNEXPECT: "参数不正确",
}

// 响应Code对应http status
var CodeStatusMap map[int]int = map[int]int{
	CODE_SYSTEM_ERROR:    500,
	CODE_PARAMS_MISSING:  400,
	CODE_PARAMS_UNEXPECT: 400,
}
