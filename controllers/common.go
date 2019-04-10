package controllers

const (
	CODE_OK = 10000 + iota
	CODE_INVALID_PARAMETER
	CODE_UNKNOWN_USER
	CODE_INTERNAL_ERROR
	CODE_PARSE_DATA_FAIL
	CODE_NO_PERMISSIONS
	CODE_TMP_SYS_ERROR
	CODE_LOST_PARAMETER
	CODE_DB_ERROR
	CODE_NO_AUTH
	CODE_REDIRECT_302 = 20001

	MSG_OK = "ok"
)