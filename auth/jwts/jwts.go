package jwts

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"go-mod/core/config"
	"go-mod/gateway/supports"
	"log"
	"strings"
	"sync"
	"time"
)
const (
	DefaultContextKey = "iris-jwt"
)
// Config是用于为jwts中间件指定配置选项的结构。
type Config struct {
	//将返回Key以验证JWT的函数。
	//它可以是共享密钥或公共密钥。
	// The function that will return the Key to validate the JWT.
	// It can be either a shared secret or a public key.
	// Default value: nil
	ValidationKeyGetter jwt.Keyfunc
	//请求中的属性名称（用户（令牌）信息）
	//将存储来自JWT的内容。
	// The name of the property in the request where the user (&token) information
	// from the JWT will be stored.
	// Default value: "jwts"
	ContextKey string
	//验证令牌出错时将调用的函数
	// The function that will be called when there's an error validating the token
	// Default value:
	ErrorHandler errorHandler
	// 一个布尔值，指示是否需要凭据
	// A boolean indicating if the credentials are required or not
	// Default value: false
	CredentialsOptional bool
	//从请求中提取令牌的函数
	// A function that extracts the token from the request
	// Default: FromAuthHeader (i.e., from Authorization header as bearer token)
	Extractor TokenExtractor
	//调试标志打开调试输出
	// Debug flag turns on debugging output
	// Default: false
	Debug bool
	//设置后，所有带有OPTIONS方法的请求都将使用身份验证
	//如果启用此选项，则还应向iris.Options（...）注册路线
	// Default: false
	EnableAuthOnOptions bool
	//设置后，中间件软件会验证令牌是否已使用特定的签名算法进行签名
	//如果签名方法不是恒定的，则ValidationKeyGetter回调可用于实现其他检查
	//对于避免此处所述的安全问题很重要：https：//auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
	// When set, the middelware verifies that tokens are signed with the specific signing algorithm
	// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
	// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
	// Default: nil
	SigningMethod jwt.SigningMethod
	//设置后，每次都会检查令牌的到期时间
	//如果令牌已过期，则会返回过期错误
	// When set, the expiration time of token will be check every time
	// if the token was expired, expiration error will be returned
	// Default: false
	Expiration bool
}

type (
	errorHandler func(context.Context, string)
	// TokenExtractor是一个将上下文作为输入并返回的函数
	//令牌或错误。 仅在尝试时返回错误
	//找到了指定令牌的信息，但是信息以某种方式不正确
	//组成。 在根本不存在令牌的情况下，这不应
	//被视为错误。 在这种情况下，应返回一个空字符串。
	// TokenExtractor is a function that takes a context as input and returns
	// either a token or an error.  An error should only be returned if an attempt
	// to specify a token was found, but the information was somehow incorrectly
	// formed.  In the case where a token is simply not present, this should not
	// be treated as an error.  An empty string should be returned in that case.
	TokenExtractor func(context.Context) (string, error)
	// Middleware the middleware for JSON Web tokens authentication method
	Jwts struct {
		Config Config
	}
)

var (
	jwts *Jwts
	lock sync.Mutex
)

// Serve the middleware's action
func Serve(ctx context.Context) bool {
	ConfigJWT()
	if err := jwts.CheckJWT(ctx); err != nil {
		//supports.Unauthorized(ctx, supports.Token_failur, nil)
		//ctx.StopExecution()
		golog.Errorf("Check jwt error, %s", err)
		return false
	}
	return true
	// If everything ok then call next.
	//ctx.Next()
}
// below 3 method is get token from url
// FromAuthHeader is a "TokenExtractor" that takes a give context and extracts
// the JWT token from the Authorization header.
func FromAuthHeader(ctx context.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}
	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", fmt.Errorf("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}

// below 3 method is get token from url
// FromParameter returns a function that extracts the token from the specified
// query string parameter
func FromParameter(param string) TokenExtractor {
	return func(ctx context.Context) (string, error) {
		return ctx.URLParam(param), nil
	}
}

// below 3 method is get token from url
// FromFirst returns a function that runs multiple token extractors and takes the
// first token it finds
func FromFirst(extractors ...TokenExtractor) TokenExtractor {
	return func(ctx context.Context) (string, error) {
		for _, ex := range extractors {
			token, err := ex(ctx)
			if err != nil {
				return "", err
			}
			if token != "" {
				return token, nil
			}
		}
		return "", nil
	}
}

func (m *Jwts) logf(format string, args ...interface{}) {
	if m.Config.Debug {
		log.Printf(format, args...)
	}
}

// Get returns the user (&token) information for this client/request
func (m *Jwts) Get(ctx context.Context) *jwt.Token {
	return ctx.Values().Get(m.Config.ContextKey).(*jwt.Token)
}

// CheckJWT the main functionality, checks for token
func (m *Jwts) CheckJWT(ctx context.Context) error {
	if !m.Config.EnableAuthOnOptions {
		if ctx.Method() == iris.MethodOptions {
			return nil
		}
	}

	// Use the specified token extractor to extract a token from the request
	token, err := m.Config.Extractor(ctx)
	// If an error occurs, call the error handler and return an error
	if err != nil {
		m.logf("Error extracting JWT: %v", err)
		m.Config.ErrorHandler(ctx, supports.TokenExactFailur)
		return fmt.Errorf("Error extracting token: %v", err)
	}

	// If the token is empty...
	if token == "" {
		// Check if it was required
		if m.Config.CredentialsOptional {
			m.logf("  No credentials found (CredentialsOptional=true)")
			// No error, just no token (and that is ok given that CredentialsOptional is true)
			return nil
		}

		m.logf("  Error: No credentials found (CredentialsOptional=false)")
		// If we get here, the required token is missing
		m.Config.ErrorHandler(ctx, supports.TokenParseFailurAndEmpty)
		return fmt.Errorf(supports.TokenParseFailurAndEmpty)
	}

	// Now parse the token

	parsedToken, err := jwt.Parse(token, m.Config.ValidationKeyGetter)
	// Check if there was an error in parsing...
	if err != nil {
		m.logf("Error parsing token1: %v", err)
		m.Config.ErrorHandler(ctx, supports.TokenExpire)
		return fmt.Errorf("Error parsing token2: %v", err)
	}

	if m.Config.SigningMethod != nil && m.Config.SigningMethod.Alg() != parsedToken.Header["alg"] {
		message := fmt.Sprintf("Expected %s signing method but token specified %s",
			m.Config.SigningMethod.Alg(),
			parsedToken.Header["alg"])
		m.logf("Error validating token algorithm: %s", message)
		m.Config.ErrorHandler(ctx, supports.TokenParseFailur) // 算法错误
		return fmt.Errorf("Error validating token algorithm: %s", message)
	}

	// Check if the parsed token is valid...
	if !parsedToken.Valid {
		m.logf(supports.TokenParseFailurAndInvalid)
		m.Config.ErrorHandler(ctx, supports.TokenParseFailurAndInvalid)
		return fmt.Errorf(supports.TokenParseFailurAndInvalid)
	}

	if m.Config.Expiration {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			if expired := claims.VerifyExpiresAt(time.Now().Unix(), true); !expired {
				return fmt.Errorf(supports.TokenExpire)
			}
		}
	}
	//m.logf("JWT: %v", parsedToken)
	// If we get here, everything worked and we can set the
	// user property in context.
	ctx.Values().Set(m.Config.ContextKey, parsedToken)

	return nil
}

// ------------------------------------------------------------------------
// ------------------------------------------------------------------------

// jwt中间件配置
func ConfigJWT() {
	if jwts != nil {
		return
	}

	lock.Lock()
	defer lock.Unlock()

	if jwts != nil {
		return
	}

	c := Config{
		ContextKey: DefaultContextKey,
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(config.O.Secret), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		ErrorHandler: func(ctx context.Context, errMsg string) {
			supports.Unauthorized(ctx, errMsg, nil)
		},
		// 指定func用于提取请求中的token
		Extractor: FromAuthHeader,
		// if the token was expired, expiration error will be returned
		Expiration:          true,
		Debug:               true,
		EnableAuthOnOptions: false,
	}
	jwts = &Jwts{Config: c}
	//return &Jwts{Config: c}
}

type AdminClaims struct {
	User AdminUserJwt `json:"user"`
	jwt.StandardClaims
}

type ApiClaims struct {
	ApiUser ApiUserJwt `json:"axeUser"`
	jwt.StandardClaims
}
// 在登录成功的时候生成token
func GenerateAdminUserToken(id,rid,bid int64,username,name string,menumap ,powermap map[string]interface{}) (string, error) {
	//expireTime := time.Now().Add(60 * time.Second)
	expireTime := time.Now().Add(time.Duration(config.O.JWTTimeout) * time.Second)
	claims := AdminClaims{
		AdminUserJwt{
			id,
			username,
			name,
			rid,
			bid,
			menumap,
			powermap,
		},
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer:    "iris-admin-jwt",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(config.O.Secret))
	return token, err
}
func GenerateApiUserToken(id,app_id int64) (string, error) {
	//expireTime := time.Now().Add(60 * time.Second)
	expireTime := time.Now().Add(time.Duration(config.O.JWTTimeout) * time.Second)
	claims := ApiClaims{
		ApiUserJwt{
			id,
			app_id,
		},
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer:    "iris-api-jwt",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(config.O.Secret))
	return token, err
}

// 解析token的信息为当前用户
func ParseAdminUserToken(ctx context.Context) (*AdminUserJwt, bool) {
	mapClaims := (jwts.Get(ctx).Claims).(jwt.MapClaims)
	userMap, ok := mapClaims["user"].(map[string]interface{})
	if !ok{
		return nil, false
	}
	id, _ := userMap["id"].(int64)
	username, _ := userMap["username"].(string)
	name, _ := userMap["name"].(string)
	rid, _ := userMap["rid"].(int64)
	bid, _ := userMap["bid"].(int64)
	menu, _ := userMap["menumap"].(map[string]interface{})
	power, _ := userMap["powermap"].(map[string]interface{})
	user := AdminUserJwt{
		Id:       id,
		Username: username,
		Name:	name,
		Rid:      rid,
		Bid:      bid,
		Menumap:  menu,
		Powermap: power,
	}
	return &user, true
}

func ParseApiUserToken(ctx context.Context) (*ApiUserJwt, bool) {
	mapClaims := (jwts.Get(ctx).Claims).(jwt.MapClaims)
	axeUserMap, ok := mapClaims["axeUser"].(map[string]interface{})
	if !ok{
		return nil, false
	}
	id, _ := axeUserMap["id"].(int64)
	app_id, _ := axeUserMap["app_id"].(int64)
	axeUser := ApiUserJwt{
		Id:       id,
		AppId: 	  app_id,
	}
	return &axeUser, true
}