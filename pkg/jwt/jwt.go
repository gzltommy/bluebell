package jwt

import (
	"bluebell/setting"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个UserID字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成 AccessToken 和 RefreshToken
/*
* RefreshToken 用于获取新的 AccessToken。这样可以缩短 AccessToken 的过期时间保证安全，同时又不会因为频繁过期重新要求用户登录。
* 用户在初次认证时，Refresh Token 会和 AccessToken 一起返回。
* 前端携带 RefreshToken 向 Authing 端点发起请求时，Authing 每次都会返回相同的 RefreshToken 和新的 AccessToken，直到 RefreshToken 过期。
* 注意：
* 为了 refresh token 的安全，Oauth2.0 要求，refresh token 一定要保存在使用方的服务器上，而绝不能存放在移动 app、PC 端软件、浏览器上，也不能在网络上随便传递。
* 调用 refresh 接口的时候，一定是从使用方服务器到鉴权服务器的 https 访问。所以，refresh token 比 access token 隐蔽得多，也安全得多。当然，这需要使用方正确的遵守 Oauth2.0 的要求。
 */
func GenToken(userID uint64, username string) (aToken, rToken string, err error) {
	// 加密并获得完整的编码后的字符串 token
	aToken, err = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		MyClaims{
			userID,   // 自定义字段
			username, // 自定义字段
			jwt.StandardClaims{ // JWT 规定的 7 个官方字段
				ExpiresAt: time.Now().
					Add(time.Duration(setting.Cfg.Auth.JwtExpire) * time.Second). // 过期时间
					Unix(),
				Issuer: setting.Cfg.AppName, // 签发人
			},
		}).SignedString([]byte(setting.Cfg.Auth.JwtSecret))
	if err != nil {
		return "", "", err
	}

	// refresh token 不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 10).Unix(), // 过期时间
			Issuer:    setting.Cfg.AppName,                        // 签发人
		}).SignedString([]byte(setting.Cfg.Auth.JwtSecret))
	return
}

// GenToken2 生成 Token
func GenToken2(userID uint64, username string) (token string, err error) {
	// 加密并获得完整的编码后的字符串token
	token, err = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		MyClaims{
			userID,   // 自定义字段
			username, // 自定义字段
			jwt.StandardClaims{ // JWT 规定的 7 个官方字段
				ExpiresAt: time.Now().
					Add(time.Duration(setting.Cfg.Auth.JwtExpire) * time.Second). // 过期时间
					Unix(),
				Issuer: setting.Cfg.AppName, // 签发人
			},
		}).SignedString([]byte(setting.Cfg.Auth.JwtSecret))
	return
}

// ParseToken 解析 token
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	// 校验 token
	if !token.Valid {
		err = errors.New("invalid token")
	}
	return
}

// RefreshToken 刷新 AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token 无效直接返回
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}

	// 从旧 access token 中解析出 claims 数据,解析出 payload 负载信息
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 当 access token 是过期错误 并且 refresh token 没有过期时就创建一个新的 access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID, claims.Username)
	}
	return
}

func keyFunc(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("token parse error")
	}
	return []byte(setting.Cfg.Auth.JwtSecret), nil
}
