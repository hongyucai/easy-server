package supports

import (
	"github.com/kataras/iris/v12"
)

const (
	// key定义
	CODE string = "code"
	MSG  string = "msg"
	DATA string = "data"

	StatusErrCode int =0
	StatusSucCode int =1

	// msg define
	Success                  = "恭喜, 成功"
	OptionSuccess     string = "恭喜, 操作成功"
	OptionFailur      string = "抱歉, 操作失败"
	ParseParamsFailur string = "解析参数失败"

	RegisteSuccess     string = "恭喜, 注册成功"
	RegisteFailur      string = "注册失败"
	LoginSuccess       string = "恭喜, 登录成功"
	LoginFailur        string = "登录失败"

	CreateFailur string = "创建失败"
	CreateSuccess string = "创建成功"
	DeleteSuccess string = "删除成功"
	DeleteFailur  string = "删除失败"

	UsernameFailur             string = "用户名错误"
	PasswordFailur             string = "密码错误"
	TokenCreateFailur          string = "生成token错误"
	TokenExactFailur           string = "token不存在或header设置不正确"
	TokenExpire                string = "回话已过期"
	TokenParseFailur           string = "token解析错误"
	TokenParseFailurAndEmpty   string = "解析错误,token为空"
	TokenParseFailurAndInvalid string = "解析错误,token无效"
	TokenRefreshFailur         string = "token刷新失败"
	NotFound                   string = "您请求的url不存在"
	PermissionsLess            string = "权限不足非法操作"
	StatusInternalServerError  string = "服务器内部错误"
	// value define
)

// 200 define
func Ok_(ctx iris.Context, msg string) {
	Ok(ctx, msg, nil)
}

func Ok(ctx iris.Context, msg string, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		CODE: StatusSucCode,
		MSG:  msg,
		DATA: data,
	})
}

// 401 error define
func Unauthorized(ctx iris.Context, msg string, data interface{}) {
	unauthorized := iris.StatusUnauthorized

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		CODE: unauthorized,
		MSG:  msg,
		DATA: data,
	})
}
// common error define
func Error(ctx iris.Context, msg string, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		CODE: StatusErrCode,
		MSG:  msg,
		DATA: data,
	})
}
