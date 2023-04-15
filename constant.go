package core

// response code
var (
	CODE_SYSTEM_ERROR    = 10001
	CODE_PARAMS_MISSING  = 10101
	CODE_PARAMS_UNEXPECT = 10102
	CODE_SERVICE_ERROR   = 20001
)

// CodeMsgMap response code message map
var CodeMsgMap map[int]string = map[int]string{
	CODE_SYSTEM_ERROR:    "System Error",
	CODE_PARAMS_MISSING:  "Params Missing",
	CODE_PARAMS_UNEXPECT: "Params Unexpect",
	CODE_SERVICE_ERROR:   "Service Error",
}

// CodeStatusMap response code status
var CodeStatusMap map[int]int = map[int]int{
	CODE_SYSTEM_ERROR:    500,
	CODE_PARAMS_MISSING:  400,
	CODE_PARAMS_UNEXPECT: 400,
	CODE_SERVICE_ERROR:   500,
}
