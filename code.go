package core

// response code
var (
	CODE_SERVER_ERROR    int64 = 10001
	CODE_SERVICE_ERROR   int64 = 10002
	CODE_PARAMS_MISSING  int64 = 10101
	CODE_PARAMS_UNEXPECT int64 = 10102
)

// response code message map
var CODE_MESSAGE_MAP = map[int64]string{
	CODE_SERVER_ERROR:    "System Error",
	CODE_SERVICE_ERROR:   "Service Error",
	CODE_PARAMS_MISSING:  "Params Missing",
	CODE_PARAMS_UNEXPECT: "Params Unexpect",
}

// response code status
var CODE_STATUS_MAP = map[int64]int{
	CODE_SERVER_ERROR:    500,
	CODE_SERVICE_ERROR:   500,
	CODE_PARAMS_MISSING:  400,
	CODE_PARAMS_UNEXPECT: 400,
}

func AddCodeMessageMap(msgMap map[int64]string) {
	for key, value := range msgMap {
		CODE_MESSAGE_MAP[key] = value
	}
}

func AddCodeStatusMap(statusMap map[int64]int) {
	for key, value := range statusMap {
		CODE_STATUS_MAP[key] = value
	}
}
