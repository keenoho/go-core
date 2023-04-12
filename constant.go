package core

// response code
var (
	CODE_SYSTEM_ERROR    = 10001
	CODE_PARAMS_MISSING  = 10101
	CODE_PARAMS_UNEXPECT = 10102
)

// CodeMsgMap response code message map
var CodeMsgMap map[int]string = map[int]string{
	CODE_SYSTEM_ERROR:    "系统异常",
	CODE_PARAMS_MISSING:  "参数缺失",
	CODE_PARAMS_UNEXPECT: "参数不正确",
}

// CodeStatusMap response code status
var CodeStatusMap map[int]int = map[int]int{
	CODE_SYSTEM_ERROR:    500,
	CODE_PARAMS_MISSING:  400,
	CODE_PARAMS_UNEXPECT: 400,
}
