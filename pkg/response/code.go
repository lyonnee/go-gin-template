package response

const (
	CODE_CALL_SUCCESS uint16 = 200

	fail_code uint16 = 10000 + iota

	CODE_NOT_TOKEN              = fail_code + iota
	CODE_TOKEN_FORMAT_INCORRECT = fail_code + iota
	CODE_TOKEN_INVALID          = fail_code + iota

	// 无效query参数
	CODE_INVALID_QUERY_ARGUMENT = fail_code + iota
	// 无效path参数
	CODE_INVALID_PATH_ARGUMENT = fail_code + iota
	// 无效body参数
	CODE_INVALID_BODY_ARGUMENT = fail_code + iota

	// 无效操作
	CODE_INVALID_OPERATION = fail_code + iota
	// 服务端错误
	CODE_SERVER_ERROR = fail_code + iota
	// 处理请求失败
	CODE_PROCESSING_REQUEST_FAILURE = fail_code + iota

	// 资源不存在
	CODE_RESOURCE_NOT_FOUND = fail_code + iota

	CODE_SEARCH_NULL = 10100
)
