package response

const (
	OK                    = 200
	NotOK                 = 406
	Unauthorized          = 401
	Forbidden             = 403
	NotFound              = 404
	InternalServerError   = 500
	ParameterBindingError = 506
)

const (
	OKMsg                      = "操作成功"
	NotOKMsg                   = "操作失败"
	UnauthorizedMsg            = "认证失败，请重新登录"
	ForbiddenMsg               = "无权限访问资源"
	NotFoundMsg                = "资源找不到"
	InternalServerErrorMsg     = "服务器内部错误"
	IdempotenceTokenEmptyMsg   = "幂等性token为空"
	IdempotenceTokenInvalidMsg = "请勿重复提交"
)

var CustomError = map[int]string{
	OK:                  OKMsg,
	NotOK:               NotOKMsg,
	Unauthorized:        UnauthorizedMsg,
	Forbidden:           ForbiddenMsg,
	NotFound:            NotFoundMsg,
	InternalServerError: InternalServerErrorMsg,
}
